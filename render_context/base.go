package render_context

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
	SchemaType    string
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
