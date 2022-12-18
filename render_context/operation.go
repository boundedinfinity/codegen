package render_context

type RenderContextOperation struct {
	RenderContextBase
	Input  string `json:"input,omitempty" yaml:"input,omitempty"`
	Output string `json:"output,omitempty" yaml:"output,omitempty"`
}

var _ RenderContext = &RenderContextOperation{}
