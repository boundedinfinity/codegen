package model

import (
	"boundedinfinity/codegen/header_strategy"
	"boundedinfinity/codegen/template_delimiter"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenSchema struct {
	Info       CodeGenSchemaInfo                              `json:"info,omitempty" yaml:"info,omitempty"`
	Mappings   mapper.Mapper[string, string]                  `json:"mappings,omitempty" yaml:"mappings,omitempty"`
	Operations mapper.Mapper[string, *CodeGenSchemaOperation] `json:"operations,omitempty" yaml:"operations,omitempty"`
	Templates  CodeGenSchemaTemplates                         `json:"templates,omitempty" yaml:"templates,omitempty"`
	Schemas    []CodeGenSchemaFile                            `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

func NewSchema() *CodeGenSchema {
	return &CodeGenSchema{
		Info:       NewInfo(),
		Mappings:   mapper.Mapper[string, string]{},
		Operations: mapper.Mapper[string, *CodeGenSchemaOperation]{},
	}
}

type CodeGenSchemaInfo struct {
	Id           o.Option[string]                               `json:"$id" yaml:"$id"`
	Name         o.Option[string]                               `json:"name,omitempty" yaml:"name,omitempty"`
	Description  o.Option[string]                               `json:"description,omitempty" yaml:"description,omitempty"`
	Version      o.Option[string]                               `json:"version,omitempty" yaml:"version,omitempty"`
	Header       o.Option[string]                               `json:"header,omitempty" yaml:"header,omitempty"`
	Namespace    o.Option[string]                               `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	DestDir      o.Option[string]                               `json:"destDir,omitempty" yaml:"destDir,omitempty"`
	FormatSource o.Option[bool]                                 `json:"formatSource,omitempty" yaml:"formatSource,omitempty"`
	TemplateDump o.Option[bool]                                 `json:"templateDump,omitempty" yaml:"templateDump,omitempty"`
	Delimiter    o.Option[template_delimiter.TemplateDelimiter] `json:"delimiter,omitempty" yaml:"delimiter,omitempty"`
}

func NewInfo() CodeGenSchemaInfo {
	return CodeGenSchemaInfo{}
}

type CodeGenSchemaOperation struct {
	Description o.Option[string] `json:"description,omitempty" yaml:"description,omitempty"`
	Input       o.Option[string] `json:"input,omitempty" yaml:"input,omitempty"`
	Output      o.Option[string] `json:"output,omitempty" yaml:"output,omitempty"`
}

type CodeGenSchemaFile struct {
	Path o.Option[string] `json:"path,omitempty" yaml:"path,omitempty"`
}

type CodeGenSchemaTemplates struct {
	Header o.Option[CodeGenSchemaHeader] `json:"header,omitempty" yaml:"header,omitempty"`
	Files  []CodeGenSchemaTemplateFile   `json:"files,omitempty" yaml:"files,omitempty"`
}

type CodeGenSchemaHeader struct {
	Content  o.Option[string]                         `json:"content,omitempty" yaml:"content,omitempty"`
	Path     o.Option[string]                         `json:"path,omitempty" yaml:"path,omitempty"`
	Strategy o.Option[header_strategy.HeaderStrategy] `json:"strategy,omitempty" yaml:"strategy,omitempty"`
}

type CodeGenSchemaTemplateFile struct {
	Header  o.Option[CodeGenSchemaHeader] `json:"header,omitempty" yaml:"header,omitempty"`
	Path    o.Option[string]              `json:"path,omitempty" yaml:"path,omitempty"`
	Content o.Option[string]              `json:"content,omitempty" yaml:"content,omitempty"`
}
