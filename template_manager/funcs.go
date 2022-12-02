package template_manager

import (
	"boundedinfinity/codegen/template_manager/dumper"

	"github.com/boundedinfinity/go-jsonschema/model"
)

func dumpJson(schema model.JsonSchema) string {
	return dumper.New().Dump(schema)
}
