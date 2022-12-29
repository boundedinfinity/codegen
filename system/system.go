package system

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/manager"
	"boundedinfinity/codegen/renderer"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type System struct {
	workDir         optioner.Option[string]
	cacheDir        optioner.Option[string]
	projectManager  *manager.CodeGenProjectManager
	typeManager     *manager.CodeGenTypeManager
	templateManager *manager.CodeGenTemplateManager
	generator       *generator.Generator
	loader          *loader.Loader
	renderer        *renderer.Renderer
	// cacher         *cacher.Cacher
}

func New(args ...Arg) (*System, error) {
	t := &System{
		typeManager:     manager.TypeManager(),
		projectManager:  manager.ProjectManager(),
		templateManager: manager.TemplateManager(),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
