package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
)

func (t *Generator) isTemplateExist(tmpl model.X_Bi_Go_Template) (string, error) {
	if tmpl.Input.IsEmpty() {
		return "", model.CannotBeEmptyErr
	}

	rel := filepath.Join(t.templateRoot, tmpl.Input.Get())
	abs, err := filepath.Abs(rel)

	if err != nil {
		return "", err
	}

	ok, err := util.PathExists(abs)

	if err != nil {
		return "", err
	}

	if !ok {
		return "", fmt.Errorf("%v not found", tmpl.Input.Get())
	}

	return abs, nil
}
