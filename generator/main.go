package generator

import (
	"boundedinfinity/codegen/template_manager"

	"github.com/boundedinfinity/go-jsonschema"
)

type Generator struct {
	destDir     string
	genExt      string
	tm          *template_manager.TemplateManager
	jsonSchemas *jsonschema.System
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
