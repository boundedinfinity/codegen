package loader

import (
	"boundedinfinity/codegen/canonical"
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-jsonschema/stringformat"
	"github.com/boundedinfinity/go-urischemer"
)

func (t *Loader) createId(v model.JsonSchema) (o.Option[string], error) {
	id := t.jsonSchemas.Id(v)

	if id.Empty() {
		return id, nil
	}

	_, path, err := urischemer.Break(id.Get())

	if err != nil {
		return id, err
	}

	pkg := t.mergedCodeGen.Info.Package.Get()
	pkg = filepath.Join(pkg, path)

	return o.Some(pkg), nil
}

func (t *Loader) convert(v model.JsonSchema, name o.Option[string]) (canonical.Canonical, error) {
	var can canonical.Canonical
	var err error

	id, err := t.createId(v)

	if err != nil {
		return can, err
	}

	switch js := v.(type) {
	case *model.JsonSchemaString:
		switch js.Format.Get() {
		case stringformat.Duration:
			can = canonical.CanonicalDuration{
				CanonicalBase: canonical.CanonicalBase{
					Id:          id,
					Name:        o.FirstOf(name, js.Title),
					Description: js.Description,
					Public:      o.Some(true),
				},
			}
		case stringformat.Date:
			can = canonical.CanonicalDate{
				CanonicalBase: canonical.CanonicalBase{
					Id:          id,
					Name:        o.FirstOf(name, js.Title),
					Description: js.Description,
					Public:      o.Some(true),
				},
			}
		default:
			if js.Enum.Defined() {
				enum := canonical.CanonicalEnum{
					CanonicalBase: canonical.CanonicalBase{
						Id:          id,
						Name:        o.FirstOf(name, js.Title),
						Description: js.Description,
						Public:      o.Some(true),
					},
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
					CanonicalBase: canonical.CanonicalBase{
						Id:          id,
						Name:        o.FirstOf(name, js.Title),
						Description: js.Description,
						Public:      o.Some(true),
					},
					Regex: js.Pattern,
				}
			}
		}

	case *model.JsonSchemaArray:
		items, err := t.convert(js.Items.Get(), o.None[string]())

		arr := canonical.CanonicalArray{
			CanonicalBase: canonical.CanonicalBase{
				Id:          id,
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			},
			Items: items,
		}

		if err != nil {
			return can, err
		}

		can = arr
	case *model.JsonSchemaObject:
		obj := canonical.CanonicalObject{
			CanonicalBase: canonical.CanonicalBase{
				Id:          id,
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			},
		}

		for name, jsProp := range js.Properties {
			if canonicalProp, err := t.convert(jsProp, o.Some(name)); err != nil {
				return can, err
			} else {
				obj.Properties = append(obj.Properties, canonicalProp)
			}
		}

		can = obj
	case *model.JsonSchemaRef:
		can = canonical.CanonicalRef{
			CanonicalBase: canonical.CanonicalBase{
				Id:          id,
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			},
			Ref: js.Ref,
		}
	}

	return can, err
}
