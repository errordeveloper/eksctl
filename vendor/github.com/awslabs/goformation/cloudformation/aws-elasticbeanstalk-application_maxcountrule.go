package cloudformation

// AWSElasticBeanstalkApplication_MaxCountRule AWS CloudFormation Resource (AWS::ElasticBeanstalk::Application.MaxCountRule)
// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html
type AWSElasticBeanstalkApplication_MaxCountRule struct {

	// DeleteSourceFromS3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-deletesourcefroms3
	DeleteSourceFromS3 bool `json:"DeleteSourceFromS3,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-enabled
	Enabled bool `json:"Enabled,omitempty"`

	// MaxCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-maxcount
	MaxCount int `json:"MaxCount,omitempty"`
}

type UntypedAWSElasticBeanstalkApplication_MaxCountRule struct {

	// DeleteSourceFromS3 AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-deletesourcefroms3
	DeleteSourceFromS3 interface{} `json:"DeleteSourceFromS3,omitempty"`

	// Enabled AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-enabled
	Enabled interface{} `json:"Enabled,omitempty"`

	// MaxCount AWS CloudFormation Property
	// Required: false
	// See: http://docs.aws.amazon.com/AWSCloudFormation/latest/UserGuide/aws-properties-elasticbeanstalk-application-maxcountrule.html#cfn-elasticbeanstalk-application-maxcountrule-maxcount
	MaxCount interface{} `json:"MaxCount,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSElasticBeanstalkApplication_MaxCountRule) AWSCloudFormationType() string {
	return "AWS::ElasticBeanstalk::Application.MaxCountRule"
}
