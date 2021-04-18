package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"errors"
	"fmt"
	"path/filepath"
)

func (t *Loader) processInput_Info() error {
	t.reportStack.Push("info")

	input := &t.input.Info
	output := &t.Output.Info

	output.DumpContext = input.DumpContext

	{
		t.reportStack.Push("dumpContext")
		t.report("%v", output.DumpContext)
		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("inputDir")

		if input.InputDir == "" {
			return model.CannotBeEmptyErr
		}

		if abs, err := filepath.Abs(input.InputDir); err != nil {
			return err
		} else {
			ok, err := util.PathExists(abs)

			if err != nil {
				return err
			}

			if !ok {
				relPath := filepath.Join(t.inputDir, input.InputDir)

				if abs, err := filepath.Abs(relPath); err != nil {
					return err
				} else {
					ok, err := util.PathExists(abs)

					if err != nil {
						return fmt.Errorf("info.templateDir error: %w", err)
					}

					if !ok {
						return errors.New("info.templateDir not found")
					} else {
						output.InputDir = abs
					}
				}
			} else {
				output.InputDir = abs
			}
		}

		t.report(output.InputDir)
		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("outputDir")

		if input.OutputDir == "" {
			return model.CannotBeEmptyErr
		}

		if filepath.IsAbs(input.OutputDir) {
			output.OutputDir = input.OutputDir
		} else {
			input.OutputDir = filepath.Join(t.inputDir, input.OutputDir)
		}

		t.report(output.OutputDir)
		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("typeMap")

		if input.TypeMap != nil {
			for k, v := range input.TypeMap {
				t.addMappedType(k, TypeInfo{
					BaseName:   v,
					ImportName: v,
					QName:      v,
					Namespace:  model.NAMESPACE_BUILTIN,
					BuiltIn:    true,
				})
			}
		}

		t.report("read %v type mappings", len(input.TypeMap))

		t.reportStack.Pop()
	}

	t.reportStack.Pop()
	return nil
}
