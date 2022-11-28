package generator

import (
	"boundedinfinity/codegen/template_manager"
	"text/template"
)

type Generator struct {
	destDir string
	genExt  string
	tm      *template_manager.TemplateManager
	tmpl    *template.Template
}

func New(args ...Arg) (*Generator, error) {
	g := &Generator{
		tmpl: template.New(""),
	}

	for _, arg := range args {
		arg(g)
	}

	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}
