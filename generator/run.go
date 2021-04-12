package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
	"github.com/ozgio/strutil"
)

func (t *Generator) run() error {

	if err := t.runGlobal(); err != nil {
		return err
	}

	if err := t.runComponentsSchemas(); err != nil {
		return err
	}

	if err := t.runPaths(); err != nil {
		return err
	}

	return nil
}

func (t *Generator) goPkg(inputPath string) string {
	var name string

	name = inputPath
	name = filepath.Dir(name)
	name = strings.TrimPrefix(name, t.templateRoot)
	name = path.Join(t.model.X_Bi_Go.Module.Name.Get(), name)

	return name
}

func (t *Generator) outputPath(name optional.StringOptional, abs string) (string, error) {
	var output string

	output = abs
	output = strings.TrimPrefix(output, t.templateRoot)
	output = filepath.Join(t.genRoot, output)
	output = util.TrimTemplateExt(output)

	if name.IsDefined() {
		var fn1 string
		var fn2 string

		fn1 = name.Get()
		fn1 = strings.ReplaceAll(fn1, "/", " ")
		fn1 = strutil.ToCamelCase(fn1)
		fn1 = fmt.Sprintf("%v%v", fn1, filepath.Ext(output))
		fn2 = filepath.Base(output)
		output = strings.ReplaceAll(output, fn2, fn1)
	}

	return output, nil
}

func (t *Generator) runGlobal() error {
	if util.IsNil(t.model.X_Bi_Go.Global, t.model.X_Bi_Go.Global.Templates) {
		return nil
	}

	x_bi := t.model.X_Bi_Go

	for _, tmpl := range x_bi.Global.Templates {
		var ctx model.XBiGoGlobalContext
		ctx.Model = t.model

		if abs, err := t.isTemplateExist(*tmpl); err != nil {
			return err
		} else {
			ctx.Input = abs
		}

		ctx.Package = t.goPkg(ctx.Input)

		if abs, err := t.outputPath(optional.NewStringEmpty(), ctx.Input); err != nil {
			return err
		} else {
			ctx.Output = abs
		}

		if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) runComponentsSchemas() error {
	if util.IsNil(t.model.X_Bi_Go.Components, t.model.X_Bi_Go.Components.Schemas, t.model.X_Bi_Go.Components.Schemas.Templates) {
		return nil
	}

	if util.IsNil(t.model.Components, t.model.Components.Schemas) {
		return nil
	}

	for sn, sv := range t.model.Components.Schemas {
		for _, tmpl := range t.model.X_Bi_Go.Components.Schemas.Templates {
			var ctx model.XBiGoSchemaContext
			ctx.Model = t.model
			ctx.Name = sn
			ctx.Schema = sv

			if abs, err := t.isTemplateExist(*tmpl); err != nil {
				return err
			} else {
				ctx.Input = abs
			}

			ctx.Package = t.goPkg(ctx.Input)

			if abs, err := t.outputPath(optional.NewStringValue(sn), ctx.Input); err != nil {
				return err
			} else {
				ctx.Output = abs
			}

			if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Generator) runPaths() error {
	if util.IsNil(t.model.X_Bi_Go.Paths, t.model.X_Bi_Go.Paths.Templates) {
		return nil
	}

	if util.IsNil(t.model.Paths) {
		return nil
	}

	for pn, pv := range t.model.Paths {
		for _, tmpl := range t.model.X_Bi_Go.Paths.Templates {
			var ctx model.XBiGoPathItemContext
			ctx.Model = t.model
			ctx.Name = pn
			ctx.PathItem = pv

			if abs, err := t.isTemplateExist(*tmpl); err != nil {
				return err
			} else {
				ctx.Input = abs
			}

			ctx.Package = t.goPkg(ctx.Input)

			if abs, err := t.outputPath(optional.NewStringValue(pn), ctx.Input); err != nil {
				return err
			} else {
				ctx.Output = abs
			}

			if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
				return err
			}
		}
	}

	return nil
}
