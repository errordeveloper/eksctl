package builder

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/pkg/errors"

	cfn "github.com/aws/aws-sdk-go/service/cloudformation"
	gfn "github.com/awslabs/goformation/cloudformation"

	"github.com/kubicorn/kubicorn/pkg/logger"
)

// func newParameter(t *gfn.Template, name, valueType, defaultValue string) interface{} {
// 	p := map[string]string{"Type": valueType}
// 	if defaultValue != "" {
// 		p["Default"] = defaultValue
// 	}
// 	t.Parameters[name] = p
// 	return makeRef(name)
// }

// func newStringParameter(t *gfn.Template, name, defaultValue string) interface{} {
// 	return newParameter(t, name, "String", defaultValue)
// }

// func newSub(sub string) interface{} {
// 	return map[string]string{"Sub": sub}
// }

type resourceSet struct {
	template *gfn.Template
	outputs  []string
}

func makeRef(refName string) *gfn.StringIntrinsic {
	return gfn.NewStringRef(refName)
}

var refStackName = makeRef("AWS::StackName")

func (r *resourceSet) newResource(name string, resource interface{}) *gfn.StringIntrinsic {
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
	r.outputs = append(r.outputs, name)
}

func (r *resourceSet) newJoinedOutput(name string, value []*gfn.StringIntrinsic, export bool) {
	r.newOutput(name, map[string][]interface{}{"Fn::Join": []interface{}{",", value}}, export)
}

func (r *resourceSet) newOutputFromAtt(name, att string, export bool) {
	r.newOutput(name, map[string]string{"Fn::GetAtt": att}, export)
}

func getOutput(stack *cfn.Stack, key string) *string {
	for _, x := range stack.Outputs {
		if *x.OutputKey == key {
			return x.OutputValue
		}
	}
	return nil
}

func setOutput(obj interface{}, key, value string) error {
	e := reflect.ValueOf(obj).Elem()
	if e.Kind() == reflect.Struct {
		f := e.FieldByName(key)
		if f.IsValid() && f.CanSet() {
			switch f.Kind() {
			case reflect.String:
				f.SetString(value)
			case reflect.Slice:
				if f.Type().String() == "[]string" {
					f.Set(reflect.ValueOf(strings.Split(value, ",")))
				}
			default:
				return fmt.Errorf("unexpected type %q of destination field for %q", f.Kind(), key)
			}
		} else {
			return fmt.Errorf("cannot set destination field for %q", key)
		}
	} else {
		return fmt.Errorf("cannot use destination interface of type %q", e.Kind())
	}
	return nil
}

func (r *resourceSet) GetAllOutputs(stackChan chan cfn.Stack, obj interface{}) error {
	defer close(stackChan)

	logger.Debug("processing stack outputs")

	stack := <-stackChan

	for _, key := range r.outputs {
		value := getOutput(&stack, key)
		if value == nil {
			return fmt.Errorf("%s is nil", key)
		}
		if err := setOutput(obj, key, *value); err != nil {
			return errors.Wrap(err, "processing stack outputs")
		}
	}
	return nil
}

func newResourceSet() *resourceSet {
	return &resourceSet{
		template: gfn.NewTemplate(),
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
