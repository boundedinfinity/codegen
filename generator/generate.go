package generator

import (
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/renderer"
	"boundedinfinity/codegen/util"
	"os"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-mimetyper/file_extention"
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
	renders, err := t.renderer.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		var outputExt string
		var output string

		outputExts, err := file_extention.GetExts(render.OutputMimeType)
		if err != nil {
			return err
		}

		inputExts, err := file_extention.GetExts(schema.Base().MimeType)

		if err != nil {
			return err
		}

		if len(outputExts) > 0 {
			outputExt = outputExts[0].String()
		}

		output = schema.Base().Source
		output = strings.Replace(
			output,
			schema.Base().Root,
			t.projectManager.Merged.Info.DestDir.Get(),
			1,
		)

		output = util.RemoveSchema(output)

		for _, inputExt := range inputExts {
			output = extentioner.Swap(output, inputExt.String(), outputExt)
		}

		render.OutputPath = output

		if err := t.writeModel(render); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) GenerateOperation(schema render_context.RenderContextOperation) error {
	// _, err := t.tm.RenderOperation(schema)

	// if err != nil {
	// 	return err
	// }

	// for _, render := range renders {
	// 	fmt.Println(render.Schema.Base().OutputPath)

	// 	if err := t.writeModel(render); err != nil {
	// 		return err
	// 	}
	// }

	return nil
}

func (t *Generator) writeModel(output renderer.ModelOutput) error {
	path := output.OutputPath

	if err := pather.DirEnsure(filepath.Dir(path)); err != nil {
		return err
	}

	if err := os.WriteFile(path, output.Output, t.fileMode); err != nil {
		return err
	}

	return nil
}
