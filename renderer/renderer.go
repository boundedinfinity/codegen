package renderer

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/render_context"
	"text/template"
)

type TemplateOutput struct {
	codegen_type.TemplateContext
	Output []byte
}

type ModelOutput struct {
	OutputPath string
	TemplateOutput
	Schema render_context.RenderContext
}

type Renderer struct {
	projectManager  *manager.CodeGenProjectManager
	typeManager     *manager.CodeGenTypeManager
	templateManager *manager.CodeGenTemplateManager
	funcs           template.FuncMap
	verbose         bool
}

func New(args ...Arg) (*Renderer, error) {
	t := &Renderer{
		funcs: make(template.FuncMap),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
