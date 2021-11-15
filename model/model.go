package model

import (
	"boundedinfinity/codegen/uritype"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/jsonschema/mimetype"
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
	Name        string                `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Input       jsonschema.JsonSchmea `json:"input,omitempty" yaml:"input,omitempty"`
	Output      jsonschema.JsonSchmea `json:"output,omitempty" yaml:"output,omitempty"`
}

type Template struct {
	Name        string                `json:"name,omitempty" yaml:"name,omitempty"`
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Input       jsonschema.JsonSchmea `json:"input,omitempty" yaml:"input,omitempty"`
	Output      jsonschema.JsonSchmea `json:"output,omitempty" yaml:"output,omitempty"`
}

type SourceInfo struct {
	SourceUri string
	UriType   uritype.UriType
	LocalPath string
	MimeType  mimetype.MimeType
}

// type InputInfo struct {
// 	InputDir                  string            `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
// 	OutputDir                 string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
// 	DumpContext               bool              `json:"dumpContext" yaml:"dumpContext"`
// 	FilenameMarker            string            `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
// 	TemplateHeader            string            `json:"templateHeader,omitempty" yaml:"templateHeader,omitempty"`
// 	Namespace                 string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
// 	DescriptionSplitCharacter string            `json:"descriptionSplitCharacter,omitempty" yaml:"descriptionSplitCharacter,omitempty"`
// 	Primitives                map[string]string `json:"primitives,omitempty" yaml:"primitives,omitempty"`
// }

// type InputTemplate struct {
// 	Header    string           `json:"header,omitempty" yaml:"header,omitempty"`
// 	Path      string           `json:"path,omitempty" yaml:"path,omitempty"`
// 	Type      TemplateTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
// 	Namespace string           `json:"namespace,omitempty" yaml:"namespace,omitempty"`
// 	Recurse   bool             `json:"recurse,omitempty" yaml:"recurse,omitempty"`
// }
