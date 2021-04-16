package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

type Loader struct {
	specPath    string
	specDir     string
	typeMapPath string
	stack       model.StrStack
	spec        model.BiSpec
	Gen         model.BiGen
	typeMap     map[string]string
	rtypeMap    map[string]string
}

func New() *Loader {
	return &Loader{
		typeMap:  make(map[string]string),
		rtypeMap: make(map[string]string),
		Gen: model.BiGen{
			Models: model.BiGenModel{
				Namespaces: make([]model.BiGenNamespace, 0),
			},
			Operations: model.BiGenOperation{
				Namespaces: make([]model.BiGenNamespace, 0),
			},
		},
	}
}

func (t *Loader) FromPath(input string) error {
	if err := util.UnmarshalFromFile(input, &t.spec); err != nil {
		return err
	}

	t.specPath = input
	t.specDir = filepath.Dir(t.specPath)

	if err := t.processSpec(); err != nil {
		return err
	}

	return nil
}
