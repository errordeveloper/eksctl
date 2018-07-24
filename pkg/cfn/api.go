package cfn

import (
	"github.com/awslabs/goformation/cloudformation"
)

// func newParameter(t *cloudformation.Template, name, valueType, defaultValue string) interface{} {
// 	p := map[string]string{"Type": valueType}
// 	if defaultValue != "" {
// 		p["Default"] = defaultValue
// 	}
// 	t.Parameters[name] = p
// 	return makeRef(name)
// }

// func newStringParameter(t *cloudformation.Template, name, defaultValue string) interface{} {
// 	return newParameter(t, name, "String", defaultValue)
// }

// func newSub(sub string) interface{} {
// 	return map[string]string{"Sub": sub}
// }

type resourceSet struct {
	template *cloudformation.Template
}

func makeRef(refName string) *cloudformation.StringIntrinsic {
	return cloudformation.NewStringRef(refName)
}

var refStackName = makeRef("AWS::StackName")

func (r *resourceSet) newResource(name string, resource interface{}) *cloudformation.StringIntrinsic {
	r.template.Resources[name] = resource
	return makeRef(name)
}

func (r *resourceSet) renderJSON() ([]byte, error) {
	return r.template.JSON()
}

func (r *resourceSet) newOutput(name string, value interface{}, export bool) {
	o := map[string]interface{}{"Value": value}
	if export {
		o["Export"] = map[string]map[string]string{
			"Name": {"Fn::Sub": "${AWS::StackName}::" + name},
		}
	}
	r.template.Outputs[name] = o
}

func (r *resourceSet) newJoinedOutput(name string, value []*cloudformation.StringIntrinsic, export bool) {
	r.newOutput(name, map[string][]interface{}{"Fn::Join": []interface{}{",", value}}, export)
}

func (r *resourceSet) newOutputFromAtt(name, att string, export bool) {
	r.newOutput(name, map[string]string{"Fn::GetAtt": att}, export)
}

func newResourceSet() *resourceSet {
	return &resourceSet{
		template: cloudformation.NewTemplate(),
	}
}

func makeAssumeRolePolicyDocument(service string) map[string]interface{} {
	return map[string]interface{}{
		"Version": "2012-10-17",
		"Statement": []interface{}{
			map[string]interface{}{
				"Effect": "Allow",
				"Principal": map[string][]string{
					"Service": []string{service},
				},
				"Action": []string{"sts:AssumeRole"},
			},
		},
	}
}
