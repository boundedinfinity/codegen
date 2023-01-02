package codegen_type

import (
	"boundedinfinity/codegen/codegen_type/template_delimiter"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenInfo struct {
	Id           o.Option[string]                               `json:"$id" yaml:"$id"`
	Name         o.Option[string]                               `json:"name,omitempty"`
	Description  o.Option[string]                               `json:"description,omitempty"`
	Version      o.Option[string]                               `json:"version,omitempty"`
	Header       o.Option[string]                               `json:"header,omitempty"`
	Namespace    o.Option[string]                               `json:"namespace,omitempty"`
	DestDir      o.Option[string]                               `json:"dest-dir,omitempty"`
	FormatSource o.Option[bool]                                 `json:"format-source,omitempty"`
	TemplateDump o.Option[bool]                                 `json:"template-dump,omitempty"`
	Delimiter    o.Option[template_delimiter.TemplateDelimiter] `json:"delimiter,omitempty"`
}

func NewInfo() CodeGenInfo {
	return CodeGenInfo{}
}
