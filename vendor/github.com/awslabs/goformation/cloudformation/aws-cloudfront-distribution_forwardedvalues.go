package cloudformation

// AWSCloudFrontDistribution_ForwardedValues AWS CloudFormation Resource (AWS::CloudFront::Distribution.ForwardedValues)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html
type AWSCloudFrontDistribution_ForwardedValues struct {

	// Cookies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-cookies
	Cookies *AWSCloudFrontDistribution_Cookies `json:"Cookies,omitempty"`

	// Headers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-headers
	Headers []string `json:"Headers,omitempty"`

	// QueryString AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-querystring
	QueryString bool `json:"QueryString,omitempty"`

	// QueryStringCacheKeys AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-querystringcachekeys
	QueryStringCacheKeys []string `json:"QueryStringCacheKeys,omitempty"`
}

type UntypedAWSCloudFrontDistribution_ForwardedValues struct {

	// Cookies AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-cookies
	Cookies *UntypedAWSCloudFrontDistribution_Cookies `json:"Cookies,omitempty"`

	// Headers AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-headers
	Headers []interface{} `json:"Headers,omitempty"`

	// QueryString AWS CloudFormation Property
	// Required: true
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-querystring
	QueryString interface{} `json:"QueryString,omitempty"`

	// QueryStringCacheKeys AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-cloudfront-distribution-forwardedvalues.html#cfn-cloudfront-distribution-forwardedvalues-querystringcachekeys
	QueryStringCacheKeys []interface{} `json:"QueryStringCacheKeys,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSCloudFrontDistribution_ForwardedValues) AWSCloudFormationType() string {
	return "AWS::CloudFront::Distribution.ForwardedValues"
}
