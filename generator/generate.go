package generator

import (
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_manager"
	"fmt"
	"os"
	"path/filepath"

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

func (t *Generator) GenerateOperation(schema render_context.RenderContextOperation) error {
	_, err := t.tm.RenderOperation(schema)

	if err != nil {
		return err
	}

	// for _, render := range renders {
	// 	fmt.Println(render.Schema.Base().OutputPath)

	// 	if err := t.writeModel(render); err != nil {
	// 		return err
	// 	}
	// }

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
