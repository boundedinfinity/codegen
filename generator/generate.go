package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/renderer"
	"boundedinfinity/codegen/util"
	"os"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *Generator) Generate() error {
	merged := t.projectManager.Merged

	for _, typ := range merged.Types {
		if err := t.GenerateModel(typ); err != nil {
			return err
		}
	}

	for _, operation := range merged.Operations {
		if err := t.GenerateOperation(*operation); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) GenerateModel(schema ct.CodeGenType) error {
	renders, err := t.renderer.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		if outputPath := util.GetOutputPath(t.projectManager.Merged.Info.DestDir.Get(), render.CodeGenProjectTemplateFile, render.Schema); outputPath.Failure() {
			return outputPath.Error
		} else {
			render.OutputPath = outputPath.Result
		}

		if err := t.writeModel(render); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) GenerateOperation(schema ct.CodeGenProjectOperation) error {

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
