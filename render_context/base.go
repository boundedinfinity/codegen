package render_context

type RenderContext interface {
	Validation() bool
}

type RenderContextBase struct {
	SourceUri   string
	RootNs      string
	CurrNs      string
	SchemaNs    string
	SchemaType  string
	Name        string
	Description string
	IsPublic    bool
	IsRequired  bool
	IsInterface bool
}
