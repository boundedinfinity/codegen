package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
)

func (t *Loader) processInput_Info() error {
	t.reportStack.Push("info")
	defer t.reportStack.Pop()

	input := &t.input.Info
	output := &t.Output.Info

	output.DumpContext = input.DumpContext

	checkDumpContext := func() error {
		t.reportStack.Push("dumpContext")
		defer t.reportStack.Pop()

		t.report("%v", output.DumpContext)

		return nil
	}

	checkFilenameMarker := func() error {
		t.reportStack.Push("filenameMarker")
		defer t.reportStack.Pop()
		t.report("%v", input.FilenameMarker)
		return nil
	}

	if err := checkDumpContext(); err != nil {
		return err
	}

	if err := checkFilenameMarker(); err != nil {
		return err
	}

	if err := t.processInput_Info_inputDir(t.input.Info, &t.Output.Info); err != nil {
		return err
	}

	if err := t.processInput_Info_outputDir(t.input.Info, &t.Output.Info); err != nil {
		return err
	}

	if err := t.processInput_Info_TypeMap(t.input.Info, &t.Output.Info); err != nil {
		return err
	}

	return nil
}

func (t *Loader) processInput_Info_inputDir(input model.BiInput_Info, output *model.BiOutput_Info) error {
	t.reportStack.Push("inputDir")
	defer t.reportStack.Pop()

	if input.InputDir == "" {
		return t.CannotBeEmpty()
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
					return err
				}

				if !ok {
					return t.NotFound()
				} else {
					output.InputDir = abs
				}
			}
		} else {
			output.InputDir = abs
		}
	}

	t.report(output.InputDir)

	return nil
}

func (t *Loader) processInput_Info_outputDir(input model.BiInput_Info, output *model.BiOutput_Info) error {
	t.reportStack.Push("outputDir")
	defer t.reportStack.Pop()

	if input.OutputDir == "" {
		return t.CannotBeEmpty()
	}

	if filepath.IsAbs(input.OutputDir) {
		output.OutputDir = input.OutputDir
	} else {
		input.OutputDir = filepath.Join(t.inputDir, input.OutputDir)
	}

	t.report(output.OutputDir)

	return nil
}

func (t *Loader) processInput_Info_TypeMap(input model.BiInput_Info, output *model.BiOutput_Info) error {
	t.reportStack.Push("typeMap")
	defer t.reportStack.Pop()

	checkBuiltIn := func() error {
		t.reportStack.Push("builtIn")
		defer t.reportStack.Pop()

		if input.TypeMap.BuiltIn != nil {
			for k, v := range input.TypeMap.BuiltIn {
				if _, ok := t.builtInTypeMap[k]; !ok {
					t.builtInTypeMap[k] = v
					t.builtInTypeMap[fmt.Sprintf("%v[]", k)] = fmt.Sprintf("%v[]", v)
					t.report("mapping %v -> %v", util.SummerySuffix(k, model.SUMMERY_SIZE), v)
				} else {
					return t.DuplicateType(k)
				}
			}
		}

		return nil
	}

	checkCustom := func() error {
		t.reportStack.Push("custom")
		defer t.reportStack.Pop()

		if input.TypeMap.Custom != nil {
			for sk, v := range input.TypeMap.Custom {
				k := t.absoluteNamespace(sk)
				if _, ok := t.customTypeMap[k]; !ok {
					t.customTypeMap[k] = v
					t.customTypeMap[fmt.Sprintf("%v[]", k)] = fmt.Sprintf("%v[]", v)
					t.report("mapping %v -> %v", util.SummerySuffix(k, model.SUMMERY_SIZE), v)
				} else {
					return t.DuplicateType(k)
				}
			}
		}

		return nil
	}

	if err := checkBuiltIn(); err != nil {
		return err
	}

	if err := checkCustom(); err != nil {
		return err
	}

	return nil
}
