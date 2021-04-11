package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"errors"
	"fmt"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/optional"
)

type Generator struct {
	path   string
	schema model.OpenApiV310
}

func New(input string) *Generator {
	return &Generator{
		path: input,
	}
}

func (t *Generator) Generate() error {
	if err := t.loadSchema(); err != nil {
		return err
	}

	if err := t.generateGo(); err != nil {
		return err
	}

	return nil
}

func (t *Generator) loadSchema() error {
	if t.path == "" {
		return t.generatorErr(model.CannotBeEmptyErr, "path")
	}

	if abs, ok := util.FileSearch(optional.NewStringValue(t.path)); !ok {
		return t.generatorSchemaErr(errors.New("not found"))
	} else {
		t.path = abs
	}

	if err := util.UnmarshalFromFile(t.path, &t.schema); err != nil {
		return t.generatorSchemaErr(err)
	}
	return nil
}

func (t *Generator) generatorErr(err error, ps ...string) error {
	p := strings.Join(ps, ".")
	return fmt.Errorf("generator.%v: %w", p, err)
}

func (t *Generator) generatorSchemaErr(err error, ps ...string) error {
	tmp := []string{fmt.Sprintf("schema[%v]", t.path)}
	tmp = append(tmp, ps...)
	return t.generatorErr(err, tmp...)
}

func (t *Generator) searchFromSchemaPath(p optional.StringOptional) (optional.StringOptional, error) {
	sd := filepath.Dir(t.path)

	if abs, ok := util.FileSearch(p, optional.NewStringValue(sd)); !ok {
		return p, model.CannotBeEmptyErr
	} else {
		return optional.NewStringValue(abs), nil
	}
}
