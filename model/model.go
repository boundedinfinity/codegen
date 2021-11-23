package model

import (
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/uritype"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/mimetyper/mime_type"
	"github.com/boundedinfinity/optional"
)

// Reference
// https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0

type Schema struct {
	Id         string                   `json:"$id" yaml:"$id"`
	Info       Info                     `json:"info,omitempty" yaml:"info,omitempty"`
	Mappings   Mappings                 `json:"mappings,omitempty" yaml:"mappings,omitempty"`
	Operations []Operation              `json:"operations,omitempty" yaml:"operations,omitempty"`
	Models     []*jsonschema.JsonSchmea `json:"models,omitempty" yaml:"models,omitempty"`
	Defs       []*jsonschema.JsonSchmea `json:"$defs,omitempty" yaml:"$defs,omitempty"`
	Templates  Templates                `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type Info struct {
	Name        optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty"`
	Description optional.StringOptional `json:"description,omitempty" yaml:"description,omitempty"`
	Version     optional.StringOptional `json:"version,omitempty" yaml:"version,omitempty"`
}

type Mappings struct {
	Language map[string]string `json:"language,omitempty" yaml:"language,omitempty"`
	Package  map[string]string `json:"package,omitempty" yaml:"package,omitempty"`
}

type Operation struct {
	Name        string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                 `json:"description,omitempty" yaml:"description,omitempty"`
	Input       *jsonschema.JsonSchmea `json:"input,omitempty" yaml:"input,omitempty"`
	Output      *jsonschema.JsonSchmea `json:"output,omitempty" yaml:"output,omitempty"`
}

type Templates struct {
	Name   string         `json:"name,omitempty" yaml:"name,omitempty"`
	Header string         `json:"header,omitempty" yaml:"header,omitempty"`
	Files  []TemplateFile `json:"files,omitempty" yaml:"files,omitempty"`
}

type TemplateFile struct {
	Name   string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Header optional.StringOptional    `json:"header,omitempty" yaml:"header,omitempty"`
	Type   template_type.TemplateType `json:"type,omitempty" yaml:"type,omitempty"`
}

type SourceInfo struct {
	SourceUri string
	UriType   uritype.UriType
	LocalPath string
	MimeType  mime_type.MimeType
}
