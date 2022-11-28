package template_manager

import "github.com/boundedinfinity/go-commoner/optioner/mapper"

type TemplateManager struct {
	pathMap mapper.Mapper[string, string]
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap: make(mapper.Mapper[string, string]),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
