package system

import (
	"boundedinfinity/codegen/canonical"
	"path/filepath"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-urischemer"
)

func (t *System) Convert() error {
	for _, v := range t.jsonSchemas.All() {
		if can, err := t.convert(v, o.None[string]()); err != nil {
			return err
		} else {
			t.canonicals.Register(can)
		}
	}

	return nil
}

func (t *System) createImport(v model.JsonSchema) (string, error) {
	id := t.jsonSchemas.Id(v).Get()
	_, path, err := urischemer.Break(id)

	if err != nil {
		return "", err
	}

	pkg := t.mergedCodeGen.Info.Package.Get()
	pkg = filepath.Join(pkg, path)

	return pkg, nil
}

func (t *System) convert(v model.JsonSchema, name o.Option[string]) (canonical.Canonical, error) {
	var can canonical.Canonical
	var err error

	path, err := t.createImport(v)

	if err != nil {
		return can, err
	}

	switch js := v.(type) {
	case *model.JsonSchemaString:
		can = canonical.CanonicalString{
			CanonicalBase: canonical.CanonicalBase{
				Id:          o.Some(path),
				Name:        js.Title,
				Description: js.Description,
				Public:      o.Some(true),
			},
		}
	case *model.JsonSchemaArray:
		i, err := t.convert(js.Items.Get(), o.None[string]())

		x := canonical.CanonicalArray{
			CanonicalBase: canonical.CanonicalBase{
				Id:          o.Some(path),
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			},
			Items: i,
		}

		if err != nil {
			return can, err
		}

		can = x
	case *model.JsonSchemaObject:
		x := canonical.CanonicalObject{
			CanonicalBase: canonical.CanonicalBase{
				Id:          o.Some(path),
				Name:        o.FirstOf(name, js.Title),
				Description: js.Description,
				Public:      o.Some(true),
			},
		}

		for name, property := range js.Properties {
			if cprop, err := t.convert(property, o.Some(name)); err != nil {
				return can, err
			} else {
				x.Properties = append(x.Properties, cprop)
			}
		}

		can = x
	}

	return can, err
}
