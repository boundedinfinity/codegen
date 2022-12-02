package generator

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-jsonschema/model"
)

func (t *Generator) GenerateJsonSchema(schema model.JsonSchema) error {
	renders, err := t.tm.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		title := t.dumpSchema(schema)
		fmt.Println()
		fmt.Println(title)
		fmt.Println(string(render.Output))
	}

	return nil
}

func (t *Generator) dumpSchema(schema model.JsonSchema) string {
	title := strings.Repeat("=", 5)
	title = fmt.Sprintf("%v %v %v", title, schema.GetId(), title)
	sep := strings.Repeat("=", len(title))
	return fmt.Sprintf("%v\n%v", sep, title)
}
