package loader

import (
	"boundedinfinity/codegen/canonical"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
)

func (t *Loader) convertJsonSchemaBase(js model.JsonSchemaBase, name o.Option[string]) canonical.CanonicalBase {
	return canonical.CanonicalBase{
		Id:          js.Id,
		Name:        o.FirstOf(name, js.Title),
		Description: js.Description,
		Public:      o.Some(true),
	}
}

func (t *Loader) convertJsonSchema(v model.JsonSchema, name o.Option[string]) (canonical.Canonical, error) {
	var can canonical.Canonical
	var err error

	switch js := v.(type) {
	case *model.JsonSchemaString:
		switch js.Format.Get() {
		case stringformat.Duration:
			can = canonical.CanonicalDuration{
				CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Date:
			can = canonical.CanonicalDate{
				CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.DateTime:
			can = canonical.CanonicalDateTime{
				CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Time:
			can = canonical.CanonicalTime{
				CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		case stringformat.Email:
			can = canonical.CanonicalEmail{
				CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			}
		default:
			if js.Enum.Defined() {
				enum := canonical.CanonicalEnum{
					CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
				}

				if js.EnumDescription.Defined() {
					for e, d := range js.EnumDescription.Get() {
						enum.Values[e] = d
					}
				} else {

				}
				can = enum
			} else {
				can = canonical.CanonicalString{
					CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
					Regex:         js.Pattern,
				}
			}
		}

	case *model.JsonSchemaArray:
		items, err := t.convertJsonSchema(js.Items.Get(), o.None[string]())

		arr := canonical.CanonicalArray{
			CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Items:         items,
		}

		if err != nil {
			return can, err
		}

		can = arr
	case *model.JsonSchemaObject:
		obj := canonical.CanonicalObject{
			CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
		}

		for name, jsProp := range js.Properties {
			if canonicalProp, err := t.convertJsonSchema(jsProp, o.Some(name)); err != nil {
				return can, err
			} else {
				obj.Properties = append(obj.Properties, canonicalProp)
			}
		}

		can = obj
	case *model.JsonSchemaRef:
		can = canonical.CanonicalRef{
			CanonicalBase: t.convertJsonSchemaBase(js.JsonSchemaBase, name),
			Ref:           js.Ref,
		}
	}

	return can, err
}
