package generator

import "boundedinfinity/codegen/model"

func (t *Generator) runOperations() error {
	t.reportStack.Push("operations")
	if t.spec.Models == nil {
		return nil
	}
	t.reportStack.Pop()

	for i, v := range t.spec.Operations {
		if err := t.runOperation(i, v); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) runOperation(i int, v model.BiOutput_Operation) error {
	t.reportStack.Push("operations[%v]", i)

	if v.Templates != nil {
		for _, tmpl := range v.Templates {
			ctx := model.BiOutput_TemplateOperationContext{
				Operation: v,
				Spec:      t.spec,
			}

			if err := t.renderFile(tmpl, ctx); err != nil {
				return err
			}
		}
	}

	t.reportStack.Pop()
	return nil
}
