package cloudformation

// AWSConfigConfigRule_Source AWS CloudFormation Resource (AWS::Config::ConfigRule.Source)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html
type AWSConfigConfigRule_Source struct {

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner
	Owner string `json:"Owner,omitempty"`

	// SourceDetails AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails
	SourceDetails []AWSConfigConfigRule_SourceDetail `json:"SourceDetails,omitempty"`

	// SourceIdentifier AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier
	SourceIdentifier string `json:"SourceIdentifier,omitempty"`
}

type UntypedAWSConfigConfigRule_Source struct {

	// Owner AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-owner
	Owner interface{} `json:"Owner,omitempty"`

	// SourceDetails AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourcedetails
	SourceDetails []UntypedAWSConfigConfigRule_SourceDetail `json:"SourceDetails,omitempty"`

	// SourceIdentifier AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-config-configrule-source.html#cfn-config-configrule-source-sourceidentifier
	SourceIdentifier interface{} `json:"SourceIdentifier,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSConfigConfigRule_Source) AWSCloudFormationType() string {
	return "AWS::Config::ConfigRule.Source"
}
