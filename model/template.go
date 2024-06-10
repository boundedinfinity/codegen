package model

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type TemplateMeta struct {
	// Type             o.Option[codegen_type_id.CodgenTypeId]
	// OutputMimeType   mime_type.MimeType
	// OutputExt        string
	// TemplateType     template_type.TemplateType
	// TemplateMimeTime mime_type.MimeType
	// TemplateExt      string
	// Template         *template.Template
}

// type CodeGenProjectTemplateFile struct {
// 	TemplateMeta
// 	Header  o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
// 	Content o.Option[string]                `json:"content,omitempty"`
// }

type CodeGenProjectTemplates struct {
	// Header     o.Option[CodeGenTemplateHeader] `json:"header,omitempty"`
	// Types      []*CodeGenProjectTemplateFile   `json:"types,omitempty"`
	// TypeList   []*CodeGenProjectTemplateFile   `json:"type-list,omitempty"`
	// Operations []*CodeGenProjectTemplateFile   `json:"operations,omitempty"`
}

// var _ LoaderContext = &CodeGenProjectTemplateFile{}
// var _ LoaderContext = &TemplateMeta{}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

func (t CodeGenProjectTemplates) Validate() error {

	return nil
}
