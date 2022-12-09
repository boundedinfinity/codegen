package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_manager"
	"io/fs"
)

type Generator struct {
	destDir       string
	genExt        string
	tm            *template_manager.TemplateManager
	canonicals    *canonical.CanonicalCombined
	codeGenSchema *model.CodeGenSchema
	fileMode      fs.FileMode
}

func New(args ...Arg) (*Generator, error) {
	g := &Generator{}

	for _, arg := range args {
		arg(g)
	}

	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}
