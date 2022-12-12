package render_context

type RenderContextArray struct {
	RenderContextBase
	Items RenderContext
}

var _ RenderContext = &RenderContextArray{}
