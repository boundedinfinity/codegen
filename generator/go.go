package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

func (t *Generator) searchFromTemplateRoot(p optional.StringOptional) (optional.StringOptional, error) {
	if abs, ok := util.FileSearch(p, t.schema.X_Bi_Go.TemplateRoot); !ok {
		return p, model.CannotBeEmptyErr
	} else {
		return optional.NewStringValue(abs), nil
	}
}

func (t *Generator) absFromGenerationRoot(p optional.StringOptional) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, model.CannotBeEmptyErr
	}

	if filepath.IsAbs(p.Get()) {
		return p, nil
	}

	p2 := filepath.Join(t.schema.X_Bi_Go.GenRoot.Get(), p.Get())
	abs, err := filepath.Abs(p2)

	if err != nil {
		return p, err
	}

	return optional.NewStringValue(abs), nil
}

func (t *Generator) generateGo() error {
	if util.IsNil(t.schema.X_Bi_Go) {
		return nil
	}

	xbigo := t.schema.X_Bi_Go

	if abs, err := t.searchFromSchemaPath(xbigo.TemplateRoot); err != nil {
		return err
	} else {
		xbigo.TemplateRoot = abs
	}

	if xbigo.GenRoot.IsDefined() {
		genRoot := xbigo.GenRoot.Get()
		if !filepath.IsAbs(genRoot) {
			p1 := filepath.Join(t.path, genRoot)
			if p2, err := filepath.Abs(p1); err != nil {
				return t.generatorSchemaErr(err, "x-bi-go", "genRoot")
			} else {
				xbigo.GenRoot = optional.NewStringValue(p2)
			}
		}
	} else {
		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "genRoot")
	}

	if util.IsNil(xbigo.Module) {
		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "module")
	}

	if xbigo.Module.Name.IsEmpty() {
		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "module", "name")
	}

	if xbigo.Module.Version.IsEmpty() {
		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "module", "version")
	}

	if util.IsDef(xbigo.Global, xbigo.Global.Templates) {
		for _, tmpl := range xbigo.Global.Templates {
			var rt model.XBiGoGlobalRuntime

			if err := t.createGlobalRuntime(&rt, *tmpl); err != nil {
				return err
			}

			if err := util.RenderFile(rt.Input.Get(), rt.Output.Get(), rt.Context); err != nil {
				return err
			}
		}
	}

	if util.IsDef(xbigo.Components, xbigo.Components.Schemas, xbigo.Components.Schemas.Templates) {
		// for _, tmpl := range xbigo.Components.Schemas.Templates {
		//     var rt model.XBiGoSchemaRuntime

		// }
	}

	return nil
}

func (t *Generator) createGlobalRuntime(rt *model.XBiGoGlobalRuntime, tmpl model.X_Bi_Go_Template) error {
	rt.Context.Schema = t.schema

	if abs, err := t.searchFromTemplateRoot(tmpl.Input); err != nil {
		return err
	} else {
		rt.Input = abs
	}

	if tmpl.Package.IsEmpty() {
		rt.Context.Package = t.goPkgBase(t.goTemplateFullPkg(tmpl))
	}

	if abs, err := t.getOutput(tmpl); err != nil {
		return err
	} else {
		rt.Output = abs
	}

	return nil
}

func (t *Generator) getOutput(tmpl model.X_Bi_Go_Template) (optional.StringOptional, error) {
	var output optional.StringOptional

	if tmpl.Output.IsDefined() {
		if !filepath.IsAbs(tmpl.Output.Get()) {
			if abs, err := t.absFromGenerationRoot(tmpl.Output); err != nil {
				return output, err
			} else {
				output = abs
			}
		}
	} else {
		var o string
		var e string
		var p string

		o = tmpl.Input.Get()
		o = filepath.Base(o)
		e = filepath.Ext(o)
		o = strings.TrimSuffix(o, e)
		p = t.goTemplateFullPkg(tmpl).Get()
		p = t.goPkgRelative(optional.NewStringValue(p)).Get()
		o = filepath.Join(p, o)

		if abs, err := t.absFromGenerationRoot(optional.NewStringValue(o)); err != nil {
			return output, err
		} else {
			output = abs
		}
	}

	return output, nil
}

func (t *Generator) goTemplateFullPkg(tmpl model.X_Bi_Go_Template) optional.StringOptional {
	if tmpl.Package.IsDefined() {
		return tmpl.Package
	} else {
		return t.schema.X_Bi_Go.Module.Name
	}
}

func (t *Generator) goPkgBase(pgk optional.StringOptional) optional.StringOptional {
	var x string
	x = pgk.Get()
	x = path.Base(x)
	return optional.NewStringValue(x)
}

func (t *Generator) goPkgRelative(pgk optional.StringOptional) optional.StringOptional {
	var x string
	x = strings.TrimPrefix(t.schema.X_Bi_Go.Module.Name.Get(), pgk.Get())
	x = strings.TrimPrefix(x, "/")
	return optional.NewStringValue(x)
}
