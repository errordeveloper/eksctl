package eks

import (
	"net"

	"github.com/awslabs/goformation/cloudformation"
)

const (
	cfnOutputClusterCertificateAuthorityData = "ClusterCertificateAuthorityData"
	cfnOutputClusterEndpoint                 = "ClusterEndpoint"
	cfnOutputClusterARN                      = "ClusterARN"

	iamAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
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

type resourceRefsForControlPlane struct {
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

func (c *clusterResourceSet) newOutput(name string, value interface{}) {
	o := map[string]interface{}{"Value": value}
	c.template.Outputs[name] = o
}

func (c *clusterResourceSet) newOutputFromAtt(name, att string) {
	c.newOutput(name, map[string]string{"Fn::GetAtt": att})
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
		GatewayId:            refIG,
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

func (c *clusterResourceSet) addResourcesForControlPlane(name, version string) {
	asrpd := map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []interface{}{
			map[string]interface{}{
				"Effect": "Allow",
				"Principal": map[string][]string{
					"Service": []string{"eks.amazonaws.com"},
				},
				"Action": []string{"sts:AssumeRole"},
			},
		},
	}
	c.newResource("ControlPlane", &cloudformation.UntypedAWSEKSCluster{
		Name: name,
		RoleArn: c.newResource("ServiceRole", &cloudformation.AWSIAMRole{
			AssumeRolePolicyDocument: asrpd,
			ManagedPolicyArns: []string{
				iamAmazonEKSServicePolicyARN,
				iamAmazonEKSClusterPolicyARN,
			},
		}),
		Version: version,
		ResourcesVpcConfig: &cloudformation.UntypedAWSEKSCluster_ResourcesVpcConfig{
			SubnetIds:        c.vpcRefs.subnets,
			SecurityGroupIds: c.vpcRefs.securityGroups,
		},
	})

	c.newOutputFromAtt(cfnOutputClusterCertificateAuthorityData, "ControlPlane.CertificateAuthorityData")
	c.newOutputFromAtt(cfnOutputClusterEndpoint, "ControlPlane.Endpoint")
	c.newOutputFromAtt(cfnOutputClusterARN, "ControlPlane.Arn")
}
