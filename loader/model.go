package loader

func (t *Loader) processModel1(ctx *WalkContext) error {
	ctx.Model.Output.Name = ctx.Model.Input.Name
	ctx.Model.Output.Description = t.splitDescription(ctx.Model.Input.Description)

	if _, ok := t.dependencies[ctx.Model.Output.SpecPath]; !ok {
		t.dependencies[ctx.Model.Output.SpecPath] = NewNode(ctx.Model.Output.SpecPath)
	}

	return nil
}

func (t *Loader) processModel2(ctx *WalkContext) error {

	return nil
}

// func (t *Loader) modelProcessor8(namespace model.BiOutput_Namespace, input model.BiInput_Model, output *model.BiOutput_Model) error {
// 	if output.Properties == nil {
// 		return nil
// 	}

// 	m := make(map[string]bool)

// 	for _, property := range output.Properties {
// 		if !strings.HasPrefix(property.Namespace, model.NAMESPACE_BUILTIN) && property.Namespace != namespace.Namespace {
// 			if _, ok := m[property.Namespace]; !ok {
// 				m[property.Namespace] = true
// 			}
// 		}
// 	}

// 	for k := range m {
// 		output.Imports = append(output.Imports, k)
// 	}

// 	return nil
// }
