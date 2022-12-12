package render_context

type RenderContextObject struct {
	RenderContextBase
	Properties []RenderContext
}

var _ RenderContext = &RenderContextObject{}
