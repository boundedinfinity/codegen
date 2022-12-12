package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextInteger struct {
	RenderContextBase
	Min        o.Option[int64]
	Max        o.Option[int64]
	MultipleOf o.Option[int64]
}

var _ RenderContext = &RenderContextInteger{}
