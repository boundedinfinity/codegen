package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"

	"github.com/boundedinfinity/optional"
)

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

func (t *Generator) Generate(rc model.RunContext) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("generator error: %w", err)
	}

	if err := t.generateGo(rc); err != nil {
		return wrapErr(err)
	}

	return nil
}

func (t *Generator) generateGo(rc model.RunContext) error {
	if rc.Model.X_Bi_Go == nil {
		return nil
	}

	if rc.Model.X_Bi_Go != nil && rc.Model.X_Bi_Go.Templates != nil {
		for _, tmpl := range rc.Model.X_Bi_Go.Templates {
			ctx := model.GoLang{
				Model:    rc.Model,
				Template: *tmpl,
			}
			if err := t.generateGoFile(*tmpl, ctx); err != nil {
				return err
			}
		}
	}

	if rc.Model.Components != nil {
		if rc.Model.Components.Schemas != nil {
			for sn, sv := range rc.Model.Components.Schemas {
				if sv.X_Bi_Go != nil {
					if sv.X_Bi_Go.Templates != nil {
						for _, tmpl := range sv.X_Bi_Go.Templates {
							ctx := model.GoLang{
								Model:    rc.Model,
								Template: *tmpl,
								Name:     optional.NewStringValue(sn),
								Schema:   &sv,
							}
							if err := t.generateGoFile(*tmpl, ctx); err != nil {
								return err
							}
						}
					}
				}
			}
		}
	}

	return nil
}

func (t *Generator) generateGoFile(tmpl model.XBiGoTemplate, ctx model.GoLang) error {

	if err := util.RenderFile(tmpl.Input.Get(), tmpl.Output.Get(), ctx); err != nil {
		return err
	}

	return nil
}
