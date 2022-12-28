package render_context

type RenderContextOperation struct {
	RenderContextBase
	Input  string `json:"input,omitempty" yaml:"input,omitempty"`
	Output string `json:"output,omitempty" yaml:"output,omitempty"`
}

func (t RenderContextOperation) HasValidation() bool {
	return false
}

var _ RenderContext = &RenderContextOperation{}
