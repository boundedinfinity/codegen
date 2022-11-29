package system

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_manager"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema"
)

type System struct {
	workDir     optioner.Option[string]
	outputDir   optioner.Option[string]
	cacheDir    optioner.Option[string]
	jsonSchemas *jsonschema.System
	pathMap     mapper.Mapper[string, model.CodeGenSchema]
	combined    *model.CodeGenSchema
	cacher      *cacher.Cacher
	generator   *generator.Generator
	tm          *template_manager.TemplateManager
}

func New(args ...Arg) (*System, error) {
	t := &System{
		jsonSchemas: jsonschema.New(),
		pathMap:     mapper.Mapper[string, model.CodeGenSchema]{},
		combined:    model.NewSchema(),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
