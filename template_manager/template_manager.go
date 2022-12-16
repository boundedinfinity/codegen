package template_manager

import (
	"boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/template_type"
	"text/template"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type TemplateContext struct {
	TemplateMimeType mime_type.MimeType
	OutputMimeType   mime_type.MimeType
	TemplateType     template_type.TemplateType
	ModelType        optioner.Option[codegen_type_id.CodgenTypeId]
	Template         *template.Template
	Source           string
	Root             string
}

type TemplateOutput struct {
	TemplateContext
	Output []byte
}

type ModelOutput struct {
	TemplateOutput
	Schema render_context.RenderContext
}

type TemplateManager struct {
	projectManager *codegen_project.CodeGenProjectManager
	typeManager    *codegen_type.CodeGenTypeManager
	pathMap        mapper.Mapper[string, TemplateContext]
	modelMap       mapper.Mapper[template_type.TemplateType, []TemplateContext]
	schemaMap      mapper.Mapper[codegen_type_id.CodgenTypeId, []TemplateContext]
	// cacher         *cacher.Cacher
	funcs   template.FuncMap
	verbose bool
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap:   make(mapper.Mapper[string, TemplateContext]),
		modelMap:  make(mapper.Mapper[template_type.TemplateType, []TemplateContext]),
		schemaMap: make(mapper.Mapper[codegen_type_id.CodgenTypeId, []TemplateContext]),
		funcs:     make(template.FuncMap),
	}

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
