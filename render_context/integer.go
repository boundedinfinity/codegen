package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextInteger struct {
	RenderContextBase
	Min        o.Option[int64]
	Max        o.Option[int64]
	MultipleOf o.Option[int64]
}

func (t RenderContextInteger) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.MultipleOf.Defined()
}

var _ RenderContext = &RenderContextInteger{}
