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
	Output      model.BiOutput
}

type TypeInfo struct {
	InNamespace  string
	OutNamespace string
	Namespace    string
}

func New() *Loader {
	return &Loader{
		typeMap: make(map[string]TypeInfo),
		Output: model.BiOutput{
			Models: model.BiGenModel{
				Namespaces: make([]model.BiOutput_Model_Namespace, 0),
			},
			Operations: model.BiGenOperation{
				Namespaces: make([]model.BiOutput_Model_Namespace, 0),
			},
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
