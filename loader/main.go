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
	OutputSpec         *model.OutputSpec
}

func (t *Loader) appendInfoTemplate(v model.InputTemplate) {
	list, ok := t.inputTemplates[v.Namespace]

	if !ok {
		list = make([]model.InputTemplate, 0)
	}

	var found bool

	for _, x := range list {
		if x.Path == v.Path {
			found = true
			break
		}
	}

	if !found {
		list = append(list, v)
	}

	t.inputTemplates[v.Namespace] = list
}

func New() *Loader {
	return &Loader{
		inputPaths:       make([]string, 0),
		inputModels:      make(map[string]model.InputModel),
		outputModels:     make(map[string]*model.OutputModel),
		inputOperations:  make(map[string]model.InputOperation),
		outputOperations: make(map[string]*model.OutputOperation),
		inputTemplates:   make(map[string][]model.InputTemplate),
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
