package cloudformation

// AWSWAFRegionalXssMatchSet_FieldToMatch AWS CloudFormation Resource (AWS::WAFRegional::XssMatchSet.FieldToMatch)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html
type AWSWAFRegionalXssMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-data
	Data string `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-type
	Type string `json:"Type,omitempty"`
}

type UntypedAWSWAFRegionalXssMatchSet_FieldToMatch struct {

	// Data AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-data
	Data interface{} `json:"Data,omitempty"`

	// Type AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-wafregional-xssmatchset-fieldtomatch.html#cfn-wafregional-xssmatchset-fieldtomatch-type
	Type interface{} `json:"Type,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSWAFRegionalXssMatchSet_FieldToMatch) AWSCloudFormationType() string {
	return "AWS::WAFRegional::XssMatchSet.FieldToMatch"
}
