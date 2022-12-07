package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"text/template"

	"github.com/boundedinfinity/go-commoner/caser"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type RenderContext struct {
	Info   model.CodeGenSchemaInfo
	Schema canonical.Canonical
}

type TemplateContext struct {
	TemplateMimeType mime_type.MimeType
	OutputMimeType   mime_type.MimeType
	TemplateType     template_type.TemplateType
	Path             string
}

type TemplateOutput struct {
	MimeType     mime_type.MimeType
	TemplateType template_type.TemplateType
	Output       []byte
	Path         string
}

type TemplateManager struct {
	codeGenSchema     *model.CodeGenSchema
	pathMap           mapper.Mapper[string, TemplateContext]
	modelMap          mapper.Mapper[string, []TemplateContext]
	cacher            *cacher.Cacher
	funcs             template.FuncMap
	formatSource      bool
	verbose           bool
	combinedTemplates *template.Template
	canonicals        *canonical.CanonicalCombined
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap:           make(mapper.Mapper[string, TemplateContext]),
		modelMap:          make(mapper.Mapper[string, []TemplateContext]),
		funcs:             make(template.FuncMap),
		combinedTemplates: template.New(""),
	}

	args = append(args,
		TemplateFunc("DUMP", dumpJson),
		TemplateFunc("PASCAL", caser.KebabToPascal[string]),
		TemplateFunc("CAMEL", caser.KebabToCamel[string]),
		TemplateFunc("IMPORT_PATH", t.importPath),
		TemplateFunc("IMPORT_TYPE", t.importType),
		TemplateFunc("OBJ_NAME", t.objName),
		TemplateFunc("OBJ_PKG", t.objPackage),
		TemplateFunc("OBJ_PKG_BASE", t.objPackageBase),
		TemplateFunc("PACKAGE_DIR", t.objPackage),
	)

	for _, arg := range args {
		arg(t)
	}

	if err := t.init(); err != nil {
		return nil, err
	}

	return t, nil
}
