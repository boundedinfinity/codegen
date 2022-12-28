package render_context

type RenderContextRef struct {
	RenderContextBase
	Ref RenderContext
}

func (t RenderContextRef) HasValidation() bool {
	return false
}

var _ RenderContext = &RenderContextRef{}
