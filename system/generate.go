package system

import (
	"boundedinfinity/codegen/model"
	"bytes"
	"text/template"
)

func (t *System) Generate() error {

	return nil
}

func (t *System) renderGoTemplate(ctx model.TemplateContext, input string, output *string) error {
	tmpl, err := template.New(ctx.Path).Funcs(functions).Parse(input)

	if err != nil {
		return err
	}

	var buff bytes.Buffer

	if err := tmpl.Execute(&buff, ctx); err != nil {
		return err
	}

	*output = buff.String()

	return nil
}
