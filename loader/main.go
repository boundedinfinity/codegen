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
	Mapper      model.BiGenLangTypeMapper
}

func New() *Loader {
	return &Loader{
		typeMapPath: "data/type-maps.yaml",
		Gen: model.BiGen{
			Lookup: make(map[string]model.BiGenType),
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

	if err := t.processMapper(); err != nil {
		return err
	}

	return nil
}
