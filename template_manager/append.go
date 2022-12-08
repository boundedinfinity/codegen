package template_manager

func (t *TemplateManager) AppendTemplateContext(tc TemplateContext) {
	if !t.modelMap.Has(tc.TemplateType) {
		t.modelMap[tc.TemplateType] = make([]TemplateContext, 0)
	}

	t.modelMap[tc.TemplateType] = append(t.modelMap[tc.TemplateType], tc)

	if tc.ModelType.Defined() {
		if !t.schemaMap.Has(tc.ModelType.Get()) {
			t.schemaMap[tc.ModelType.Get()] = make([]TemplateContext, 0)
		}

		t.schemaMap[tc.ModelType.Get()] = append(t.schemaMap[tc.ModelType.Get()], tc)
	}
}
