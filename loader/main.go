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
	namespaceStack model.StrStack
	reportStack    model.StrStack
	depNodes       map[string]*Node
	typeMap        map[string]*model.TypeInfo
	templateMap    map[string][]model.BiInput_Template
	namespaceMap   map[string]*model.BiOutput_Namespace
	modelMap       map[string]*model.BiOutput_Model
	propertyMap    map[string]*model.BiOutput_Property
	operationMap   map[string]*model.BiOutput_Operation
	Output         model.BiOutput
}

func New() *Loader {
	return &Loader{
		typeMap:      make(map[string]*model.TypeInfo),
		namespaceMap: make(map[string]*model.BiOutput_Namespace),
		modelMap:     make(map[string]*model.BiOutput_Model),
		propertyMap:  make(map[string]*model.BiOutput_Property),
		operationMap: make(map[string]*model.BiOutput_Operation),
		templateMap:  make(map[string][]model.BiInput_Template),
		depNodes:     make(map[string]*Node),
		Output:       model.New_BiOutput(),
	}
}

func (t *Loader) FromPath(inputs []string) error {
	if err := util.UnmarshalFromFile(inputs[0], &t.input); err != nil {
		return err
	}

	t.inputPath = inputs[0]
	t.inputDir = filepath.Dir(t.inputPath)

	if err := t.processInput(); err != nil {
		t.report("%v", err.Error())
	}

	return nil
}
