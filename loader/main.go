package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
)

type Loader struct {
	inputPath      string
	inputSpec      model.InputSpec
	namespaceStack model.StrStack
	reportStack    model.StrStack
	dependencies   map[string]*Node
	namespaceMap   map[string]*model.OutputNamespace
	modelMap       map[string]*model.OutputModel
	operationMap   map[string]*model.OutputOperation
	templateMap    map[string][]model.InputTemplate
	OutputSpec     *model.OutputSpec
}

func New() *Loader {
	return &Loader{
		namespaceMap: make(map[string]*model.OutputNamespace),
		modelMap:     make(map[string]*model.OutputModel),
		operationMap: make(map[string]*model.OutputOperation),
		dependencies: make(map[string]*Node),
		templateMap:  make(map[string][]model.InputTemplate),
		OutputSpec:   model.NewOutputSpec(),
	}
}

func (t *Loader) FromPath(inputs []string) error {
	if err := util.UnmarshalFromFile(inputs[0], &t.inputSpec); err != nil {
		return err
	}

	t.inputPath = inputs[0]

	if err := t.processInput(); err != nil {
		t.reportErr(err)
	}

	return nil
}
