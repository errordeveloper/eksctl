package cfn

import "github.com/awslabs/goformation/cloudformation"

const (
	cfnOutputNodeInstanceRoleARN = "NodeInstanceRole"

	iamPolicyAmazonEKSWorkerNodePolicyARN           = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
	iamPolicyAmazonEKSCNIPolicyARN                  = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
	iamPolicyAmazonEC2ContainerRegistryPowerUserARN = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
	iamPolicyAmazonEC2ContainerRegistryReadOnlyARN  = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
)

var (
	defaultPolicyARNs = []*cloudformation.StringIntrinsic{
		cloudformation.NewString(iamPolicyAmazonEKSWorkerNodePolicyARN),
		cloudformation.NewString(iamPolicyAmazonEKSCNIPolicyARN),
	}

	maxPodsPerNode = map[string]int{
		"c4.large":    29,
		"c4.xlarge":   58,
		"c4.2xlarge":  58,
		"c4.4xlarge":  234,
		"c4.8xlarge":  234,
		"c5.large":    29,
		"c5.xlarge":   58,
		"c5.2xlarge":  58,
		"c5.4xlarge":  234,
		"c5.9xlarge":  234,
		"c5.18xlarge": 737,
		"i3.large":    29,
		"i3.xlarge":   58,
		"i3.2xlarge":  58,
		"i3.4xlarge":  234,
		"i3.8xlarge":  234,
		"i3.16xlarge": 737,
		"m3.medium":   12,
		"m3.large":    29,
		"m3.xlarge":   58,
		"m3.2xlarge":  118,
		"m4.large":    20,
		"m4.xlarge":   58,
		"m4.2xlarge":  58,
		"m4.4xlarge":  234,
		"m4.10xlarge": 234,
		"m5.large":    29,
		"m5.xlarge":   58,
		"m5.2xlarge":  58,
		"m5.4xlarge":  234,
		"m5.12xlarge": 234,
		"m5.24xlarge": 737,
		"p2.xlarge":   58,
		"p2.8xlarge":  234,
		"p2.16xlarge": 234,
		"p3.2xlarge":  58,
		"p3.8xlarge":  234,
		"p3.16xlarge": 234,
		"r3.xlarge":   58,
		"r3.2xlarge":  58,
		"r3.4xlarge":  234,
		"r3.8xlarge":  234,
		"r4.large":    29,
		"r4.xlarge":   58,
		"r4.2xlarge":  58,
		"r4.4xlarge":  234,
		"r4.8xlarge":  234,
		"r4.16xlarge": 737,
		"t2.small":    8,
		"t2.medium":   17,
		"t2.large":    35,
		"t2.xlarge":   44,
		"t2.2xlarge":  44,
		"x1.16xlarge": 234,
		"x1.32xlarge": 234,
	}
)

type nodeGroupResourceSet struct {
	resourceSet *resourceSet
	vpc         *resourceRefsForVPC
}

func newNodeGroupResourceSet() *nodeGroupResourceSet {
	return &nodeGroupResourceSet{
		resourceSet: newResourceSet(),
	}
}

func (n *nodeGroupResourceSet) newResource(name string, resource interface{}) *cloudformation.StringIntrinsic {
	return n.resourceSet.newResource(name, resource)
}

func (n *nodeGroupResourceSet) newOutputFromAtt(name, att string, export bool) {
	n.resourceSet.newOutputFromAtt(name, att, export)
}

func (n *nodeGroupResourceSet) addResourcesForNodeGroup() {
	n.newResource("NodeInstanceProfile", &cloudformation.AWSIAMInstanceProfile{
		Path: cloudformation.NewString("/"),
		Roles: []*cloudformation.StringIntrinsic{
			n.newResource("NodeInstanceRole", &cloudformation.AWSIAMRole{
				Path: cloudformation.NewString("/"),
				AssumeRolePolicyDocument: makeAssumeRolePolicyDocument("ec2.amazonaws.com"),
				ManagedPolicyArns:        defaultPolicyARNs, // TODO parametrise
			}),
		},
	})

	n.newResource("NodeSecurityGroup", &cloudformation.AWSEC2SecurityGroup{
		VpcId: n.vpc.vpc,
	})

	n.newOutputFromAtt(cfnOutputNodeInstanceRoleARN, "NodeInstanceRole.Arn", false)
}
