package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/template_type"
	"text/template"

	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-jsonschema"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type TemplateContext struct {
	TemplateMimeType mime_type.MimeType
	TemplateType     template_type.TemplateType
	Template         *template.Template
	OutputMimeType   mime_type.MimeType
	Path             string
}

type TemplateOutput struct {
	MimeType     mime_type.MimeType
	TemplateType template_type.TemplateType
	Output       []byte
	Path         string
}

type TemplateManager struct {
	pathMap      mapper.Mapper[string, TemplateContext]
	cacher       *cacher.Cacher
	funcs        template.FuncMap
	formatSource bool
	jsonSchemas  *jsonschema.System
	verbose      bool
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap: make(mapper.Mapper[string, TemplateContext]),
		funcs:   make(template.FuncMap),
	}

	args = append(args,
		TemplateFunc("dumpJson", dumpJson),
	)

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
