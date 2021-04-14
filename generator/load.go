package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"

	"github.com/boundedinfinity/optional"
)

func (t *Generator) load() error {
	if t.specPath == "" {
		return t.generatorErr(model.CannotBeEmptyErr, "spec")
	}

	if abs, err := filepath.Abs(t.specPath); err != nil {
		return t.generatorErr(err, "spec")
	} else {
		t.specPath = abs
	}

	if err := util.UnmarshalFromFile(t.specPath, &t.spec); err != nil {
		return t.generatorSchemaErr(err)
	}

	if err := t.normalizeTemplateDir(); err != nil {
		return err
	}

	if err := t.normalizeOutputDir(); err != nil {
		return err
	}

	return nil
}

func (t *Generator) normalizeFromSpecPath(p optional.StringOptional, errpath []string) (optional.StringOptional, error) {
	if p.IsEmpty() {
		return p, t.generatorErr(model.CannotBeEmptyErr, errpath...)
	}

	abs, err := filepath.Abs(p.Get())

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	ok, err := util.PathExists(abs)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	if ok {
		return optional.NewStringValue(abs), nil
	}

	specDir := filepath.Dir(t.specPath)
	newPath := filepath.Join(specDir, p.Get())
	abs, err = filepath.Abs(newPath)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	ok, err = util.PathExists(abs)

	if err != nil {
		return p, t.generatorErr(err, errpath...)
	}

	if ok {
		return optional.NewStringValue(abs), nil
	}

	return p, t.generatorErr(model.NotFoundErr, errpath...)
}

func (t *Generator) normalizeTemplateDir() error {
	errpath := []string{"generation", "templateDir"}
	abs, err := t.normalizeFromSpecPath(t.spec.Generation.TemplateDir, errpath)

	if err != nil {
		return err
	} else {
		t.spec.Generation.TemplateDir = abs
	}

	return nil
}

func (t *Generator) normalizeOutputDir() error {
	errpath := []string{"generation", "outputDir"}
	abs, err := t.normalizeFromSpecPath(t.spec.Generation.OutputDir, errpath)

	if err != nil {
		return err
	} else {
		t.spec.Generation.OutputDir = abs
	}

	return nil
}

// func (t *Generator) loadGo() error {
// 	if util.IsNil(t.spec.Generation) {
// 		return nil
// 	}

// 	x_bi := t.spec.X_Bi_Go
// 	inputDir := filepath.Dir(t.specPath)

// 	if x_bi.TemplateRoot.IsEmpty() {
// 		return t.generatorSchemaErr(model.CannotBeEmptyErr, "x-bi-go", "templateRoot")
// 	}

// 	ok, err := util.PathExists(x_bi.TemplateRoot.Get())

// 	if err != nil {
// 		return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
// 	}

// 	if ok {
// 		t.templateRoot = x_bi.TemplateRoot.Get()
// 	} else {
// 		p := filepath.Join(inputDir, x_bi.TemplateRoot.Get())
// 		if abs, err := filepath.Abs(p); err != nil {
// 			return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
// 		} else {
// 			ok, err := util.PathExists(abs)

// 			if err != nil {
// 				return t.generatorSchemaErr(err, "x-bi-go", "templateRoot")
// 			}

// 			if !ok {
// 				return t.generatorSchemaErr(fmt.Errorf("file not found"), "x-bi-go", "templateRoot")
// 			}

// 			t.templateRoot = abs
// 		}
// 	}

// 	if x_bi.GenRoot.IsDefined() {
// 		if filepath.IsAbs(x_bi.GenRoot.Get()) {
// 			t.genRoot = x_bi.GenRoot.Get()
// 		} else {
// 			if abs, err := filepath.Abs(x_bi.GenRoot.Get()); err != nil {
// 				return t.generatorErr(err, "x-bi-go", "genRoot")
// 			} else {
// 				t.genRoot = abs
// 			}
// 		}
// 	} else {
// 		if abs, err := filepath.Abs(".x-bi-gen"); err != nil {
// 			return t.generatorErr(err, "x-bi-go", "genRoot")
// 		} else {
// 			t.genRoot = abs
// 		}
// 	}

// 	return nil
// }
