package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
	"github.com/ozgio/strutil"
)

func (t *Generator) run() error {
	if util.IsDef(t.spec.Generation, t.spec.Generation.Global, t.spec.Generation.Global.Templates) {
		for i, tmpl := range t.spec.Generation.Global.Templates {
			errpath := []string{"generation", "global", fmt.Sprintf("templates[%v]", i)}

			ctx := model.BiSpecGenerationTemplateGlobalContext{
				Spec: t.spec,
			}
			if err := t.runTemplate(optional.NewStringEmpty(), tmpl, ctx, errpath); err != nil {
				return err
			}
		}
	}

	if util.IsDef(t.spec.Types) {
		if util.IsDef(t.spec.Generation, t.spec.Generation.Types, t.spec.Generation.Types.Templates) {
			for _, typ := range t.spec.Types {
				for i, tmpl := range t.spec.Generation.Types.Templates {
					errpath := []string{"generation", "types", fmt.Sprintf("templates[%v]", i)}

					ctx := model.BiSpecGenerationTemplateTypeContext{
						Type: typ,
						Spec: t.spec,
					}

					if err := t.runTemplate(typ.Name, tmpl, ctx, errpath); err != nil {
						return err
					}
				}
			}
		}
	}

	return nil
}

func (t *Generator) runTemplate(n optional.StringOptional, tmpl model.BiSpecGenerationTemplate, ctx interface{}, errpath []string) error {
	inputAbs, err := t.normalizeFromTemplateDir(tmpl.Input, errpath)

	if err != nil {
		return err
	}

	outputAbs, err := t.normalizeFromOutputDir(n, tmpl.Input, errpath)

	if err != nil {
		return err
	}

	if err := util.RenderFile(inputAbs.Get(), outputAbs.Get(), ctx); err != nil {
		return err
	}

	return nil
}

func (t *Generator) normalizeFromTemplateDir(p optional.StringOptional, errpath []string) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, t.generatorErr(model.CannotBeEmptyErr, errpath...)
	}

	np := filepath.Join(t.spec.Generation.TemplateDir.Get(), p.Get())
	abs, err := filepath.Abs(np)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	ok, err := util.PathExists(abs)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	if ok {
		return optional.NewStringValue(abs), nil
	}

	return p, t.generatorErr(model.NotFoundErr, errpath...)
}

func (t *Generator) normalizeFromOutputDir(n, p optional.StringOptional, errpath []string) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, t.generatorErr(model.CannotBeEmptyErr, errpath...)
	}

	abs := filepath.Join(t.spec.Generation.OutputDir.Get(), p.Get())
	abs, err := filepath.Abs(abs)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	abs = util.TrimTemplateExt(abs)

	if n.IsDefined() {
		var fn1 string
		var fn2 string

		fn1 = n.Get()
		fn1 = strings.ReplaceAll(fn1, "/", " ")
		fn1 = strutil.ToCamelCase(fn1)
		fn1 = fmt.Sprintf("%v%v", fn1, filepath.Ext(abs))
		fn2 = filepath.Base(abs)
		abs = strings.ReplaceAll(abs, fn2, fn1)
	}

	return optional.NewStringValue(abs), nil
}

// func (t *Generator) goPkg(inputPath string) string {
// 	var name string

// 	name = inputPath
// 	name = filepath.Dir(name)
// 	name = strings.TrimPrefix(name, t.templateRoot)
// 	name = path.Join(t.spec.X_Bi_Go.Module.Name.Get(), name)

// 	return name
// }

// func (t *Generator) runGlobal() error {
// 	if util.IsNil(t.spec.X_Bi_Go.Global, t.spec.X_Bi_Go.Global.Templates) {
// 		return nil
// 	}

// 	x_bi := t.spec.X_Bi_Go

// 	for _, tmpl := range x_bi.Global.Templates {
// 		var ctx model.XBiGoGlobalContext
// 		ctx.Model = t.spec

// 		if abs, err := t.isTemplateExist(*tmpl); err != nil {
// 			return err
// 		} else {
// 			ctx.Input = abs
// 		}

// 		ctx.Package = t.goPkg(ctx.Input)

// 		if abs, err := t.outputPath(optional.NewStringEmpty(), ctx.Input); err != nil {
// 			return err
// 		} else {
// 			ctx.Output = abs
// 		}

// 		if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
// 			return err
// 		}
// 	}

// 	return nil
// }

// func (t *Generator) runComponentsSchemas() error {
// 	if util.IsNil(t.spec.X_Bi_Go.Components, t.spec.X_Bi_Go.Components.Schemas, t.spec.X_Bi_Go.Components.Schemas.Templates) {
// 		return nil
// 	}

// 	if util.IsNil(t.spec.Components, t.spec.Components.Schemas) {
// 		return nil
// 	}

// 	for sn, sv := range t.spec.Components.Schemas {
// 		for _, tmpl := range t.spec.X_Bi_Go.Components.Schemas.Templates {
// 			var ctx model.XBiGoSchemaContext
// 			ctx.Model = t.spec
// 			ctx.Name = sn
// 			ctx.Schema = sv

// 			if abs, err := t.isTemplateExist(*tmpl); err != nil {
// 				return err
// 			} else {
// 				ctx.Input = abs
// 			}

// 			ctx.Package = t.goPkg(ctx.Input)

// 			if abs, err := t.outputPath(optional.NewStringValue(sn), ctx.Input); err != nil {
// 				return err
// 			} else {
// 				ctx.Output = abs
// 			}

// 			if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }

// func (t *Generator) runPaths() error {
// 	if util.IsNil(t.spec.X_Bi_Go.Paths, t.spec.X_Bi_Go.Paths.Templates) {
// 		return nil
// 	}

// 	if util.IsNil(t.spec.Paths) {
// 		return nil
// 	}

// 	for pn, pv := range t.spec.Paths {
// 		for _, tmpl := range t.spec.X_Bi_Go.Paths.Templates {
// 			var ctx model.XBiGoPathItemContext
// 			ctx.Model = t.spec
// 			ctx.Name = pn
// 			ctx.PathItem = pv

// 			if abs, err := t.isTemplateExist(*tmpl); err != nil {
// 				return err
// 			} else {
// 				ctx.Input = abs
// 			}

// 			ctx.Package = t.goPkg(ctx.Input)

// 			if abs, err := t.outputPath(optional.NewStringValue(pn), ctx.Input); err != nil {
// 				return err
// 			} else {
// 				ctx.Output = abs
// 			}

// 			if err := util.RenderFile(ctx.Input, ctx.Output, ctx); err != nil {
// 				return err
// 			}
// 		}
// 	}

// 	return nil
// }
