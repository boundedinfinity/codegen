package template_manager

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/canonical/canonical_type"
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_type"
	"fmt"
	"text/template"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
)

type TemplateContext struct {
	TemplateMimeType mime_type.MimeType
	OutputMimeType   mime_type.MimeType
	TemplateType     template_type.TemplateType
	ModelType        optioner.Option[canonical_type.CanonicalType]
	Path             string
}

type TemplateOutput struct {
	TemplateContext
	Output []byte
}

type ModelOutput struct {
	TemplateOutput
	Schema canonical.Canonical
}

type TemplateManager struct {
	codeGenSchema     *model.CodeGenSchema
	pathMap           mapper.Mapper[string, TemplateContext]
	modelMap          mapper.Mapper[template_type.TemplateType, []TemplateContext]
	schemaMap         mapper.Mapper[canonical_type.CanonicalType, []TemplateContext]
	cacher            *cacher.Cacher
	funcs             template.FuncMap
	formatSource      bool
	verbose           bool
	combinedTemplates *template.Template
	canonicals        *canonical.CanonicalCombined
}

func New(args ...Arg) (*TemplateManager, error) {
	t := &TemplateManager{
		pathMap:   make(mapper.Mapper[string, TemplateContext]),
		modelMap:  make(mapper.Mapper[template_type.TemplateType, []TemplateContext]),
		schemaMap: make(mapper.Mapper[canonical_type.CanonicalType, []TemplateContext]),
		funcs:     make(template.FuncMap),
	}

	args = append(args,
		TemplateFunc("DUMP", dumpJson),
		TemplateFunc("PASCAL", t.pascal),
		TemplateFunc("CAMEL", t.camel),
		TemplateFunc("PACKAGE", t.getPackage),
		TemplateFunc("BASE_TYPE", t.baseType),
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

func (t *TemplateManager) init() error {
	if t.cacher == nil {
		return fmt.Errorf("cacher is nil")
	}

	if t.pathMap == nil {
		return fmt.Errorf("pathMap is nil")
	}

	if t.funcs == nil {
		return fmt.Errorf("funcs is nil")
	}

	if t.canonicals == nil {
		return fmt.Errorf("canonicals is nil")
	}

	t.combinedTemplates = template.New("").Funcs(t.funcs)

	return nil
}
