package generator

// func (t *Generator) runNamespaces() error {
// 	t.reportStack.Push("namespaces")
// 	if t.spec.Namespaces == nil {
// 		return nil
// 	}
// 	t.reportStack.Pop()

// 	for i, v := range t.spec.Namespaces {
// 		if err := t.runNamespace(i, *v); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *Generator) runNamespace(i int, v model.OutputNamespace) error {
// 	t.reportStack.Push("namespaces[%v]", i)

// 	if v.Templates != nil && v.Name != "" {
// 		for _, tmpl := range v.Templates {
// 			ctx := model.OutputTemplateNamespaceContext{
// 				Namespace: v,
// 				Spec:      t.spec,
// 			}

// 			if err := t.renderFile(*tmpl, ctx); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	t.reportStack.Pop()
// 	return nil
// }
