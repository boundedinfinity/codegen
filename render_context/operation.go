package render_context

type RenderContextOperation struct {
	RenderContextBase
	Input  RenderContext
	Output RenderContext
}

func (t RenderContextOperation) HasValidation() bool {
	return false
}

var _ RenderContext = &RenderContextOperation{}
