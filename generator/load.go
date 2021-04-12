package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
)

func (t *Generator) load() error {
	if t.inputPath == "" {
		return t.generatorErr(model.CannotBeEmptyErr, "inputPath")
	}

	if abs, err := filepath.Abs(t.inputPath); err != nil {
		return t.generatorErr(err, "inputPath")
	} else {
		t.inputPath = abs
	}

	if err := util.UnmarshalFromFile(t.inputPath, &t.model); err != nil {
		return t.generatorSchemaErr(err)
	}

	if err := t.loadGo(); err != nil {
		return err
	}

	return nil
}

func (t *Generator) loadGo() error {
	if util.IsNil(t.model.X_Bi_Go) {
		return nil
	}

	x_bi := t.model.X_Bi_Go
	inputDir := filepath.Dir(t.inputPath)

	if x_bi.TemplateRoot.IsEmpty() {
		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "templateRoot")
	}

	ok, err := util.PathExists(x_bi.TemplateRoot.Get())

	if err != nil {
		return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
	}

	if ok {
		t.templateRoot = x_bi.TemplateRoot.Get()
	} else {
		p := filepath.Join(inputDir, x_bi.TemplateRoot.Get())
		if abs, err := filepath.Abs(p); err != nil {
			return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
		} else {
			ok, err := util.PathExists(abs)

			if err != nil {
				return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
			}

			if !ok {
				return t.generatorSchemaErr(fmt.Errorf("file not found"), "x-bi-go", "templateRoot")
			}

			t.templateRoot = abs
		}
	}

	if x_bi.GenRoot.IsDefined() {
		if filepath.IsAbs(x_bi.GenRoot.Get()) {
			t.genRoot = x_bi.GenRoot.Get()
		} else {
			if abs, err := filepath.Abs(x_bi.GenRoot.Get()); err != nil {
				return t.generatorErr(err, "x-bi-go", "genRoot")
			} else {
				t.genRoot = abs
			}
		}
	} else {
		if abs, err := filepath.Abs(".x-bi-gen"); err != nil {
			return t.generatorErr(err, "x-bi-go", "genRoot")
		} else {
			t.genRoot = abs
		}
	}

	return nil
}
