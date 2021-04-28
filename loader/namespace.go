package loader

func (t *Loader) processNamespace1(ctx *WalkContext) error {
	if ctx.Namespace.Input.Namespaces != nil {
		for _, child := range ctx.Namespace.Input.Namespaces {
			ctx.Namespace.Output.Children = append(ctx.Namespace.Output.Children, child.Name)
		}
	}

	return nil
}

func (t *Loader) processNamespace2(ctx *WalkContext) error {
	// specPath := ctx.Namespace.Output.SpecPath

	// if tmpls, ok := t.templateMap[specPath]; ok {
	// 	for _, tmpl := range tmpls {
	// 		t.templateMap[output.Namespace] = append(t.templateMap[output.Namespace], tmpl)
	// 	}
	// }

	return nil
}

// func (t Loader) namespaceProcssor1(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
// 	pns := path.Dir(output.Namespace)

// 	if tmpls, ok := t.templateMap[pns]; ok {
// 		for _, tmpl := range tmpls {
// 			t.templateMap[output.Namespace] = append(t.templateMap[output.Namespace], tmpl)
// 		}
// 	}

// 	if input.Templates != nil {
// 		for tmplIndex, tmpl := range input.Templates {
// 			templateWrapper := func() error {
// 				t.reportStack.Push("template[%v]", tmplIndex)
// 				defer t.reportStack.Pop()

// 				var processed model.BiInput_Template

// 				if err := t.processTemplate1(*output, tmpl, &processed); err != nil {
// 					return err
// 				}

// 				t.templateMap[output.Namespace] = append(t.templateMap[output.Namespace], processed)

// 				return nil
// 			}

// 			if err := templateWrapper(); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func (t Loader) namespaceProcssor7(input model.BiInput_Namespace, output *model.BiOutput_Namespace) error {
// 	// tmpls, ok := t.templateMap[output.Namespace]

// 	// if !ok {
// 	// 	return nil
// 	// }

// 	// namespaceTemplates := func(inputTemplate model.BiInput_Template) error {
// 	// 	namespaces := filterNamespace(t.namespaceMap, func(v model.BiOutput_Namespace) bool {
// 	// 		return v.Namespace == output.Namespace
// 	// 	})

// 	// 	for _, namescape := range namespaces {
// 	// 		outputTemplate := model.New_BiOutput_Template()

// 	// 		if err := t.processTemplate2(*output, "namespace", inputTemplate, outputTemplate); err != nil {
// 	// 			return err
// 	// 		}

// 	// 		namescape.Templates = append(namescape.Templates, outputTemplate)
// 	// 	}

// 	// 	return nil
// 	// }

// 	// operationTemplates := func(inputTemplate model.BiInput_Template) error {
// 	// 	operations := filterOperations(t.operationMap, func(v model.BiOutput_Operation) bool {
// 	// 		return v.Namespace == output.Namespace
// 	// 	})

// 	// 	for _, operation := range operations {
// 	// 		outputTemplate := model.New_BiOutput_Template()

// 	// 		if err := t.processTemplate2(*output, operation.Name, inputTemplate, outputTemplate); err != nil {
// 	// 			return err
// 	// 		}

// 	// 		operation.Templates = append(operation.Templates, outputTemplate)
// 	// 	}

// 	// 	return nil
// 	// }

// 	// modelTemplates := func(inputTemplate model.BiInput_Template) error {
// 	// 	models := filterModels(t.modelMap, func(v model.BiOutput_Model) bool {
// 	// 		return v.Namespace == output.Namespace
// 	// 	})

// 	// 	for _, model1 := range models {
// 	// 		outputTemplate := model.New_BiOutput_Template()

// 	// 		if err := t.processTemplate2(*output, model1.Name, inputTemplate, outputTemplate); err != nil {
// 	// 			return err
// 	// 		}

// 	// 		model1.Templates = append(model1.Templates, outputTemplate)
// 	// 	}

// 	// 	return nil
// 	// }

// 	// for _, tmpl := range tmpls {
// 	// 	switch tmpl.Type {
// 	// 	case string(model.TemplateType_NAMESPACE):
// 	// 		if err := namespaceTemplates(tmpl); err != nil {
// 	// 			return err
// 	// 		}
// 	// 	case string(model.TemplateType_MODEL):
// 	// 		if err := modelTemplates(tmpl); err != nil {
// 	// 			return err
// 	// 		}
// 	// 	case string(model.TemplateType_OPERATION):
// 	// 		if err := operationTemplates(tmpl); err != nil {
// 	// 			return err
// 	// 		}
// 	// 	default:
// 	// 		return t.InvalidateType()
// 	// 	}
// 	// }

// 	return nil
// }
