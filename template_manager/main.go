package template_manager

import (
	"boundedinfinity/codegen/cacher"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type TemplateManager struct {
	pathMap mapper.Mapper[string, []byte]
	cacher  *cacher.Cacher
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap: make(mapper.Mapper[string, []byte]),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
