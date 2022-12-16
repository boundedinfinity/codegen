package codegen_project

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenProjectTemplates struct {
	Header o.Option[CodeGenProjectHeader] `json:"header,omitempty" yaml:"header,omitempty"`
	Root   o.Option[string]               `json:"root,omitempty" yaml:"root,omitempty"`
	Files  []CodeGenProjectTemplateFile   `json:"files,omitempty" yaml:"files,omitempty"`
}

type CodeGenProjectTemplateFile struct {
	Header  o.Option[CodeGenProjectHeader] `json:"header,omitempty" yaml:"header,omitempty"`
	Path    o.Option[string]               `json:"path,omitempty" yaml:"path,omitempty"`
	Content o.Option[string]               `json:"content,omitempty" yaml:"content,omitempty"`
}
