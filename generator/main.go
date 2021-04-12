package generator

import (
	"boundedinfinity/codegen/model"
)

type Generator struct {
	inputPath    string
	templateRoot string
	genRoot      string
	model        model.OpenApiV310
}

func New(inputPath string) *Generator {
	return &Generator{
		inputPath: inputPath,
	}
}

func (t *Generator) Generate() error {
	if err := t.load(); err != nil {
		return err
	}

	if err := t.run(); err != nil {
		return err
	}

	return nil
}
