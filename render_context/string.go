package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextString struct {
	RenderContextBase
	Min   o.Option[int]
	Max   o.Option[int]
	Regex o.Option[string]
}

var _ RenderContext = &RenderContextString{}
