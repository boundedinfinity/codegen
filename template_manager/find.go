package template_manager

import (
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/template_type"
)

func (t *TemplateManager) FindTemplateType(typ template_type.TemplateType) []TemplateContext {
	return t.modelMap.Get(typ).OrElse(make([]TemplateContext, 0))
}

func (t *TemplateManager) FindSchemaTemplate(typ codegen_type_id.CodgenTypeId) []TemplateContext {
	return t.schemaMap.Get(typ).OrElse(make([]TemplateContext, 0))
}
