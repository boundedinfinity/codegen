package render_context

import "boundedinfinity/codegen/canonical/canonical_type"

type RenderContext interface {
	Base() *RenderContextBase
}

type RenderContextBase struct {
	OutputPath    string
	SourceUri     string
	Id            string
	RootNs        string
	CurrNs        string
	SchemaNs      string
	RelNs         string
	SchemaType    canonical_type.CanonicalType
	Name          string
	Description   string
	IsPublic      bool
	IsRequired    bool
	IsInterface   bool
	Header        string
	HasValidation bool
}

func (t *RenderContextBase) Base() *RenderContextBase {
	return t
}
