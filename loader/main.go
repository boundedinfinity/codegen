package loader

import (
	"boundedinfinity/codegen/model"
)

type Loader struct {
	inputPaths         []string
	reportStack        model.StrStack
	primitiveMap       map[string]string
	inputModels        map[string]model.InputModel
	outputModels       map[string]*model.OutputModel
	dependencies       map[string]*Node
	solvedDependencies Graph
	inputOperations    map[string]model.InputOperation
	outputOperations   map[string]*model.OutputOperation
	templateMap        map[string][]model.InputTemplate
	OutputSpec         *model.OutputSpec
}

func New() *Loader {
	return &Loader{
		inputPaths:       make([]string, 0),
		primitiveMap:     make(map[string]string),
		inputModels:      make(map[string]model.InputModel),
		outputModels:     make(map[string]*model.OutputModel),
		inputOperations:  make(map[string]model.InputOperation),
		outputOperations: make(map[string]*model.OutputOperation),
		dependencies:     make(map[string]*Node),
		templateMap:      make(map[string][]model.InputTemplate),
		OutputSpec:       model.NewOutputSpec(),
	}
}

func (t *Loader) FromPaths(inputPaths []string) error {
	t.inputPaths = append(t.inputPaths, inputPaths...)

	if err := t.processInput(); err != nil {
		return err
	}

	return nil
}
