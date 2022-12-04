package template_manager

func (t *TemplateManager) AppendTemplateContext(tc TemplateContext) {
	typ := string(tc.TemplateType)

	if !t.modelMap.Has(typ) {
		t.modelMap[typ] = make([]TemplateContext, 0)
	}

	t.modelMap[typ] = append(t.modelMap[typ], tc)
}
