package loader

import (
	"boundedinfinity/codegen/util"
	"errors"
	"fmt"
	"path/filepath"
)

func (t *Loader) processSpec() error {
	t.Gen.Name = t.spec.Name

	if t.spec.Info.TemplateDir == "" {
		return errors.New("info.templateDir missing")
	}

	if abs, err := filepath.Abs(t.spec.Info.TemplateDir); err != nil {
		return fmt.Errorf("info.templateDir error: %w", err)
	} else {
		ok, err := util.PathExists(abs)

		if err != nil {
			return fmt.Errorf("info.templateDir error: %w", err)
		}

		if !ok {
			relPath := filepath.Join(t.specDir, t.spec.Info.TemplateDir)

			if abs, err := filepath.Abs(relPath); err != nil {
				return fmt.Errorf("info.templateDir error: %w", err)
			} else {
				ok, err := util.PathExists(abs)

				if err != nil {
					return fmt.Errorf("info.templateDir error: %w", err)
				}

				if !ok {
					return errors.New("info.templateDir not found")
				} else {
					t.Gen.Info.TemplateDir = abs
				}
			}
		} else {
			t.Gen.Info.TemplateDir = abs
		}
	}

	if t.spec.Info.OutputDir == "" {
		return errors.New("info.outputDir missing")
	}

	if filepath.IsAbs(t.spec.Info.OutputDir) {
		t.Gen.Info.OutputDir = t.spec.Info.OutputDir
	} else {
		t.Gen.Info.OutputDir = filepath.Join(t.specDir, t.spec.Info.OutputDir)
	}

	if t.spec.Info.TypeMap != nil {
		for k, v := range t.spec.Info.TypeMap {
			t.addMappedType(k, v)
		}
	}

	if t.spec.Models.Namespaces != nil {
		for _, ns := range t.spec.Models.Namespaces {
			if err := t.processNamespace1(ns); err != nil {
				return err
			}
		}

		for _, sNamespace := range t.spec.Models.Namespaces {
			gNamespace, err := t.processNamespace2(sNamespace, t.spec.Models.Templates)

			if err != nil {
				return err
			}

			t.Gen.Models.Namespaces = append(t.Gen.Models.Namespaces, gNamespace)
		}
	}

	if t.spec.Operations.Namespaces != nil {
		for _, ns := range t.spec.Operations.Namespaces {
			if err := t.processNamespace1(ns); err != nil {
				return err
			}
		}

		for _, sNamespace := range t.spec.Operations.Namespaces {
			gNamespace, err := t.processNamespace2(sNamespace, t.spec.Operations.Templates)

			if err != nil {
				return err
			}

			t.Gen.Operations.Namespaces = append(t.Gen.Operations.Namespaces, gNamespace)
		}
	}

	return nil
}
