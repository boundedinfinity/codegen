package model

import "errors"

type RunContext struct {
	Input string
	Model OpenApiV310
}

func (t RunContext) Validate() error {
	if t.Input == "" {
		return errors.New("modelPath is empty")
	}

	return nil
}

type TemplateContext struct {
	Model    OpenApiV310
	Template XBiGoTemplate
}
