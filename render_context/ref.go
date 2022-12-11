package render_context

type RenderContextRef struct {
	RenderContextBase
	Ref string
}

func (t RenderContextRef) Validation() bool {
	return false
}

var _ RenderContext = &RenderContextRef{}
