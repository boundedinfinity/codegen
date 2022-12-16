package loader

import (
	ct "boundedinfinity/codegen/codegen_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

func (t *Loader) ConvertJsonSchema() error {
	for _, js := range t.jsonSchemas.AllPath() {
		if schema, err := t.convertJsonSchema(js, o.None[string]()); err != nil {
			return err
		} else {
			source := t.jsonSchemas.GetSource(js.Base().Id.Get())
			root := t.jsonSchemas.GetRoot(js.Base().Id.Get())

			if err := t.typeManager.Register(root.Get(), source.Get(), schema); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) convertJsonSchemaBase(js model.JsonSchemaBase, name o.Option[string]) ct.CodeGenTypeBase {
	return ct.CodeGenTypeBase{
		Id:          js.Id,
		Name:        o.FirstOf(name, js.Title),
		Description: js.Description,
		Public:      o.Some(true),
	}
}

func (t *Loader) convertJsonSchema(v model.JsonSchema, name o.Option[string]) (ct.CodeGenType, error) {
	var can ct.CodeGenType
	var err error

	switch js := v.(type) {
	case *model.JsonSchemaString:
		switch js.Format.Get() {
		case stringformat.Duration:
			can = &ct.CodeGenTypeDuration{
				CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Date:
			can = &ct.CodeGenTypeDate{
				CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.DateTime:
			can = &ct.CodeGenTypeDateTime{
				CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Time:
			can = &ct.CodeGenTypeTime{
				CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Email:
			can = &ct.CodeGenTypeEmail{
				CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		default:
			if js.Enum.Defined() {
				enum := ct.CodeGenTypeEnum{
					CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
				}

				if js.EnumDescription.Defined() {
					for e, d := range js.EnumDescription.Get() {
						enum.Values[e] = d
					}
				} else {

				}
				can = &enum
			} else {
				can = &ct.CodeGenTypeString{
					CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
					Regex:           js.Pattern,
				}
			}
		}
	case *model.JsonSchemaArray:
		items, err := t.convertJsonSchema(js.Items.Get(), o.None[string]())

		arr := ct.CodeGenTypeArray{
			CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Items:           items,
		}

		if err != nil {
			return can, err
		}

		can = &arr
	case *model.JsonSchemaObject:
		obj := ct.CodeGenTypeObject{
			CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
		}

		for name, jsProp := range js.Properties {
			if cgProp, err := t.convertJsonSchema(jsProp, o.Some(name)); err != nil {
				return can, err
			} else {
				obj.Properties = append(obj.Properties, cgProp)
			}
		}

		can = &obj
	case *model.JsonSchemaRef:
		can = &ct.CodeGenTypeRef{
			CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Ref:             js.Ref,
		}
	case *model.JsonSchemaInteger:
		can = &ct.CodeGenTypeInteger{
			CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Min:             js.Maximum,
			Max:             js.Maximum,
			MultipleOf:      js.MultipleOf,
		}
	case *model.JsonSchemaNumber:
		can = &ct.CodeGenTypeFloat{
			CodeGenTypeBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Min:             js.Maximum,
			Max:             js.Maximum,
			MultipleOf:      js.MultipleOf,
		}
	}

	return can, err
}
