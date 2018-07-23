package cloudformation

import (
	"encoding/json"
	"errors"
	"fmt"
)

// AWSServerlessFunction AWS CloudFormation Resource (AWS::Serverless::Function)
// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
type AWSServerlessFunction struct {

	// CodeUri AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	CodeUri *AWSServerlessFunction_CodeUri `json:"CodeUri,omitempty"`

	// DeadLetterQueue AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	DeadLetterQueue *AWSServerlessFunction_DeadLetterQueue `json:"DeadLetterQueue,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Description string `json:"Description,omitempty"`

	// Environment AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Environment *AWSServerlessFunction_FunctionEnvironment `json:"Environment,omitempty"`

	// Events AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Events map[string]AWSServerlessFunction_EventSource `json:"Events,omitempty"`

	// FunctionName AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	FunctionName string `json:"FunctionName,omitempty"`

	// Handler AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Handler string `json:"Handler,omitempty"`

	// KmsKeyArn AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	KmsKeyArn string `json:"KmsKeyArn,omitempty"`

	// MemorySize AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	MemorySize int `json:"MemorySize,omitempty"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Policies *AWSServerlessFunction_Policies `json:"Policies,omitempty"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Role string `json:"Role,omitempty"`

	// Runtime AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Runtime string `json:"Runtime,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tags map[string]string `json:"Tags,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Timeout int `json:"Timeout,omitempty"`

	// Tracing AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tracing string `json:"Tracing,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	VpcConfig *AWSServerlessFunction_VpcConfig `json:"VpcConfig,omitempty"`
}

type UntypedAWSServerlessFunction struct {

	// CodeUri AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	CodeUri *AWSServerlessFunction_CodeUri `json:"CodeUri,omitempty"`

	// DeadLetterQueue AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	DeadLetterQueue *UntypedAWSServerlessFunction_DeadLetterQueue `json:"DeadLetterQueue,omitempty"`

	// Description AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Description interface{} `json:"Description,omitempty"`

	// Environment AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Environment *UntypedAWSServerlessFunction_FunctionEnvironment `json:"Environment,omitempty"`

	// Events AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Events map[string]UntypedAWSServerlessFunction_EventSource `json:"Events,omitempty"`

	// FunctionName AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	FunctionName interface{} `json:"FunctionName,omitempty"`

	// Handler AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Handler interface{} `json:"Handler,omitempty"`

	// KmsKeyArn AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	KmsKeyArn interface{} `json:"KmsKeyArn,omitempty"`

	// MemorySize AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	MemorySize interface{} `json:"MemorySize,omitempty"`

	// Policies AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Policies *AWSServerlessFunction_Policies `json:"Policies,omitempty"`

	// Role AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Role interface{} `json:"Role,omitempty"`

	// Runtime AWS CloudFormation Property
	// Required: true
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Runtime interface{} `json:"Runtime,omitempty"`

	// Tags AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tags map[string]interface{} `json:"Tags,omitempty"`

	// Timeout AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Timeout interface{} `json:"Timeout,omitempty"`

	// Tracing AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	Tracing interface{} `json:"Tracing,omitempty"`

	// VpcConfig AWS CloudFormation Property
	// Required: false
	// See: https://github.com/awslabs/serverless-application-model/blob/master/versions/2016-10-31.md#awsserverlessfunction
	VpcConfig *UntypedAWSServerlessFunction_VpcConfig `json:"VpcConfig,omitempty"`
}

// AWSCloudFormationType returns the AWS CloudFormation resource type
func (r *AWSServerlessFunction) AWSCloudFormationType() string {
	return "AWS::Serverless::Function"
}

// MarshalJSON is a custom JSON marshalling hook that embeds this object into
// an AWS CloudFormation JSON resource's 'Properties' field and adds a 'Type'.
func (r *AWSServerlessFunction) MarshalJSON() ([]byte, error) {
	type Properties AWSServerlessFunction
	return json.Marshal(&struct {
		Type       string
		Properties Properties
	}{
		Type:       r.AWSCloudFormationType(),
		Properties: (Properties)(*r),
	})
}

// UnmarshalJSON is a custom JSON unmarshalling hook that strips the outer
// AWS CloudFormation resource object, and just keeps the 'Properties' field.
func (r *AWSServerlessFunction) UnmarshalJSON(b []byte) error {
	type Properties AWSServerlessFunction
	res := &struct {
		Type       string
		Properties *Properties
	}{}
	if err := json.Unmarshal(b, &res); err != nil {
		fmt.Printf("ERROR: %s\n", err)
		return err
	}

	// If the resource has no Properties set, it could be nil
	if res.Properties != nil {
		*r = AWSServerlessFunction(*res.Properties)
	}

	return nil
}

// GetAllAWSServerlessFunctionResources retrieves all AWSServerlessFunction items from an AWS CloudFormation template
func (t *Template) GetAllAWSServerlessFunctionResources() map[string]AWSServerlessFunction {
	results := map[string]AWSServerlessFunction{}
	for name, untyped := range t.Resources {
		switch resource := untyped.(type) {
		case AWSServerlessFunction:
			// We found a strongly typed resource of the correct type; use it
			results[name] = resource
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Serverless::Function" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServerlessFunction
						if err := json.Unmarshal(b, &result); err == nil {
							results[name] = result
						}
					}
				}
			}
		}
	}
	return results
}

// GetAWSServerlessFunctionWithName retrieves all AWSServerlessFunction items from an AWS CloudFormation template
// whose logical ID matches the provided name. Returns an error if not found.
func (t *Template) GetAWSServerlessFunctionWithName(name string) (AWSServerlessFunction, error) {
	if untyped, ok := t.Resources[name]; ok {
		switch resource := untyped.(type) {
		case AWSServerlessFunction:
			// We found a strongly typed resource of the correct type; use it
			return resource, nil
		case map[string]interface{}:
			// We found an untyped resource (likely from JSON) which *might* be
			// the correct type, but we need to check it's 'Type' field
			if resType, ok := resource["Type"]; ok {
				if resType == "AWS::Serverless::Function" {
					// The resource is correct, unmarshal it into the results
					if b, err := json.Marshal(resource); err == nil {
						var result AWSServerlessFunction
						if err := json.Unmarshal(b, &result); err == nil {
							return result, nil
						}
					}
				}
			}
		}
	}
	return AWSServerlessFunction{}, errors.New("resource not found")
}
