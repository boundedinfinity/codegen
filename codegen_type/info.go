package codegen_type

import (
	"boundedinfinity/codegen/template_delimiter"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectInfo struct {
	Id           o.Option[string]                               `json:"$id" yaml:"$id"`
	Name         o.Option[string]                               `json:"name,omitempty"`
	Description  o.Option[string]                               `json:"description,omitempty"`
	Version      o.Option[string]                               `json:"version,omitempty"`
	Header       o.Option[string]                               `json:"header,omitempty"`
	Namespace    o.Option[string]                               `json:"namespace,omitempty"`
	DestDir      o.Option[string]                               `json:"destDir,omitempty"`
	FormatSource o.Option[bool]                                 `json:"formatSource,omitempty"`
	TemplateDump o.Option[bool]                                 `json:"templateDump,omitempty"`
	Delimiter    o.Option[template_delimiter.TemplateDelimiter] `json:"delimiter,omitempty"`
}

func NewInfo() CodeGenProjectInfo {
	return CodeGenProjectInfo{}
}
