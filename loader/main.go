package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

type Loader struct {
	inputPath   string
	inputDir    string
	input       model.BiInput
	modelStack  model.StrStack
	reportStack model.StrStack
	typeMapPath string
	typeMap     map[string]TypeInfo
	templateMap map[string][]model.BiInput_Template
	Output      model.BiOutput
}

type TypeInfo struct {
	BaseName   string
	ImportName string
	QName      string
	Namespace  string
	BuiltIn    bool
}

func New() *Loader {
	return &Loader{
		typeMap:     make(map[string]TypeInfo),
		templateMap: make(map[string][]model.BiInput_Template),
		Output: model.BiOutput{
			Models:     make([]model.BiOutput_Model, 0),
			Operations: make([]model.BiOutput_Operation, 0),
			Namespaces: make([]model.BiOutput_Namespace, 0),
		},
	}
}

func (t *Loader) FromPath(input string) error {
	if err := util.UnmarshalFromFile(input, &t.input); err != nil {
		return err
	}

	t.inputPath = input
	t.inputDir = filepath.Dir(t.inputPath)

	if err := t.processInput(); err != nil {
		return t.wrapError(err)
	}

	return nil
}
