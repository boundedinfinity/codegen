package generator

import "boundedinfinity/codegen/model"

type Generator struct {
	spec        model.OutputSpec
	reportStack model.StrStack
}

func New(spec model.OutputSpec) *Generator {
	return &Generator{
		spec: spec,
	}
}

func (t *Generator) Generate() error {
	t.reportStack.Push("generator")

	if err := t.runModels(); err != nil {
		return err
	}

	if err := t.runOperations(); err != nil {
		return err
	}

	if err := t.runNamespaces(); err != nil {
		return err
	}

	t.reportStack.Pop()
	return nil
}
