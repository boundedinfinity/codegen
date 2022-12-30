package generator

import (
	rc "boundedinfinity/codegen/render_context"
)

func (t *Generator) Process() error {
	typs := make([]rc.RenderContext, 0)

	for _, operation := range t.projectManager.Operations {
		var ctx rc.RenderContextOperation

		if err := t.processOperation(*operation, &ctx); err != nil {
			return err
		} else {
			t.operations = append(t.operations, ctx)
			typs = append(typs, ctx.Input, ctx.Output)
		}
	}

	err := rc.NewWalker().Base(func(s rc.RenderContext, b *rc.RenderContextBase) error {
		if b.Id != "" {
			t.types = append(t.types, s)
		}

		return nil
	}).Walk(typs...)

	return err
}
