package cloudformation

// AWSSESReceiptRule_WorkmailAction AWS CloudFormation Resource (AWS::SES::ReceiptRule.WorkmailAction)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html
type AWSSESReceiptRule_WorkmailAction struct {

	// OrganizationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-organizationarn
	OrganizationArn string `json:"OrganizationArn,omitempty"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-topicarn
	TopicArn string `json:"TopicArn,omitempty"`
}

type UntypedAWSSESReceiptRule_WorkmailAction struct {

	// OrganizationArn AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-organizationarn
	OrganizationArn interface{} `json:"OrganizationArn,omitempty"`

	// TopicArn AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-ses-receiptrule-workmailaction.html#cfn-ses-receiptrule-workmailaction-topicarn
	TopicArn interface{} `json:"TopicArn,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSSESReceiptRule_WorkmailAction) AWSCloudFormationType() string {
	return "AWS::SES::ReceiptRule.WorkmailAction"
}
