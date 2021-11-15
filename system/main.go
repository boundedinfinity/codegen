package system

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/jsonschema"
)

type System struct {
	sourceInfo map[string]*model.SourceInfo
	codeGen    map[string]*model.Schema
	jsonSchema *jsonschema.System
}

func New() *System {
	return &System{
		sourceInfo: make(map[string]*model.SourceInfo),
		codeGen:    make(map[string]*model.Schema),
		jsonSchema: jsonschema.New(),
	}
}
