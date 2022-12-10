package render_context

type RenderContextArray struct {
	RenderContextBase
	Items RenderContext
}

func (t RenderContextArray) Validation() bool {
	return true
}

var _ RenderContext = &RenderContextArray{}
