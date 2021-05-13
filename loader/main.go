package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
)

type Loader struct {
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
	for _, p := range util.SchemaTypePrimitives {
		t.primitiveMap[string(p)] = ""
		t.dependencies[string(p)] = NewNode(string(p))
	}

	for _, inputPath := range inputPaths {
		if err := t.combine(inputPath); err != nil {
			return err
		}
	}

	if err := t.processInput(); err != nil {
		return err
	}

	return nil
}
