package system

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/loader_context"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type System struct {
	workDir         optioner.Option[string]
	cacheDir        optioner.Option[string]
	projectManager  *loader_context.CodeGenProjectManager
	typeManager     *loader_context.CodeGenTypeManager
	templateManager *loader_context.CodeGenTemplateManager
	generator       *generator.Generator
	loader          *loader.Loader
	// tm              *template_manager.TemplateManager
	// cacher         *cacher.Cacher
}

func New(args ...Arg) (*System, error) {
	t := &System{
		typeManager:     loader_context.TypeManager(),
		projectManager:  loader_context.ProjectManager(),
		templateManager: loader_context.TemplateManager(),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
