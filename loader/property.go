package loader

import (
	"boundedinfinity/codegen/model"
)

func (t *Loader) processProperty1(ctx *model.WalkContext) error {
	// for _, property := range ctx.Model.Properties {
	// 	if model.IsSchemaTypeEnum(property.Type) {
	// 		pType, err := model.SchemaTypeEnumParse(property.Type)

	// 		if err != nil {
	// 			return err
	// 		}

	// 		switch pType {
	// 		case model.SchemaType_Boolean:

	// 		}
	// 	}

	// 	// pType := t.getModelType(property.Type)
	// }

	return nil
}
func (t *Loader) processProperty2(ctx *model.WalkContext) error {
	// st := ctx.Property.Output.Type
	// pm, ok := t.modelMap[ctx.Property.Output.Type]

	// if !ok {
	// 	return t.ErrInvalidType(st)
	// }

	// output := ctx.Property.Output

	// if strings.HasSuffix(pm.Namespace, model.NAMESPACE_BUILTIN) {
	// 	output.Type = pm.Type
	// } else if ctx.Model.Output.Namespace != pm.Namespace {
	// 	output.Type = fmt.Sprintf("%v.%v", path.Base(pm.Namespace), pm.Name)
	// 	ctx.Model.Output.Imports = append(ctx.Model.Output.Imports, pm.Namespace)
	// } else {
	// 	output.Type = pm.Name
	// }

	// if output.Collection {
	// 	output.Type = fmt.Sprintf("%v%v", model.COLLECTION_SUFFIX, output.Type)
	// }

	// ctx.Model.Output.Imports = util.StrSliceDedup(ctx.Model.Output.Imports)

	return nil
}
