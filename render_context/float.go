package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextFloat struct {
	RenderContextBase
	Min        o.Option[float64]
	Max        o.Option[float64]
	MultipleOf o.Option[float64]
}

func (t RenderContextFloat) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

var _ RenderContext = &RenderContextFloat{}
