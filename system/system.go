package system

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_manager"

	"github.com/boundedinfinity/go-commoner/optioner"
)

type System struct {
	workDir       optioner.Option[string]
	outputDir     optioner.Option[string]
	cacheDir      optioner.Option[string]
	mergedCodeGen *model.CodeGenSchema
	canonicals    *canonical.CanonicalCombined
	cacher        *cacher.Cacher
	generator     *generator.Generator
	tm            *template_manager.TemplateManager
	loader        *loader.Loader
}

func New(args ...Arg) (*System, error) {
	t := &System{
		canonicals:    canonical.NewCombinded(),
		mergedCodeGen: model.NewSchema(),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
