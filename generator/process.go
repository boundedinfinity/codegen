package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/render_context"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/gertd/go-pluralize"
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
	var err error

	b, err := t.convertBase(ci)

	if err != nil {
		return rci, err
	}

	switch c := ci.(type) {
	case canonical.CanonicalArray:
		rci, err = t.handleRenderContextArray(c, b)
	case *canonical.CanonicalArray:
		rci, err = t.handleRenderContextArray(*c, b)
	case canonical.CanonicalObject:
		rci, err = t.handleRenderContextObject(c, b)
	case *canonical.CanonicalObject:
		rci, err = t.handleRenderContextObject(*c, b)
	case canonical.CanonicalUrl:
		rci, err = t.handleRenderContextUrl(c, b)
	case *canonical.CanonicalUrl:
		rci, err = t.handleRenderContextUrl(*c, b)
	case canonical.CanonicalString:
		rci, err = t.handleRenderContextString(c, b)
	case *canonical.CanonicalString:
		rci, err = t.handleRenderContextString(*c, b)
	case canonical.CanonicalRef:
		rci, err = t.handleRenderContextRef(c, b)
	case *canonical.CanonicalRef:
		rci, err = t.handleRenderContextRef(*c, b)
	case canonical.CanonicalInteger:
		rci, err = t.handleRenderContextInteger(c, b)
	case *canonical.CanonicalInteger:
		rci, err = t.handleRenderContextInteger(*c, b)
	case canonical.CanonicalFloat:
		rci, err = t.handleRenderContextFloat(c, b)
	case *canonical.CanonicalFloat:
		rci, err = t.handleRenderContextFloat(*c, b)
	default:
		return rci, fmt.Errorf("unsupported type: %v", c)
	}

	return rci, nil
}

func (t *Generator) handleRenderContextRef(c canonical.CanonicalRef, b render_context.RenderContextBase) (*render_context.RenderContextRef, error) {
	ref := t.canonicals.Find(c.Ref)

	if ref.Empty() {
		// TODO
	}

	rb, err := t.convert(ref.Get())

	if err != nil {
		return nil, err
	}

	rc := render_context.RenderContextRef{
		RenderContextBase: b,
		Ref:               rb,
	}

	if rc.Name == "" || rc.Name == "." {
		rc.Name = rb.Base().Name
	}

	return &rc, nil
}

func (t *Generator) handleRenderContextArray(c canonical.CanonicalArray, b render_context.RenderContextBase) (*render_context.RenderContextArray, error) {
	rc := render_context.RenderContextArray{
		RenderContextBase: b,
	}

	if i, err := t.convert(c.Items); err != nil {
		return &rc, err
	} else {
		if rc.SchemaNs == "" {
			rc.SchemaNs = i.Base().SchemaNs
		}

		if rc.Name == "" || rc.Name == "." {
			rc.Name = i.Base().Name
			rc.Name = pluralize.NewClient().Plural(rc.Name)
		}

		rc.Items = i
	}

	return &rc, nil
}

func (t *Generator) handleRenderContextObject(c canonical.CanonicalObject, b render_context.RenderContextBase) (*render_context.RenderContextObject, error) {
	rc := render_context.RenderContextObject{
		RenderContextBase: b,
	}

	for _, cp := range c.Properties {
		if rcp, err := t.convert(cp); err != nil {
			return &rc, err
		} else {
			rc.Properties = append(rc.Properties, rcp)
		}
	}

	return &rc, nil
}

func (t *Generator) handleRenderContextUrl(c canonical.CanonicalUrl, b render_context.RenderContextBase) (*render_context.RenderContextUrl, error) {
	return &render_context.RenderContextUrl{
		RenderContextBase: b,
	}, nil
}

func (t *Generator) handleRenderContextString(c canonical.CanonicalString, b render_context.RenderContextBase) (*render_context.RenderContextString, error) {
	return &render_context.RenderContextString{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		Regex:             c.Regex,
	}, nil
}

func (t *Generator) handleRenderContextInteger(c canonical.CanonicalInteger, b render_context.RenderContextBase) (*render_context.RenderContextInteger, error) {
	return &render_context.RenderContextInteger{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		MultipleOf:        c.MultipleOf,
	}, nil
}

func (t *Generator) handleRenderContextFloat(c canonical.CanonicalFloat, b render_context.RenderContextBase) (*render_context.RenderContextFloat, error) {
	return &render_context.RenderContextFloat{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		MultipleOf:        c.MultipleOf,
	}, nil
}

func (t *Generator) convertBase(ci canonical.Canonical) (render_context.RenderContextBase, error) {
	b := ci.Base()

	return render_context.RenderContextBase{
		SourceUri:     t.loader.FindSource(b.Id).Get(),
		Id:            b.Id.Get(),
		SchemaType:    string(ci.SchemaType()),
		RootNs:        t.codeGenSchema.Info.Namespace.Get(),
		SchemaNs:      util.SchemaNamepace(t.codeGenSchema.Info, ci),
		Name:          o.FirstOf(b.Name, o.Some(path.Base(b.Id.Get()))).Get(),
		Description:   b.Description.Get(),
		IsPublic:      b.Public.OrElse(true),
		IsInterface:   false,
		IsRequired:    b.Required.OrElse(false),
		HasValidation: ci.HasValidation(),
	}, nil
}
