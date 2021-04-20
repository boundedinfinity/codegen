package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

type Loader struct {
	inputPath      string
	inputDir       string
	input          model.BiInput
	modelStack     model.StrStack
	reportStack    model.StrStack
	builtInTypeMap map[string]string
	customTypeMap  map[string]string
	templateMap    map[string][]model.BiInput_Template
	modelMap       map[string]*model.BiOutput_Model
	namespaceMap   map[string]*model.BiOutput_Namespace
	operationMap   map[string]*model.BiOutput_Operation
	Output         model.BiOutput
}

func New() *Loader {
	return &Loader{
		customTypeMap:  make(map[string]string),
		builtInTypeMap: make(map[string]string),
		templateMap:    make(map[string][]model.BiInput_Template),
		modelMap:       make(map[string]*model.BiOutput_Model),
		namespaceMap:   make(map[string]*model.BiOutput_Namespace),
		operationMap:   make(map[string]*model.BiOutput_Operation),
		Output: model.BiOutput{
			Models:     make([]model.BiOutput_Model, 0),
			Operations: make([]model.BiOutput_Operation, 0),
			Namespaces: make([]model.BiOutput_Namespace, 0),
		},
	}
}

func (t *Loader) FromPath(inputs []string) error {
	if err := util.UnmarshalFromFile(inputs[0], &t.input); err != nil {
		return err
	}

	t.inputPath = inputs[0]
	t.inputDir = filepath.Dir(t.inputPath)

	if err := t.processInput(); err != nil {
		return t.wrapError(err)
	}

	return nil
}
