package render_context

type RenderContextObject struct {
	RenderContextBase
	Properties []RenderContext
}

func (t RenderContextObject) HasValidation() bool {
	for _, property := range t.Properties {
		if property.HasValidation() {
			return true
		}
	}

	return false
}

var _ RenderContext = &RenderContextObject{}
