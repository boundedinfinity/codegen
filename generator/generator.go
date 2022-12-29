package generator

import (
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/renderer"
	"io/fs"
)

type Generator struct {
	genExt          string
	typeManager     *manager.CodeGenTypeManager
	projectManager  *manager.CodeGenProjectManager
	templateManager *manager.CodeGenTemplateManager
	types           []render_context.RenderContext
	operations      []render_context.RenderContextOperation
	renderer        *renderer.Renderer
	fileMode        fs.FileMode
}

func New(args ...Arg) (*Generator, error) {
	g := &Generator{
		types: make([]render_context.RenderContext, 0),
	}

	for _, arg := range args {
		arg(g)
	}

	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}
