package generator

import (
	"boundedinfinity/codegen/model"
	"fmt"
)

type Generator struct {
	rc model.RunContext
}

func New(rc model.RunContext) *Generator {
	return &Generator{
		rc: rc,
	}
}

func (t *Generator) Generate() error {
	wrapErr := func(err error) error {
		return fmt.Errorf("generator error: %w", err)
	}

	if err := t.generateGo(); err != nil {
		return wrapErr(err)
	}

	return nil
}
