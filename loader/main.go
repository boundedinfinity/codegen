package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

type Loader struct {
	inputPath      string
	inputDir       string
	inputSpec      model.InputSpec
	namespaceStack model.StrStack
	reportStack    model.StrStack
	dependencies   map[string]*Node
	namespaceMap   map[string]*model.OutputNamespace
	modelMap       map[string]*model.OutputModel
	OutputSpec     *model.OutputSpec
}

func New() *Loader {
	return &Loader{
		namespaceMap: make(map[string]*model.OutputNamespace),
		modelMap:     make(map[string]*model.OutputModel),
		dependencies: make(map[string]*Node),
		OutputSpec:   model.NewOutputSpec(),
	}
}

func (t *Loader) FromPath(inputs []string) error {
	if err := util.UnmarshalFromFile(inputs[0], &t.inputSpec); err != nil {
		return err
	}

	t.inputPath = inputs[0]
	t.inputDir = filepath.Dir(t.inputPath)

	if err := t.processInput(); err != nil {
		t.reportErr(err)
	}

	return nil
}
