package model

import "github.com/boundedinfinity/jsonschema"

// Reference
// https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0

type File struct {
	Name       string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Version    string                  `json:"version,omitempty" yaml:"version,omitempty"`
	TypeMaps   string                  `json:"typeMaps,omitempty" yaml:"typeMaps,omitempty"`
	Operations []Operation             `json:"operations,omitempty" yaml:"operations,omitempty"`
	Models     []jsonschema.JsonSchmea `json:"models,omitempty" yaml:"models,omitempty"`
	Defs       []jsonschema.JsonSchmea `json:"$defs,omitempty" yaml:"$defs,omitempty"`
}

type TypeMaps struct {
	LangType map[string]string `json:"lang-type,omitempty" yaml:"lang-type,omitempty"`
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
