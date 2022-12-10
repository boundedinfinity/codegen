package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) Process() error {
	for _, c := range t.canonicals.All() {
		if rc, err := t.convert(c); err != nil {
			return err
		} else {
			t.rcs = append(t.rcs, rc)
		}
	}

	return nil
}

func (t *Generator) convert(ci canonical.Canonical) (render_context.RenderContext, error) {
	var rci render_context.RenderContext
	b, err := t.convertBase(ci)

	if err != nil {
		return rci, err
	}

	switch c := ci.(type) {
	case canonical.CanonicalArray:
		rc := render_context.RenderContextArray{
			RenderContextBase: b,
		}

		if i, err := t.convert(c.Items); err != nil {
			return rci, err
		} else {
			rc.Items = i
		}

		rci = rc
	case canonical.CanonicalObject:
		rc := render_context.RenderContextObject{
			RenderContextBase: b,
		}

		for _, cp := range c.Properties {
			if rcp, err := t.convert(cp); err != nil {
				return rci, err
			} else {
				rc.Properties = append(rc.Properties, rcp)
			}
		}

		rci = rc
	case canonical.CanonicalUrl:
		rc := render_context.RenderContextUrl{
			RenderContextBase: b,
		}
		rci = rc
	case *canonical.CanonicalUrl:
		rc := render_context.RenderContextUrl{
			RenderContextBase: b,
		}
		rci = rc
	case canonical.CanonicalString:
		rc := render_context.RenderContextString{
			RenderContextBase: b,
			Min:               c.Min,
			Max:               c.Max,
			Regex:             c.Regex,
		}

		rci = rc
	default:
		return rci, fmt.Errorf("unsupported type: %v", c)
	}

	return rci, nil
}

func (t *Generator) convertBase(ci canonical.Canonical) (render_context.RenderContextBase, error) {
	b := ci.Base()

	return render_context.RenderContextBase{
		SourceUri:   t.loader.FindSource(b.Id).Get(),
		SchemaType:  string(ci.SchemaType()),
		RootNs:      t.codeGenSchema.Info.Namespace.Get(),
		SchemaNs:    util.SchemaNamepace(t.codeGenSchema.Info, ci),
		Name:        o.FirstOf(b.Name, o.Some(path.Base(b.Id.Get()))).Get(),
		Description: b.Description.Get(),
		IsPublic:    true,
		IsInterface: false,
		IsRequired:  b.Required.Get(),
	}, nil
}
