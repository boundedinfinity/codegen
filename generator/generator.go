package generator

import (
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/renderer"
	"io/fs"
)

type Generator struct {
	genExt          string
	typeManager     *manager.CodeGenTypeManager
	projectManager  *manager.CodeGenProjectManager
	templateManager *manager.CodeGenTemplateManager
	renderer        *renderer.Renderer
	fileMode        fs.FileMode
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
