package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_manager"
	"io/fs"
)

type Generator struct {
	destDir       string
	genExt        string
	tm            *template_manager.TemplateManager
	canonicals    *canonical.CanonicalCombined
	codeGenSchema *model.CodeGenSchema
	rcs           []render_context.RenderContext
	fileMode      fs.FileMode
	loader        *loader.Loader
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
