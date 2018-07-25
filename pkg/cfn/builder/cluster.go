package builder

import (
	"net"
	"strings"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	gfn "github.com/awslabs/goformation/cloudformation"
)

const (
	cfnOutputClusterCertificateAuthorityData = "CertificateAuthorityData"
	cfnOutputClusterEndpoint                 = "Endpoint"
	cfnOutputClusterARN                      = "ARN"
	cfnOutputClusterVPC                      = "VPC"
	cfnOutputClusterSubnets                  = "Subnets"
	cfnOutputClusterSecurityGroup            = "SecurityGroup"

	iamAmazonEKSServicePolicyARN = "arn:aws:iam::aws:policy/AmazonEKSServicePolicy"
	iamAmazonEKSClusterPolicyARN = "arn:aws:iam::aws:policy/AmazonEKSClusterPolicy"
)

type clusterResourceSet struct {
	resourceSet      *resourceSet
	vpcRefs          *resourceRefsForVPC
	controlPlaneRefs *resourceRefsForControlPlane
}

type resourceRefsForVPC struct {
	vpc            *gfn.StringIntrinsic
	subnets        []*gfn.StringIntrinsic
	securityGroups []*gfn.StringIntrinsic
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

func (c *clusterResourceSet) newResource(name string, resource interface{}) *gfn.StringIntrinsic {
	return c.resourceSet.newResource(name, resource)
}

func (c *clusterResourceSet) newOutput(name string, value interface{}, export bool) {
	c.resourceSet.newOutput(name, value, export)
}

func (c *clusterResourceSet) newJoinedOutput(name string, value []*gfn.StringIntrinsic, export bool) {
	c.resourceSet.newJoinedOutput(name, value, export)
}

func (c *clusterResourceSet) newOutputFromAtt(name, att string, export bool) {
	c.resourceSet.newOutputFromAtt(name, att, export)
}

func (c *clusterResourceSet) addResourcesForVPC(globalCIDR *net.IPNet, subnets map[string]*net.IPNet) {
	refs := &resourceRefsForVPC{}
	refs.vpc = c.newResource("VPC", &gfn.AWSEC2VPC{
		CidrBlock:          gfn.NewString(globalCIDR.String()),
		EnableDnsSupport:   true,
		EnableDnsHostnames: true,
	})

	refIG := c.newResource("InternetGateway", &gfn.AWSEC2InternetGateway{})
	c.newResource("VPCGatewayAttachment", &gfn.AWSEC2VPCGatewayAttachment{
		InternetGatewayId: refIG,
		VpcId:             refs.vpc,
	})

	refRT := c.newResource("RouteTable", &gfn.AWSEC2RouteTable{
		VpcId: refs.vpc,
	})

	c.newResource("Route", &gfn.AWSEC2Route{
		RouteTableId:         refRT,
		DestinationCidrBlock: gfn.NewString("0.0.0.0/0"),
		GatewayId:            refIG,
	})

	for az, subnet := range subnets {
		alias := strings.ToUpper(strings.Join(strings.Split(az, "-"), ""))
		refSubnet := c.newResource("Subnet"+alias, &gfn.AWSEC2Subnet{
			AvailabilityZone: gfn.NewString(az),
			CidrBlock:        gfn.NewString(subnet.String()),
			VpcId:            refs.vpc,
		})
		c.newResource("RouteTableAssociation"+alias, &gfn.AWSEC2SubnetRouteTableAssociation{
			SubnetId:     refSubnet,
			RouteTableId: refRT,
		})
		refs.subnets = append(refs.subnets, refSubnet)
	}

	refSG := c.newResource("ControlPlaneSecurityGroup", &gfn.AWSEC2SecurityGroup{
		GroupDescription: gfn.NewString("For communication between the cluster control plane and worker nodes"),
		VpcId:            refs.vpc,
	})
	refs.securityGroups = []*gfn.StringIntrinsic{refSG}

	c.vpcRefs = refs

	c.newOutput(cfnOutputClusterVPC, refs.vpc, true)
	c.newJoinedOutput(cfnOutputClusterSecurityGroup, refs.securityGroups, true)
	c.newJoinedOutput(cfnOutputClusterSubnets, refs.subnets, true)
}

func (c *clusterResourceSet) addResourcesForControlPlane(version string) {
	c.newResource("ServiceRole", &gfn.AWSIAMRole{
		AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("eks.amazonaws.com"),
		ManagedPolicyArns: []*gfn.StringIntrinsic{
			gfn.NewString(iamAmazonEKSServicePolicyARN),
			gfn.NewString(iamAmazonEKSClusterPolicyARN),
		},
	})
	c.newResource("ControlPlane", &gfn.AWSEKSCluster{
		Name:    refStackName,
		RoleArn: gfn.NewStringIntrinsic("Fn::GetAtt", "ServiceRole.Arn"),
		Version: gfn.NewString(version),
		ResourcesVpcConfig: &gfn.AWSEKSCluster_ResourcesVpcConfig{
			SubnetIds:        c.vpcRefs.subnets,
			SecurityGroupIds: c.vpcRefs.securityGroups,
		},
	})

	c.newOutputFromAtt(cfnOutputClusterCertificateAuthorityData, "ControlPlane.CertificateAuthorityData", false)
	c.newOutputFromAtt(cfnOutputClusterEndpoint, "ControlPlane.Endpoint", true)
	c.newOutputFromAtt(cfnOutputClusterARN, "ControlPlane.Arn", true)
}

func (c *clusterResourceSet) GetAllOutputs(stackChan chan cfn.Stack, obj interface{}) error {
	return c.resourceSet.GetAllOutputs(stackChan, obj)
}
