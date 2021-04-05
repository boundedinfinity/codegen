package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

type Generator struct {
}

func New() *Generator {
	return &Generator{}
}

func (t *Generator) Generate(rctx model.RunContext) error {
	wrapErr := func(err error) error {
		return fmt.Errorf("generator error: %w", err)
	}

	if err := t.generateGo(rctx); err != nil {
		return wrapErr(err)
	}

	return nil
}

func (t *Generator) generateGo(rctx model.RunContext) error {
	if err := t.generateGoFile(rctx, rctx.Project.Global); err != nil {
		return err
	}

	if err := t.generateGoFile(rctx, rctx.Project.Schemas); err != nil {
		return err
	}

	return nil
}

func (t *Generator) generateGoFile(rctx model.RunContext, tmpls []model.XBiGoTemplate) error {
	if tmpls == nil {
		return nil
	}

	for i, tmpl := range tmpls {
		ctx := model.GoFileContext{
			Project: rctx.Project,
			Config:  tmpl,
			Model:   rctx.Model,
		}

		if err := util.RenderFile(*tmpl.Input, *tmpl.Output, ctx); err != nil {
			return fmt.Errorf("project[%v].global[%v].template[%v] render error: %w", rctx.ProjectPath, i, *tmpl.Input, err)
		}
	}

	return nil
}
