package generator

// import "boundedinfinity/codegen/model"

// func (t *Generator) runModels() error {
// 	t.reportStack.Push("models")
// 	if t.spec.Models == nil {
// 		return nil
// 	}
// 	t.reportStack.Pop()

// 	for i, m := range t.spec.Models {
// 		if err := t.runModel(i, *m); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *Generator) runModel(i int, m model.OutputModel) error {
// 	t.reportStack.Push("models[%v]", i)

// 	if m.Templates != nil {
// 		for _, tmpl := range m.Templates {
// 			ctx := model.OutputTemplateModelContext{
// 				Model: m,
// 				Spec:  t.spec,
// 			}

// 			if err := t.renderFile(*tmpl, ctx); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	t.reportStack.Pop()
// 	return nil
// }
