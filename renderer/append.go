package renderer

func (t *Renderer) AppendTemplateContext(tc TemplateContext) {
	// if !t.templateType2Context.Has(tc.TemplateType) {
	// 	t.templateType2Context[tc.TemplateType] = make([]TemplateContext, 0)
	// }

	// t.templateType2Context[tc.TemplateType] = append(t.templateType2Context[tc.TemplateType], tc)

	// if tc.ModelType.Defined() {
	// 	if !t.typeId2Context.Has(tc.ModelType.Get()) {
	// 		t.typeId2Context[tc.ModelType.Get()] = make([]TemplateContext, 0)
	// 	}

	// 	t.typeId2Context[tc.ModelType.Get()] = append(t.typeId2Context[tc.ModelType.Get()], tc)
	// }
}
