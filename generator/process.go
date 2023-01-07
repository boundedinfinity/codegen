package generator

func (t *Generator) Process() error {
	// typs := make([]ct.CodeGenProjectOperation, 0)

	// for _, operation := range t.projectManager.Operations {
	// 	var ctx ct.CodeGenProjectOperation

	// 	if err := t.processOperation(*operation, &ctx); err != nil {
	// 		return err
	// 	} else {
	// 		t.operations = append(t.operations, ctx)
	// 		typs = append(typs, ctx.Input, ctx.Output)
	// 	}
	// }

	// err := rc.NewWalker().Base(func(s rc.RenderContext, b *rc.RenderContextBase) error {
	// 	if b.Id != "" {
	// 		t.types = append(t.types, s)
	// 	}

	// 	return nil
	// }).Walk(typs...)

	return nil
}
