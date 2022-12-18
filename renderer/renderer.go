package renderer

import (
	"boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/render_context"
	"text/template"
)

type TemplateOutput struct {
	loader_context.TemplateLoaderContext
	Output []byte
}

type ModelOutput struct {
	TemplateOutput
	Schema render_context.RenderContext
}

type Renderer struct {
	projectManager  *loader_context.CodeGenProjectManager
	typeManager     *loader_context.CodeGenTypeManager
	templateManager *loader_context.CodeGenTemplateManager
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
