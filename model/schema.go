package model

import (
	"boundedinfinity/codegen/header_strategy"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema/model"
)

type CodeGenSchema struct {
	Info       o.Option[CodeGenSchemaInfo]                    `json:"info,omitempty" yaml:"info,omitempty"`
	Mappings   mapper.Mapper[string, *CodeGenSchemaMappings]  `json:"mappings,omitempty" yaml:"mappings,omitempty"`
	Operations mapper.Mapper[string, *CodeGenSchemaOperation] `json:"operations,omitempty" yaml:"operations,omitempty"`
	Templates  CodeGenSchemaTemplates                         `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewSchema() *CodeGenSchema {
	return &CodeGenSchema{
		Info:       o.Some(NewInfo()),
		Mappings:   mapper.Mapper[string, *CodeGenSchemaMappings]{},
		Operations: mapper.Mapper[string, *CodeGenSchemaOperation]{},
	}
}

type CodeGenSchemaInfo struct {
	Id          o.Option[string] `json:"$id" yaml:"$id"`
	Name        o.Option[string] `json:"name,omitempty" yaml:"name,omitempty"`
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Version     o.Option[string] `json:"version,omitempty" yaml:"version,omitempty"`
	Header      o.Option[string] `json:"header,omitempty" yaml:"header,omitempty"`
}

func NewInfo() CodeGenSchemaInfo {
	return CodeGenSchemaInfo{}
}

type CodeGenSchemaMappings struct {
	Package o.Option[string]              `json:"package,omitempty" yaml:"package,omitempty"`
	RootDir o.Option[string]              `json:"rootDir,omitempty" yaml:"rootDir,omitempty"`
	Replace mapper.Mapper[string, string] `json:"replace,omitempty" yaml:"replace,omitempty"`
}

func NewMappings() *CodeGenSchemaMappings {
	return &CodeGenSchemaMappings{
		Replace: mapper.Mapper[string, string]{},
	}
}

type CodeGenSchemaOperation struct {
	Description o.Option[string]    `json:"description,omitempty" yaml:"description,omitempty"`
	Input       o.Option[model.IdT] `json:"input,omitempty" yaml:"input,omitempty"`
	Output      o.Option[model.IdT] `json:"output,omitempty" yaml:"output,omitempty"`
}

type CodeGenSchemaTemplates struct {
	Header CodeGenSchemaHeader         `json:"header,omitempty" yaml:"header,omitempty"`
	Files  []CodeGenSchemaTemplateFile `json:"files,omitempty" yaml:"files,omitempty"`
}

type CodeGenSchemaHeader struct {
	Content  o.Option[string]                         `json:"content,omitempty" yaml:"content,omitempty"`
	Path     o.Option[string]                         `json:"path,omitempty" yaml:"path,omitempty"`
	Strategy o.Option[header_strategy.HeaderStrategy] `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}

type CodeGenSchemaTemplateFile struct {
	Path   o.Option[string]              `json:"path,omitempty" yaml:"path,omitempty"`
	Header o.Option[CodeGenSchemaHeader] `json:"header,omitempty" yaml:"header,omitempty"`
}
