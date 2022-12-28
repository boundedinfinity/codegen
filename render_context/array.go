package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextArray struct {
	RenderContextBase
	Items RenderContext
	Min   o.Option[int] `json:"min,omitempty" yaml:"min,omitempty"`
	Max   o.Option[int] `json:"max,omitempty" yaml:"max,omitempty"`
}

func (t RenderContextArray) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Items.HasValidation()
}

var _ RenderContext = &RenderContextArray{}
