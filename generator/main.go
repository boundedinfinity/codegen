package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"encoding/json"
	"fmt"
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
			if err := t.generateGoFile(rc, tmpl); err != nil {
				return err
			}
		}
	}

	if rc.Model.Components != nil {
		if rc.Model.Components.Schemas != nil {
			if rc.Model.Components.X_Bi_Go_Schemas != nil {
				if rc.Model.Components.X_Bi_Go_Schemas.Templates != nil {
					for _, tmpl := range rc.Model.Components.X_Bi_Go_Schemas.Templates {
						if err := t.generateGoFile(rc, tmpl); err != nil {
							return err
						}
					}
				}
			}
		}
	}

	return nil
}

func (t *Generator) generateGoFile(rc model.RunContext, tmpl model.XBiGoTemplate) error {
	// ctx := model.TemplateContext{
	// 	Model:    rc.Model,
	// 	Template: tmpl,
	// }

	ctxJson := `{
		"package": "model",
		"name": "tag",
		"schema": {
			"type": "object",
			"properties": {
				"id": { "type": "string" },
				"text": { "type": "string" },
				"description": { "type": "string" }
			}
		}
	}
	`

	var ctx model.GoLang

	if err := json.Unmarshal([]byte(ctxJson), &ctx); err != nil {
		return err
	}

	input := "/Users/bbabb200/dev/github.com/codegen-templates/go/server/echo/go/model.gotmpl"
	output := "/tmp/generator/go/server/echo/go/model/tag.gen.go"

	if err := util.RenderFile(input, output, ctx); err != nil {
		return err
	}
	// if err := util.RenderFile(*tmpl.Input, *tmpl.Output, ctx); err != nil {
	// 	return err
	// }

	return nil
}
