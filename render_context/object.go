package render_context

type RenderContextObject struct {
	RenderContextBase
	Properties []RenderContext
}

func (t RenderContextObject) Validation() bool {
	return true
}

var _ RenderContext = &RenderContextObject{}
