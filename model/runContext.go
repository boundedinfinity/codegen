package model

import "errors"

type RunContext struct {
	ModelPath string
	Model     OpenApiV310
}

func (t RunContext) Validate() error {
	if t.ModelPath == "" {
		return errors.New("modelPath is empty")
	}

	return nil
}
