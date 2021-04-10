package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

func (t *Generator) generateGo() error {
	if t.rc.Model.X_Bi_Go == nil {
		return nil
	}

	if t.rc.Model.X_Bi_Go.Module == nil {
		return model.CannotBeEmptyErr
	}

	if t.rc.Model.X_Bi_Go.Templates != nil {
		for _, tmpl := range t.rc.Model.X_Bi_Go.Templates {
			var rt model.XBiGoGlobalRuntime

			if err := t.goNonGoTemplate2GoGlobalRuntime(tmpl, &rt); err != nil {
				return nil
			}

			if err := util.RenderFile(rt.Input.Get(), rt.Output.Get(), rt.Context); err != nil {
				return err
			}
		}
	}

	if t.rc.Model.Components == nil || t.rc.Model.Components.Schemas == nil ||
		t.rc.Model.Components.X_Bi_Go == nil ||
		t.rc.Model.Components.X_Bi_Go.Schemas == nil ||
		t.rc.Model.Components.X_Bi_Go.Schemas.Templates == nil {
		return nil
	}

	for sn, sv := range t.rc.Model.Components.Schemas {
		for _, tmpl := range t.rc.Model.Components.X_Bi_Go.Schemas.Templates {
			var rt model.XBiGoSchemaRuntime

			if err := t.goSchema2GoContext(&rt, tmpl, sn, sv); err != nil {
				return nil
			}

			if err := util.RenderFile(rt.Input.Get(), rt.Output.Get(), rt.Context); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Generator) fullPkg2FilePkg(pkg optional.StringOptional) (optional.StringOptional, error) {
	if pkg.IsEmpty() {
		return pkg, model.CannotBeEmptyErr
	}

	var pkg2 string
	mn := t.rc.Model.X_Bi_Go.Module.Name.Get()
	pkg2 = pkg.Get()
	pkg2 = strings.TrimPrefix(pkg2, mn)
	pkg2 = strings.TrimPrefix(pkg2, "/")

	if pkg2 == "" {
		pkg2 = "main"
	}

	return optional.NewStringValue(pkg2), nil
}

func (t *Generator) goSchema2GoContext(r *model.XBiGoSchemaRuntime, tmpl *model.X_Bi_GoTemplate, sn string, sv model.JsonSchema_Draft07) error {
	r.Context.Model = t.rc.Model
	r.Context.Name = optional.NewStringValue(sn)
	r.Context.Schema = sv

	if abs, ok := util.FileSearch(tmpl.Input, t.rc.Model.X_Bi_Go.TemplateRoot); ok {
		r.Input = optional.NewStringValue(abs)
	} else {
		return model.CannotBeEmptyErr
	}

	if sv.X_Bi_Go.Package.IsDefined() {
		if pkg, err := t.fullPkg2FilePkg(sv.X_Bi_Go.Package); err != nil {
			return err
		} else {
			r.Context.Package = pkg
		}
	} else if tmpl.Package.IsDefined() {
		if pkg, err := t.fullPkg2FilePkg(tmpl.Package); err != nil {
			return err
		} else {
			r.Context.Package = pkg
		}
	}

	var o string
	o = tmpl.Input.Get()
	o = filepath.Base(o)
	o = strings.TrimSuffix(o, filepath.Ext(o))
	o = fmt.Sprintf("%v.%v", sn, o)
	o = filepath.Join(r.Context.Package.Get(), o)
	o = filepath.Join(t.rc.Model.X_Bi_Go.GenRoot.Get(), o)

	r.Output = optional.NewStringValue(o)

	return nil
}

func (t *Generator) goNonGoTemplate2GoGlobalRuntime(tmpl *model.X_Bi_GoTemplate, r *model.XBiGoGlobalRuntime) error {
	id := filepath.Dir(t.rc.Model.X_Bi_Go.TemplateRoot.Get())

	if abs, ok := util.FileSearch(tmpl.Input, optional.NewStringValue(id)); ok {
		r.Input = optional.NewStringValue(abs)
	} else {
		return model.CannotBeEmptyErr
	}

	if tmpl.Package.IsDefined() {
		if pkg, err := t.fullPkg2FilePkg(tmpl.Package); err != nil {
			return err
		} else {
			r.Context.Package = pkg
		}
	}

	if tmpl.Output.IsDefined() {
		if abs, err := util.AbsFromDirPath(t.rc.Model.X_Bi_Go.GenRoot, tmpl.Output); err != nil {
			return err
		} else {
			r.Output = abs
		}
	} else {
		var o string
		var ie string

		o = r.Input.Get()
		o = filepath.Base(o)
		ie = filepath.Ext(o)
		o = strings.TrimSuffix(o, ie)
		o = filepath.Join(t.rc.Model.X_Bi_Go.GenRoot.Get(), o)

		if r.Context.Package.Get() != "main" {
			o = filepath.Join(o, r.Context.Package.Get())
		}

		r.Output = optional.NewStringValue(o)
	}

	r.Context.Model = t.rc.Model
	return nil
}
