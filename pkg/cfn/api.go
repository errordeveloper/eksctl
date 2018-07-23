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

func makeRef(refName string) interface{} {
	return map[string]string{"Ref": refName}
}

func (r *resourceSet) newResource(name string, resource interface{}) interface{} {
	r.template.Resources[name] = resource
	return makeRef(name)
}

func (r *resourceSet) newOutput(name string, value interface{}) {
	o := map[string]interface{}{"Value": value}
	r.template.Outputs[name] = o
}

func (r *resourceSet) newOutputFromAtt(name, att string) {
	r.newOutput(name, map[string]string{"Fn::GetAtt": att})
}

func newResourceSet() *resourceSet {
	return &resourceSet{
		template: cloudformation.NewTemplate(),
	}
}
