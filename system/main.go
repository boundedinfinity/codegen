package system

import "boundedinfinity/codegen/model"

type System struct {
	sourceInfo map[string]*model.SourceInfo
	// codeGen    map[string]*model.Schema
	// template   map[string]*model.TemplateFile
	// jsonSchema *jsonschema.System
}

func New() *System {
	return &System{
		sourceInfo: make(map[string]*model.SourceInfo),
		// codeGen:    make(map[string]*model.Schema),
		// template:   make(map[string]*model.TemplateFile),
		// jsonSchema: jsonschema.New(),
	}
}
