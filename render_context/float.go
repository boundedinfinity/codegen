package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextFloat struct {
	RenderContextBase
	Min        o.Option[float64]
	Max        o.Option[float64]
	MultipleOf o.Option[float64]
}

var _ RenderContext = &RenderContextFloat{}
