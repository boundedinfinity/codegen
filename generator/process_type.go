package generator

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"

	rc "boundedinfinity/codegen/render_context"
	"fmt"
	"path"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Generator) processType(currNs o.Option[string], lctx ct.CodeGenTypeContext, schema ct.CodeGenType, rctx *rc.RenderContext) error {
	var err error
	var base rc.RenderContextBase
	found := t.typeManager.Resolve(schema)

	if found.Defined() {
		if err := t.convertBase(found.Get(), &base); err != nil {
			return err
		}
	} else {
		if err := t.convertBase(ct.CodeGenTypeContext{Schema: schema}, &base); err != nil {
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

func (t *Generator) convertBase(lctx ct.CodeGenTypeContext, base *rc.RenderContextBase) error {
	fileInfo := lctx.FileInfo
	schemaBase := lctx.Schema.Base()
	rootNs := t.projectManager.Merged.Info.Namespace.Get()
	var schemaNs string
	var relNs string
	var name o.Option[string]

	if schemaBase.Id.Defined() {
		schemaNs = fileInfo.Source
		schemaNs = strings.Replace(schemaNs, fileInfo.Root, "", 1)
		schemaNs = path.Join(rootNs, schemaNs)
		schemaNs = extentioner.Strip(schemaNs)
		schemaNs = extentioner.Strip(schemaNs)
	}

	if schemaNs != "" {
		relNs = schemaNs
		relNs = strings.ReplaceAll(schemaNs, rootNs, "")
		relNs = strings.Replace(relNs, "/", "", 1)
	}

	name = schemaBase.Name
	id := path.Base(schemaBase.Id.Get())
	name = o.FirstOf(name, o.Some(id))

	*base = rc.RenderContextBase{
		Namespace: ct.Namespace{
			RootNs:   rootNs,
			SchemaNs: schemaNs,
			RelNs:    relNs,
			CurrNs:   schemaNs,
		},
		FileInfo:    fileInfo,
		SchemaType:  lctx.Schema.SchemaType(),
		Id:          schemaBase.Id.Get(),
		Name:        name.Get(),
		Description: schemaBase.Description.Get(),
		IsPublic:    schemaBase.Public.OrElse(true),
		IsInterface: false,
		IsRequired:  schemaBase.Required.OrElse(false),
	}

	return nil
}
