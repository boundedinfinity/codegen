package generator

import (
	"boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/render_context"
	"io/fs"
)

type Generator struct {
	genExt          string
	typeManager     *loader_context.CodeGenTypeManager
	projectManager  *loader_context.CodeGenProjectManager
	templateManager *loader_context.CodeGenTemplateManager
	rcs             []render_context.RenderContext
	fileMode        fs.FileMode
	// loader          *loader.Loader
}

func New(args ...Arg) (*Generator, error) {
	g := &Generator{
		rcs: make([]render_context.RenderContext, 0),
	}

	for _, arg := range args {
		arg(g)
	}

	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}
