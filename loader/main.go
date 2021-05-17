package loader

import (
	"boundedinfinity/codegen/model"
)

type Loader struct {
	inputPaths         []string
	reportStack        model.StrStack
	primitiveMap       map[string]string
	inputModels        map[string]model.InputModel
	dependencies       map[string]*Node
	solvedDependencies Graph
	modelMap           map[string]*model.OutputModel
	operationMap       map[string]*model.OutputOperation
	templateMap        map[string][]model.InputTemplate
	OutputSpec         *model.OutputSpec
}

func New() *Loader {
	return &Loader{
		inputPaths:   make([]string, 0),
		primitiveMap: make(map[string]string),
		inputModels:  make(map[string]model.InputModel),
		modelMap:     make(map[string]*model.OutputModel),
		operationMap: make(map[string]*model.OutputOperation),
		dependencies: make(map[string]*Node),
		templateMap:  make(map[string][]model.InputTemplate),
		OutputSpec:   model.NewOutputSpec(),
	}
}

func (t *Loader) FromPaths(inputPaths []string) error {
	t.inputPaths = append(t.inputPaths, inputPaths...)

	if err := t.processInput(); err != nil {
		return err
	}

	return nil
}
