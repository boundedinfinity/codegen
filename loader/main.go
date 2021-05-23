package loader

import (
	"boundedinfinity/codegen/model"
)

type Loader struct {
	inputPaths         []string
	reportStack        model.StrStack
	inputModels        map[string]model.InputModel
	outputModels       map[string]*model.OutputModel
	dependencies       map[string]*Node
	solvedDependencies Graph
	inputOperations    map[string]model.InputOperation
	outputOperations   map[string]*model.OutputOperation
	inputTemplates     map[string][]model.InputTemplate
	outputNamespace    map[string]*model.OutputNamespace
	OutputSpec         *model.OutputSpec
}

func New() *Loader {
	return &Loader{
		inputPaths:       make([]string, 0),
		inputModels:      make(map[string]model.InputModel),
		outputModels:     make(map[string]*model.OutputModel),
		inputOperations:  make(map[string]model.InputOperation),
		outputOperations: make(map[string]*model.OutputOperation),
		inputTemplates:   make(map[string][]model.InputTemplate),
		outputNamespace:  make(map[string]*model.OutputNamespace),
		dependencies:     make(map[string]*Node),
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
