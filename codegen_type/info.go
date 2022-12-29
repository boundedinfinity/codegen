package codegen_type

import (
	"boundedinfinity/codegen/template_delimiter"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectInfo struct {
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

func NewInfo() CodeGenProjectInfo {
	return CodeGenProjectInfo{}
}
