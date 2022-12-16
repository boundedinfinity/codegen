package generator

import (
	"boundedinfinity/codegen/codegen_project"
	ct "boundedinfinity/codegen/codegen_type"
	rc "boundedinfinity/codegen/render_context"
	"fmt"
	"path"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/gertd/go-pluralize"
)

func (t *Generator) Process() error {
	for _, c := range t.typeManager.All() {
		if rc, err := t.convert(c); err != nil {
			return err
		} else {
			t.rcs = append(t.rcs, rc)
		}
	}

	return nil
}

func (t *Generator) convert(ci ct.CodeGenType) (rc.RenderContext, error) {
	var rci rc.RenderContext
	var err error

	b, err := t.convertBase(ci)

	if err != nil {
		return rci, err
	}

	switch c := ci.(type) {
	case *ct.CodeGenTypeArray:
		rci, err = t.handleRenderContextArray(*c, b)
	case *ct.CodeGenTypeObject:
		rci, err = t.handleRenderContextObject(*c, b)
	case *ct.CodeGenTypeUrl:
		rci, err = t.handleRenderContextUrl(*c, b)
	case *ct.CodeGenTypeString:
		rci, err = t.handleRenderContextString(*c, b)
	case *ct.CodeGenTypeRef:
		rci, err = t.handleRenderContextRef(*c, b)
	case *ct.CodeGenTypeInteger:
		rci, err = t.handleRenderContextInteger(*c, b)
	case *ct.CodeGenTypeFloat:
		rci, err = t.handleRenderContextFloat(*c, b)
	default:
		return rci, fmt.Errorf("unsupported type: %v", c)
	}

	return rci, nil
}

func (t *Generator) handleRenderContextRef(c ct.CodeGenTypeRef, b rc.RenderContextBase) (*rc.RenderContextRef, error) {
	ref := t.typeManager.Find(c.Ref)

	if ref.Empty() {
		// TODO
	}

	rb, err := t.convert(ref.Get())

	if err != nil {
		return nil, err
	}

	rc := rc.RenderContextRef{
		RenderContextBase: b,
		Ref:               rb,
	}

	if rc.Name == "" || rc.Name == "." {
		rc.Name = rb.Base().Name
	}

	return &rc, nil
}

func (t *Generator) handleRenderContextArray(c ct.CodeGenTypeArray, b rc.RenderContextBase) (*rc.RenderContextArray, error) {
	rc := rc.RenderContextArray{
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

func (t *Generator) handleRenderContextObject(c ct.CodeGenTypeObject, b rc.RenderContextBase) (*rc.RenderContextObject, error) {
	rc := rc.RenderContextObject{
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

func (t *Generator) handleRenderContextUrl(c ct.CodeGenTypeUrl, b rc.RenderContextBase) (*rc.RenderContextUrl, error) {
	return &rc.RenderContextUrl{
		RenderContextBase: b,
	}, nil
}

func (t *Generator) handleRenderContextString(c ct.CodeGenTypeString, b rc.RenderContextBase) (*rc.RenderContextString, error) {
	return &rc.RenderContextString{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		Regex:             c.Regex,
	}, nil
}

func (t *Generator) handleRenderContextInteger(c ct.CodeGenTypeInteger, b rc.RenderContextBase) (*rc.RenderContextInteger, error) {
	return &rc.RenderContextInteger{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		MultipleOf:        c.MultipleOf,
	}, nil
}

func (t *Generator) handleRenderContextFloat(c ct.CodeGenTypeFloat, b rc.RenderContextBase) (*rc.RenderContextFloat, error) {
	return &rc.RenderContextFloat{
		RenderContextBase: b,
		Min:               c.Min,
		Max:               c.Max,
		MultipleOf:        c.MultipleOf,
	}, nil
}

func (t *Generator) convertBase(ci ct.CodeGenType) (rc.RenderContextBase, error) {
	b := ci.Base()

	return rc.RenderContextBase{
		SourceUri:     t.loader.FindSource(b.Id).Get(),
		Id:            b.Id.Get(),
		SchemaType:    ci.SchemaType(),
		RootNs:        t.projectManager.Merged.Info.Namespace.Get(),
		SchemaNs:      codegen_project.SchemaNamepace(t.projectManager.Merged.Info, ci),
		RelNs:         codegen_project.RelNamepace(t.projectManager.Merged.Info, ci),
		Name:          o.FirstOf(b.Name, o.Some(path.Base(b.Id.Get()))).Get(),
		Description:   b.Description.Get(),
		IsPublic:      b.Public.OrElse(true),
		IsInterface:   false,
		IsRequired:    b.Required.OrElse(false),
		HasValidation: ci.HasValidation(),
	}, nil
}
