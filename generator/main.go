package generator

import "boundedinfinity/codegen/model"

type Generator struct {
	spec model.BiGen
}

func New(spec model.BiGen) *Generator {
	return &Generator{
		spec: spec,
	}
}

func (t *Generator) Generate() error {
	if err := t.run(); err != nil {
		return err
	}

	return nil
}

func (t *Generator) run() error {

	for _, ns := range t.spec.Models.Namespaces {
		if err := t.runNamespace(ns); err != nil {
			return err
		}
	}

	return nil
}
