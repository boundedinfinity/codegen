package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"strings"

	"github.com/aymerick/raymond"
	"github.com/ozgio/strutil"
)

func init() {
	raymond.RegisterHelper("uc", uc)
	raymond.RegisterHelper("ucFirst", ucFirst)
	raymond.RegisterHelper("ifeq", ifeq)
	raymond.RegisterHelper("basePath", basePath)
	raymond.RegisterHelper("operationId", operationId)

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

func operationId(path string, operation string, v model.OpenApiV310Operation) string {
	var operationId string

	if v.OperationId.IsDefined() {
		operationId = v.OperationId.Get()
	} else {

		operationId = strings.Join(strings.Split(path, "/"), " ")
		operationId = fmt.Sprintf("%v %v", operationId, strings.ToLower(operation))
		operationId = strutil.ToCamelCase(operationId)
	}

	return operationId
}