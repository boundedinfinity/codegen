package loader

import (
	"boundedinfinity/codegen/model"
)

// func (t *Loader) processOperationModel(op *model.OutputOperation, input model.InputModel, output *model.OutputModel) error {
// 	// output.Name = input.Name
// 	// typ := path.Join(t.rootName(), input.Type)

// 	// if strings.HasSuffix(typ, model.COLLECTION_SUFFIX) {
// 	// 	output.Collection = true
// 	// 	typ = strings.ReplaceAll(typ, model.COLLECTION_SUFFIX, "")
// 	// }

// 	// output.Type = typ

// 	// if mt, ok := t.modelMap[output.Type]; !ok {
// 	// 	return t.ErrCustomTypeNotFound(output.Type)
// 	// } else {
// 	// 	output.Namespace = mt.Namespace
// 	// 	output.JsonStr = mt.JsonStr

// 	// 	if output.Namespace != mt.Namespace {
// 	// 		op.Imports = append(op.Imports, mt.Namespace)
// 	// 		output.Type = fmt.Sprintf("%v.%v", path.Base(mt.Namespace), mt.Name)
// 	// 	} else {
// 	// 		output.Type = mt.Name
// 	// 	}

// 	// 	if output.Collection {
// 	// 		output.Type = fmt.Sprintf("%v%v", model.COLLECTION_SUFFIX, output.Type)
// 	// 	}
// 	// }

// 	return nil
// }

func (t *Loader) processOperation1(ctx *model.WalkContext) error {
	// input := ctx.Operation.Input
	// output := ctx.Operation.Output

	// output.Name = input.Name
	// output.Description = t.splitDescription(input.Description)

	// if err := t.processOperationModel(output, input.Input, &output.Input); err != nil {
	// 	return err
	// }

	// if err := t.processOperationModel(output, input.Output, &output.Output); err != nil {
	// 	return err
	// }

	// output.Imports = util.StrSliceDedup(output.Imports)

	return nil
}

func (t *Loader) processOperation2(ctx *model.WalkContext) error {
	// output := ctx.Operation.Output
	// namespace := output.Namespace

	// if strings.HasSuffix(namespace, model.NAMESPACE_BUILTIN) {
	// 	return nil
	// }

	// vs := t.getTemplates(namespace, model.TemplateType_Operation)

	// for _, v := range vs {
	// 	outputTemplate := model.NewOutputTemplate()

	// 	if err := t.processTemplate2(ctx, output.Name, v, outputTemplate); err != nil {
	// 		return err
	// 	}

	// 	output.Templates = append(output.Templates, outputTemplate)
	// }

	return nil
}
