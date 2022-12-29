package generator

import (
	rc "boundedinfinity/codegen/render_context"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) Process() error {
	for _, lc := range t.typeManager.All() {
		var rc rc.RenderContext

		if err := t.processType(o.None[string](), lc, lc.Schema, &rc); err != nil {
			return err
		} else {
			t.types = append(t.types, rc)
		}
	}

	for _, operation := range t.projectManager.Operations {
		var ctx rc.RenderContextOperation

		if err := t.processOperation(*operation, &ctx); err != nil {
			return err
		} else {
			t.operations = append(t.operations, ctx)
		}
	}

	return nil
}
