package render_context

type RenderContextUrl struct {
	RenderContextBase
}

func (t RenderContextUrl) HasValidation() bool {
	return true
}

var _ RenderContext = &RenderContextUrl{}
