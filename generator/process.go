package generator

import (
	cp "boundedinfinity/codegen/codegen_project"
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	lc "boundedinfinity/codegen/loader_context"

	rc "boundedinfinity/codegen/render_context"
	"fmt"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) Process() error {
	for _, lc := range t.typeManager.All() {
		var rc rc.RenderContext

		if err := t.processType(o.None[string](), lc, lc.Schema, &rc); err != nil {
			return err
		} else {
			t.rcs = append(t.rcs, rc)
		}
	}

	// for _, o := range t.projectManager.Merged. {
	// 	if rc, err := t.convertOperation(o); err != nil {
	// 		return err
	// 	} else {
	// 		t.rcs = append(t.rcs, rc)
	// 	}
	// }

	return nil
}

func (t *Generator) processType(currNs o.Option[string], lctx lc.TypeLoaderContext, schema ct.CodeGenType, rctx *rc.RenderContext) error {
	var err error
	var base rc.RenderContextBase
	found := t.typeManager.Resolve(schema)

	if found.Defined() {
		if err := t.convertBase(found.Get(), &base); err != nil {
			return err
		}
	} else {
		if err := t.convertBase(lc.TypeLoaderContext{Schema: schema}, &base); err != nil {
			return err
		}
	}

	if currNs.Empty() {
		new := path.Dir(base.SchemaNs)
		currNs = o.Some(currNs.OrElse(new))
	}

	base.CurrNs = currNs.Get()

	switch s := schema.(type) {
	case *ct.CodeGenTypeArray:
		var items rc.RenderContext
		base.SchemaType = codegen_type_id.Array

		if err = t.processType(currNs, lctx, s.Items, &items); err != nil {
			return err
		}

		*rctx = &rc.RenderContextArray{
			RenderContextBase: base,
			Items:             items,
		}
	case *ct.CodeGenTypeObject:
		obj := rc.RenderContextObject{
			RenderContextBase: base,
		}

		for _, cgtProperty := range s.Properties {
			var property rc.RenderContext

			if err = t.processType(currNs, lctx, cgtProperty, &property); err != nil {
				return err
			} else {
				obj.Properties = append(obj.Properties, property)
			}
		}

		*rctx = &obj
	case *ct.CodeGenTypeUrl:
		*rctx = &rc.RenderContextUrl{RenderContextBase: base}
	case *ct.CodeGenTypeString:
		*rctx = &rc.RenderContextString{
			RenderContextBase: base,
			Min:               s.Min,
			Max:               s.Max,
			Regex:             s.Regex,
		}
	case *ct.CodeGenTypeRef:
		found := t.typeManager.Find(s.Ref)
		var ref rc.RenderContext

		if found.Defined() {
			if err = t.processType(currNs, lctx, found.Get().Schema, &ref); err != nil {
				return err
			}
		}

		*rctx = &rc.RenderContextRef{
			RenderContextBase: base,
			Ref:               ref,
		}

	case *ct.CodeGenTypeInteger:
		*rctx = &rc.RenderContextInteger{
			RenderContextBase: base,
			Min:               s.Min,
			Max:               s.Max,
			MultipleOf:        s.MultipleOf,
		}
	case *ct.CodeGenTypeFloat:
		*rctx = &rc.RenderContextFloat{
			RenderContextBase: base,
			Min:               s.Min,
			Max:               s.Max,
			MultipleOf:        s.MultipleOf,
		}
	default:
		return fmt.Errorf("unsupported type: %v", s)
	}

	return err
}

func (t *Generator) convertBase(lctx lc.TypeLoaderContext, base *rc.RenderContextBase) error {
	fi := lctx.FileInfo
	sb := lctx.Schema.Base()
	rootNs := t.projectManager.Merged.Info.Namespace.Get()
	var schemaNs string
	var relNs string
	var name o.Option[string]

	if lctx.Schema.Base().Base().Id.Defined() {
		schemaNs = lctx.FileInfo.Source
		schemaNs = strings.Replace(schemaNs, lctx.FileInfo.Root, "", 1)
		schemaNs = path.Join(rootNs, schemaNs)
		schemaNs = extentioner.Strip(schemaNs)
		schemaNs = extentioner.Strip(schemaNs)
	}

	if schemaNs != "" {
		relNs = schemaNs
		relNs = strings.ReplaceAll(schemaNs, rootNs, "")
		relNs = strings.Replace(relNs, "/", "", 1)
	}

	name = lctx.Schema.Base().Name
	id := path.Base(lctx.Schema.Base().Id.Get())
	name = o.FirstOf(name, o.Some(id))

	*base = rc.RenderContextBase{
		Root:        fi.Root,
		Source:      fi.Source,
		SchemaType:  lctx.Schema.SchemaType(),
		CurrNs:      schemaNs,
		RootNs:      rootNs,
		SchemaNs:    schemaNs,
		RelNs:       relNs,
		MimeType:    lctx.FileInfo.MimeType,
		Id:          sb.Id.Get(),
		Name:        name.Get(),
		Description: sb.Description.Get(),
		IsPublic:    sb.Public.OrElse(true),
		IsInterface: false,
		IsRequired:  sb.Required.OrElse(false),
	}

	return nil
}

func (t *Generator) convertOperation(o *cp.CodeGenProjectOperation) error {
	// var rci rc.RenderContext
	var err error

	// base := rc.RenderContextBase{
	// 	Root:          b.Root,
	// 	Source:        b.Source,
	// 	SchemaType:    ci.SchemaType(),
	// 	RootNs:        t.projectManager.Merged.Info.Namespace.Get(),
	// 	SchemaNs:      cp.SchemaNamepace(t.projectManager.Merged.Info, ci),
	// 	RelNs:         cp.RelNamepace(t.projectManager.Merged.Info, ci),
	// 	Id:            b.Id.Get(),
	// 	Name:          o.FirstOf(b.Name, o.Some(path.Base(b.Id.Get()))).Get(),
	// 	Description:   b.Description.Get(),
	// 	IsPublic:      b.Public.OrElse(true),
	// 	IsInterface:   false,
	// 	IsRequired:    b.Required.OrElse(false),
	// 	HasValidation: ci.HasValidation(),
	// }

	// ro := rc.RenderContextOperation{
	// 	RenderContextBase: base,
	// }

	// rci = &ro

	return err
}
