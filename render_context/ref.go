package render_context

type RenderContextRef struct {
	RenderContextBase
	Ref RenderContext
}

var _ RenderContext = &RenderContextRef{}
