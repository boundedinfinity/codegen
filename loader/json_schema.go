package loader

import (
	ct "boundedinfinity/codegen/codegen_type"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

func (t *Loader) ConvertJsonSchema(lc *ct.JsonSchemaContext, js model.JsonSchema) error {
	// if schema, err := t.convertJsonSchema(js, o.None[string]()); err != nil {
	// 	return err
	// } else {
	// 	lc.Schema = schema
	// }

	return nil
}

func (t *Loader) convertJsonSchema(v model.JsonSchema, name o.Option[string]) (ct.CodeGenType, error) {
	var base ct.CodeGenTypeBase
	var schema ct.CodeGenType

	w := jsonschema.Walk().
		Base(func(js *model.JsonSchemaBase) error {
			base = ct.CodeGenTypeBase{
				Id:          js.Id,
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			}

			return nil
		}).
		String(func(js *model.JsonSchemaString) error {
			switch js.Format.Get() {
			case stringformat.Duration:
				schema = &ct.CodeGenTypeDuration{
					CodeGenTypeBase: base,
				}
			case stringformat.Date:
				schema = &ct.CodeGenTypeDate{
					CodeGenTypeBase: base,
				}
			case stringformat.DateTime:
				schema = &ct.CodeGenTypeDateTime{
					CodeGenTypeBase: base,
				}
			case stringformat.Time:
				schema = &ct.CodeGenTypeTime{
					CodeGenTypeBase: base,
				}
			case stringformat.Email:
				schema = &ct.CodeGenTypeEmail{
					CodeGenTypeBase: base,
				}
			default:
				if js.Enum.Defined() {
					enum := ct.CodeGenTypeEnum{
						CodeGenTypeBase: base,
					}

					if js.EnumDescription.Defined() {
						for e, d := range js.EnumDescription.Get() {
							enum.Values[e] = d
						}
					}

					schema = &enum
				} else {
					schema = &ct.CodeGenTypeString{
						CodeGenTypeBase: base,
						Regex:           js.Pattern,
					}
				}
			}
			return nil
		}).
		Array(func(js *model.JsonSchemaArray) error {
			items, err := t.convertJsonSchema(js.Items.Get(), o.None[string]())

			if err != nil {
				return err
			}

			schema = &ct.CodeGenTypeArray{
				CodeGenTypeBase: base,
				Items:           items,
			}

			return nil
		}).
		Object(func(js *model.JsonSchemaObject) error {
			obj := ct.CodeGenTypeObject{
				CodeGenTypeBase: base,
			}

			for name, jsProp := range js.Properties {
				if cgProp, err := t.convertJsonSchema(jsProp, o.Some(name)); err != nil {
					return err
				} else {
					obj.Properties = append(obj.Properties, cgProp)
				}
			}

			schema = &obj
			return nil
		}).
		Ref(func(js *model.JsonSchemaRef) error {
			schema = &ct.CodeGenTypeRef{
				CodeGenTypeBase: base,
				Ref:             js.Ref,
			}
			return nil
		}).
		Integer(func(js *model.JsonSchemaInteger) error {
			schema = &ct.CodeGenTypeInteger{
				CodeGenTypeBase: base,
				Min:             js.Maximum,
				Max:             js.Maximum,
				MultipleOf:      js.MultipleOf,
			}
			return nil
		}).
		Number(func(js *model.JsonSchemaNumber) error {
			schema = &ct.CodeGenTypeFloat{
				CodeGenTypeBase: base,
				Min:             js.Maximum,
				Max:             js.Maximum,
				MultipleOf:      js.MultipleOf,
			}
			return nil
		})

	err := w.Run(v)

	return schema, err
}
