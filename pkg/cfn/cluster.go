package cfn

import (
	"net"
	"strings"

	"github.com/awslabs/goformation/cloudformation"
)

const (
	OutputClusterCertificateAuthorityData  = "ClusterCertificateAuthorityData"
	OutputClusterEndpoint                  = "ClusterEndpoint"
	OutputClusterARN                       = "ClusterARN"
	OutputClusterVPC                       = "ClusterVPC"
	OutputClusterSubnets                   = "ClusterSubnets"
	OutputClusterControlPlaneSecurityGroup = "ClusterControlPlaneSecurityGroup"

	iamAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
)

type clusterResourceSet struct {
	resourceSet      *resourceSet
	vpcRefs          *resourceRefsForVPC
	controlPlaneRefs *resourceRefsForControlPlane
}

type resourceRefsForVPC struct {
	vpc            *cloudformation.StringIntrinsic
	subnets        []*cloudformation.StringIntrinsic
	securityGroups []*cloudformation.StringIntrinsic
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

	c.resourceSet.template.Description = "Dynamically generated EKS cluster template (with dedicated VPC & IAM role)"
}

func (c *clusterResourceSet) RenderJSON() ([]byte, error) {
	return c.resourceSet.renderJSON()
}

func (c *clusterResourceSet) newResource(name string, resource interface{}) *cloudformation.StringIntrinsic {
	return c.resourceSet.newResource(name, resource)
}

func (c *clusterResourceSet) newOutput(name string, value interface{}, export bool) {
	c.resourceSet.newOutput(name, value, export)
}

func (c *clusterResourceSet) newJoinedOutput(name string, value []*cloudformation.StringIntrinsic, export bool) {
	c.resourceSet.newJoinedOutput(name, value, export)
}

func (c *clusterResourceSet) newOutputFromAtt(name, att string, export bool) {
	c.resourceSet.newOutputFromAtt(name, att, export)
}

func (c *clusterResourceSet) addResourcesForVPC(globalCIDR *net.IPNet, subnets map[string]*net.IPNet) {
	refs := &resourceRefsForVPC{}
	refs.vpc = c.newResource("VPC", &cloudformation.AWSEC2VPC{
		CidrBlock:          cloudformation.NewString(globalCIDR.String()),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
	})

	refIG := c.newResource("InternetGateway", &cloudformation.AWSEC2InternetGateway{})
	c.newResource("VPCGatewayAttachment", &cloudformation.AWSEC2VPCGatewayAttachment{
		InternetGatewayId: refIG,
		VpcId:             refs.vpc,
	})

	refRT := c.newResource("RouteTable", &cloudformation.AWSEC2RouteTable{
		VpcId: refs.vpc,
	})

	c.newResource("Route", &cloudformation.AWSEC2Route{
		RouteTableId:         refRT,
		DestinationCidrBlock: cloudformation.NewString("0.0.0.0/0"),
		GatewayId:            refIG,
	})

	for az, subnet := range subnets {
		alias := strings.ToUpper(strings.Join(strings.Split(az, "-"), ""))
		refSubnet := c.newResource("Subnet"+alias, &cloudformation.AWSEC2Subnet{
			AvailabilityZone: cloudformation.NewString(az),
			CidrBlock:        cloudformation.NewString(subnet.String()),
			VpcId:            refs.vpc,
		})
		c.newResource("RouteTableAssociation"+alias, &cloudformation.AWSEC2SubnetRouteTableAssociation{
			SubnetId:     refSubnet,
			RouteTableId: refRT,
		})
		refs.subnets = append(refs.subnets, refSubnet)
	}

	refSG := c.newResource("ControlPlaneSecurityGroup", &cloudformation.AWSEC2SecurityGroup{
		GroupDescription: cloudformation.NewString("For communication between the cluster control plane and worker nodes"),
		VpcId:            refs.vpc,
	})
	refs.securityGroups = []*cloudformation.StringIntrinsic{refSG}

	c.vpcRefs = refs

	c.newOutput(OutputClusterVPC, refs.vpc, true)
	c.newJoinedOutput(OutputClusterControlPlaneSecurityGroup, refs.securityGroups, true)
	c.newJoinedOutput(OutputClusterSubnets, refs.subnets, true)
}

func (c *clusterResourceSet) addResourcesForControlPlane(version string) {
	c.newResource("ServiceRole", &cloudformation.AWSIAMRole{
		AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("eks.amazonaws.com"),
		ManagedPolicyArns: []*cloudformation.StringIntrinsic{
			cloudformation.NewString(iamAmazonEKSServicePolicyARN),
			cloudformation.NewString(iamAmazonEKSClusterPolicyARN),
		},
	})
	c.newResource("ControlPlane", &cloudformation.AWSEKSCluster{
		Name:    refStackName,
		RoleArn: cloudformation.NewStringIntrinsic("Fn::GetAtt", "ServiceRole.Arn"),
		Version: cloudformation.NewString(version),
		ResourcesVpcConfig: &cloudformation.AWSEKSCluster_ResourcesVpcConfig{
			SubnetIds:        c.vpcRefs.subnets,
			SecurityGroupIds: c.vpcRefs.securityGroups,
		},
	})

	c.newOutputFromAtt(OutputClusterCertificateAuthorityData, "ControlPlane.CertificateAuthorityData", false)
	c.newOutputFromAtt(OutputClusterEndpoint, "ControlPlane.Endpoint", true)
	c.newOutputFromAtt(OutputClusterARN, "ControlPlane.Arn", true)
}
