package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextString struct {
	RenderContextBase
	Min   o.Option[int]
	Max   o.Option[int]
	Regex o.Option[string]
}

func (t RenderContextString) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Regex.Defined()
}

var _ RenderContext = &RenderContextString{}
