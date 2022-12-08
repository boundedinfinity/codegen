package template_manager

import (
	"boundedinfinity/codegen/canonical/canonical_type"
	"boundedinfinity/codegen/template_type"
)

func (t *TemplateManager) FindTemplateType(typ template_type.TemplateType) []TemplateContext {
	return t.modelMap.Get(typ).OrElse(make([]TemplateContext, 0))
}

func (t *TemplateManager) FindSchemaTemplate(typ canonical_type.CanonicalType) []TemplateContext {
	return t.schemaMap.Get(typ).OrElse(make([]TemplateContext, 0))
}
