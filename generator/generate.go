package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_manager"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *Generator) Generate() error {
	for _, rc := range t.rcs {
		if err := t.GenerateModel(rc); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) GenerateModel(schema render_context.RenderContext) error {
	renders, err := t.tm.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		fmt.Println(render.Schema.Base().OutputPath)

		if err := t.writeModel(render); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) writeModel(output template_manager.ModelOutput) error {
	path := output.Schema.Base().OutputPath

	if err := pather.DirEnsure(filepath.Dir(path)); err != nil {
		return err
	}

	if err := os.WriteFile(path, output.Output, t.fileMode); err != nil {
		return err
	}

	return nil
}

func (t *Generator) dumpSchema(schema canonical.Canonical) string {
	title := strings.Repeat("=", 5)
	title = fmt.Sprintf("%v %v %v", title, schema.SchemaId(), title)
	sep := strings.Repeat("=", len(title))
	return fmt.Sprintf("%v\n%v", sep, title)
}
