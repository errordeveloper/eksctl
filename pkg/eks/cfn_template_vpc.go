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
// 	return newRef(name)
// }

// func newStringParameter(t *cloudformation.Template, name, defaultValue string) interface{} {
// 	return newParameter(t, name, "String", defaultValue)
// }

// func newSub(sub string) interface{} {
// 	return map[string]string{"Sub": sub}
// }

func newRef(refName string) interface{} {
	return map[string]string{"Ref": refName}
}

func newResource(t *cloudformation.Template, name string, resource interface{}) interface{} {
	t.Resources[name] = resource
	return newRef(name)
}

type resourceRefsForVPC struct {
	VPC            interface{}
	Subnets        []interface{}
	SecurityGroups []interface{}
}

func addResourcesForVPC(t *cloudformation.Template, stackName string, globalCIDR *net.IPNet, subnets map[string]*net.IPNet) *resourceRefsForVPC {
	refs := &resourceRefsForVPC{}
	refs.VPC = newResource(t, "VPC", &cloudformation.UntypedAWSEC2VPC{
		CidrBlock:          globalCIDR.String(),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
		// Tags: []cloudformation.Tag{{
		// 	Key:   "Name",
		// 	Value: stackName + ".VPC",
		// }},
	})

	refIG := newResource(t, "InternetGateway", &cloudformation.AWSEC2InternetGateway{})
	newResource(t, "VPCGatewayAttachment", &cloudformation.UntypedAWSEC2VPCGatewayAttachment{
		InternetGatewayId: refIG,
		VpcId:             refs.VPC,
	})

	refRT := newResource(t, "RouteTable", &cloudformation.UntypedAWSEC2RouteTable{
		VpcId: refs.VPC,
		// Tags: []cloudformation.Tag{
		// 	{Key: "Name", Value: "Public Subnets"},
		// 	{Key: "Network", Value: "Public"},
		// },
	})

	newResource(t, "Route", &cloudformation.UntypedAWSEC2Route{
		RouteTableId:         refRT,
		DestinationCidrBlock: "0.0.0.0/0",
		GatewayId:            newRef("InternetGateway"),
	})

	for az, subnet := range subnets {
		refSubnet := newResource(t, "Subnet_"+az, &cloudformation.UntypedAWSEC2Subnet{
			AvailabilityZone: az,
			CidrBlock:        subnet.String(),
			VpcId:            refs.VPC,
			// Tags: []cloudformation.Tag{{
			// 	Key:   "Name",
			// 	Value: stackName + "." + name,
			// }},
		})
		newResource(t, "RouteTableAssociation_"+az, &cloudformation.UntypedAWSEC2SubnetRouteTableAssociation{
			SubnetId:     refSubnet,
			RouteTableId: refRT,
		})
		refs.Subnets = append(refs.Subnets, refSubnet)
	}

	refSG := newResource(t, "ControlPlaneSecurityGroup", &cloudformation.UntypedAWSEC2SecurityGroup{
		GroupDescription: "Cluster communication with worker nodes",
		VpcId:            refs.VPC,
	})
	refs.SecurityGroups = []interface{}{refSG}

	return refs
}
