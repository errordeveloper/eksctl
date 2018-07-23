package cfn

const (
	cfnOutputNodeInstanceRoleARN = "NodeInstanceRole"

	iamPolicyAmazonEKSWorkerNodePolicyARN           = "arn:aws:iam::aws:policy/AmazonEKSWorkerNodePolicy"
	iamPolicyAmazonEKSCNIPolicyARN                  = "arn:aws:iam::aws:policy/AmazonEKS_CNI_Policy"
	iamPolicyAmazonEC2ContainerRegistryPowerUserARN = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryPowerUser"
	iamPolicyAmazonEC2ContainerRegistryReadOnlyARN  = "arn:aws:iam::aws:policy/AmazonEC2ContainerRegistryReadOnly"
)

var (
	defaultPolicyARNs = []string{
		iamPolicyAmazonEKSWorkerNodePolicyARN,
		iamPolicyAmazonEKSCNIPolicyARN,
	}
)

type nodeGroupResourceSet struct {
	resourceSet *resourceSet
}

func newNodeGroupResourceSet() *nodeGroupResourceSet {
	return &nodeGroupResourceSet{
		resourceSet: newResourceSet(),
	}
}

func (r *nodeGroupResourceSet) newResource(name string, resource interface{}) interface{} {
	return r.resourceSet.newResource(name, resource)
}

func (r *nodeGroupResourceSet) newOutput(name string, value interface{}) {
	r.resourceSet.newOutput(name, value)
}

func (r *nodeGroupResourceSet) newOutputFromAtt(name, att string) {
	r.resourceSet.newOutputFromAtt(name, att)
}

func (c *clusterResourceSet) addResourcesForNodeGroup() {
	c.newOutputFromAtt(cfnOutputNodeInstanceRoleARN, "NodeInstanceRole.Arn")
}
