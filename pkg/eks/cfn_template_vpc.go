package eks

import (
	"net"

	"github.com/awslabs/goformation/cloudformation"
)

// func newParameter(t *cloudformation.Template, name, valueType, defaultValue string) interface{} {
// 	p := map[string]string{"Type": valueType}
// 	if defaultValue != "" {
// 		p["Default"] = defaultValue
// 	}
// 	t.Parameters[name] = p
// 	return makeRef(name)
// }

// func newStringParameter(t *cloudformation.Template, name, defaultValue string) interface{} {
// 	return newParameter(t, name, "String", defaultValue)
// }

// func newSub(sub string) interface{} {
// 	return map[string]string{"Sub": sub}
// }

func makeRef(refName string) interface{} {
	return map[string]string{"Ref": refName}
}

type clusterResourceSet struct {
	template *cloudformation.Template
	vpcRefs  *resourceRefsForVPC
}

type resourceRefsForVPC struct {
	vpc            interface{}
	subnets        []interface{}
	securityGroups []interface{}
}

func newClusterResourceSet() *clusterResourceSet {
	return &clusterResourceSet{
		template: cloudformation.NewTemplate(),
	}
}

func (c *clusterResourceSet) newResource(name string, resource interface{}) interface{} {
	c.template.Resources[name] = resource
	return makeRef(name)
}

func (c *clusterResourceSet) addResourcesForVPC(globalCIDR *net.IPNet, subnets map[string]*net.IPNet) {
	refs := &resourceRefsForVPC{}
	refs.vpc = c.newResource("VPC", &cloudformation.UntypedAWSEC2VPC{
		CidrBlock:          globalCIDR.String(),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
	})

	refIG := c.newResource("InternetGateway", &cloudformation.AWSEC2InternetGateway{})
	c.newResource("VPCGatewayAttachment", &cloudformation.UntypedAWSEC2VPCGatewayAttachment{
		InternetGatewayId: refIG,
		VpcId:             refs.vpc,
	})

	refRT := c.newResource("RouteTable", &cloudformation.UntypedAWSEC2RouteTable{
		VpcId: refs.vpc,
	})

	c.newResource("Route", &cloudformation.UntypedAWSEC2Route{
		RouteTableId:         refRT,
		DestinationCidrBlock: "0.0.0.0/0",
		GatewayId:            makeRef("InternetGateway"),
	})

	for az, subnet := range subnets {
		refSubnet := c.newResource("Subnet_"+az, &cloudformation.UntypedAWSEC2Subnet{
			AvailabilityZone: az,
			CidrBlock:        subnet.String(),
			VpcId:            refs.vpc,
		})
		c.newResource("RouteTableAssociation_"+az, &cloudformation.UntypedAWSEC2SubnetRouteTableAssociation{
			SubnetId:     refSubnet,
			RouteTableId: refRT,
		})
		refs.subnets = append(refs.subnets, refSubnet)
	}

	refSG := c.newResource("ControlPlaneSecurityGroup", &cloudformation.UntypedAWSEC2SecurityGroup{
		GroupDescription: "Cluster communication with worker nodes",
		VpcId:            refs.vpc,
	})
	refs.securityGroups = []interface{}{refSG}

	c.vpcRefs = refs
}
