package generator

import (
	"boundedinfinity/codegen/model"
)

type Generator struct {
	specPath     string
	templateRoot string
	genRoot      string
	spec         model.BiSpec
}

func New(specPath string) *Generator {
	return &Generator{
		specPath: specPath,
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
