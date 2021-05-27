package generator

import (
	"boundedinfinity/codegen/util"

	"github.com/aymerick/raymond"
)

func init() {
	raymond.RegisterHelper("uc", util.Uc)
	raymond.RegisterHelper("ucFirst", util.UcFirst)
	raymond.RegisterHelper("ifeq", ifeq)
	raymond.RegisterHelper("basePath", util.PathBase)
	// raymond.RegisterHelper("operationId", operationId)

	// raymond.RegisterHelper("type_go", t)
	raymond.RegisterHelper("jdump", util.Jdump)
}

func (t *Generator) renderHandlebars(s string, d interface{}) (string, error) {
	o, err := raymond.Render(s, d)

	if err != nil {
		return o, err
	}

	return o, nil
}

// func operationId(path string, operation string, v model.OpenApiV310Operation) string {
// 	var operationId string

// 	if v.OperationId.IsDefined() {
// 		operationId = v.OperationId.Get()
// 	} else {

// 		operationId = strings.Join(strings.Split(path, "/"), " ")
// 		operationId = fmt.Sprintf("%v %v", operationId, strings.ToLower(operation))
// 		operationId = strutil.ToCamelCase(operationId)
// 	}

// 	return operationId
// }
