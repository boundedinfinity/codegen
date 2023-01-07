package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/renderer"
	"io/fs"
)

type Generator struct {
	genExt          string
	typeManager     *manager.CodeGenTypeManager
	projectManager  *manager.CodeGenProjectManager
	templateManager *manager.CodeGenTemplateManager
	types           []ct.CodeGenType
	operations      []ct.CodeGenProjectOperationTemplateFile
	renderer        *renderer.Renderer
	fileMode        fs.FileMode
}

func New(args ...Arg) (*Generator, error) {
	g := &Generator{
		types: make([]ct.CodeGenType, 0),
	}

	for _, arg := range args {
		arg(g)
	}

	if err := g.init(); err != nil {
		return nil, err
	}

	return g, nil
}
