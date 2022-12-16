package system

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/template_manager"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type System struct {
	workDir        optioner.Option[string]
	cacheDir       optioner.Option[string]
	projectManager *codegen_project.CodeGenProjectManager
	typeManager    *codegen_type.CodeGenTypeManager
	// cacher         *cacher.Cacher
	generator *generator.Generator
	tm        *template_manager.TemplateManager
	loader    *loader.Loader
}

func New(args ...Arg) (*System, error) {
	t := &System{
		typeManager:    codegen_type.Manager(),
		projectManager: codegen_project.Manager(),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
