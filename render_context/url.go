package render_context

type RenderContextUrl struct {
	RenderContextBase
}

func (t RenderContextUrl) Validation() bool {
	return true
}

var _ RenderContext = &RenderContextUrl{}
