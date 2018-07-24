package cfn

import (
	"net"

	"github.com/awslabs/goformation/cloudformation"
)

const (
	OutputClusterCertificateAuthorityData  = "Cluster.CertificateAuthorityData"
	OutputClusterEndpoint                  = "Cluster.Endpoint"
	OutputClusterARN                       = "Cluster.ARN"
	OutputClusterVPC                       = "Cluster.VPC"
	OutputClusterSubnets                   = "Cluster.Subnets"
	OutputClusterControlPlaneSecurityGroup = "Cluster.ControlPlaneSecurityGroup"

	iamAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
)

type clusterResourceSet struct {
	resourceSet      *resourceSet
	vpcRefs          *resourceRefsForVPC
	controlPlaneRefs *resourceRefsForControlPlane
}

type resourceRefsForVPC struct {
	vpc            interface{}
	subnets        []interface{}
	securityGroups []interface{}
}

type resourceRefsForControlPlane struct {
}

func NewClusterResourceSet() *clusterResourceSet {
	return &clusterResourceSet{
		resourceSet: newResourceSet(),
	}
}

func (c *clusterResourceSet) AddAllResources(availabilityZones []string) {
	_, globalCIDR, _ := net.ParseCIDR("192.168.0.0/16")

	subnets := map[string]*net.IPNet{}
	_, subnets[availabilityZones[0]], _ = net.ParseCIDR("192.168.64.0/18")
	_, subnets[availabilityZones[1]], _ = net.ParseCIDR("192.168.128.0/18")
	_, subnets[availabilityZones[2]], _ = net.ParseCIDR("192.168.192.0/18")

	c.addResourcesForVPC(globalCIDR, subnets)
	c.addResourcesForControlPlane("1.10")
}

func (c *clusterResourceSet) RenderJSON() ([]byte, error) {
	return c.resourceSet.renderJSON()
}

func (c *clusterResourceSet) newResource(name string, resource interface{}) interface{} {
	return c.resourceSet.newResource(name, resource)
}

func (c *clusterResourceSet) newOutput(name string, value interface{}) {
	c.resourceSet.newOutput(name, value, true)
}

func (c *clusterResourceSet) newJoinedOutput(name string, value []interface{}) {
	c.resourceSet.newJoinedOutput(name, value, true)
}

func (c *clusterResourceSet) newOutputFromAtt(name, att string) {
	c.resourceSet.newOutputFromAtt(name, att, true)
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

	c.newOutput(OutputClusterVPC, refs.vpc)
	c.newJoinedOutput(OutputClusterControlPlaneSecurityGroup, refs.securityGroups)
	c.newJoinedOutput(OutputClusterSubnets, refs.subnets)
}

func (c *clusterResourceSet) addResourcesForControlPlane(version string) {
	c.newResource("ControlPlane", &cloudformation.UntypedAWSEKSCluster{
		Name: refStackName,
		RoleArn: c.newResource("ServiceRole", &cloudformation.AWSIAMRole{
			AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("eks.amazonaws.com"),
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

	c.newOutputFromAtt(OutputClusterCertificateAuthorityData, "ControlPlane.CertificateAuthorityData")
	c.newOutputFromAtt(OutputClusterEndpoint, "ControlPlane.Endpoint")
	c.newOutputFromAtt(OutputClusterARN, "ControlPlane.Arn")
}
