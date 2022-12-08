package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
)

func (t *Generator) GenerateModel(schema canonical.Canonical) error {
	renders, err := t.tm.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		title := t.dumpSchema(schema)
		fmt.Println()
		fmt.Println(title)
		fmt.Println(string(render.Output))

		if err := t.writeModel(render); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) writeModel(output template_manager.ModelOutput) error {
	filename := output.Path
	filename = filepath.Base(filename)
	filename = extentioner.Strip(filename)
	path := filepath.Join(t.destDir, t.codeGenSchema.Info.RootDir.Get())

	fmt.Println(path)

	return nil
}

func (t *Generator) dumpSchema(schema canonical.Canonical) string {
	title := strings.Repeat("=", 5)
	title = fmt.Sprintf("%v %v %v", title, schema.SchemaId(), title)
	sep := strings.Repeat("=", len(title))
	return fmt.Sprintf("%v\n%v", sep, title)
}
