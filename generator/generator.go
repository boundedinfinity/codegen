package generator

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_manager"
	"io/fs"
)

type Generator struct {
	genExt         string
	tm             *template_manager.TemplateManager
	typeManager    *codegen_type.CodeGenTypeManager
	projectManager *codegen_project.CodeGenProjectManager
	rcs            []render_context.RenderContext
	fileMode       fs.FileMode
	loader         *loader.Loader
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
