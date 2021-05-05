package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path/filepath"
)

func (t *Loader) processInput_Info() error {
	t.reportStack.Push("info")
	defer t.reportStack.Pop()

	input := t.inputSpec.Info
	output := &t.OutputSpec.Info

	if err := t.processInput_Info_inputDir(input, output); err != nil {
		return err
	}

	if err := t.processInput_Info_outputDir(input, output); err != nil {
		return err
	}

	checkDumpContext := func() error {
		t.reportStack.Push("dumpContext")
		defer t.reportStack.Pop()

		output.DumpContext = input.DumpContext
		t.reportf(t.reportStack.S(), "%v", output.DumpContext)

		return nil
	}

	checkFilenameMarker := func() error {
		t.reportStack.Push("filenameMarker")
		defer t.reportStack.Pop()

		if input.FilenameMarker == "" {
			output.FilenameMarker = model.DEFAULT_FILENAME_MARKER
		} else {
			output.FilenameMarker = input.FilenameMarker
		}

		t.reportf(t.reportStack.S(), "%v", output.FilenameMarker)
		return nil
	}

	if err := checkDumpContext(); err != nil {
		return err
	}

	if err := checkFilenameMarker(); err != nil {
		return err
	}

	return nil
}

func (t *Loader) processInput_Info_inputDir(input model.InputInfo, output *model.OutputInfo) error {
	t.reportStack.Push("inputDir")
	defer t.reportStack.Pop()

	if input.InputDir == "" {
		t.OutputSpec.Info.InputDir = filepath.Dir(t.inputPath)
	} else if abs, err := filepath.Abs(input.InputDir); err != nil {
		return err
	} else {
		ok, err := util.PathExists(abs)

		if err != nil {
			return err
		}

		if !ok {
			relPath := filepath.Join(t.OutputSpec.Info.OutputDir, input.InputDir)

			if abs, err := filepath.Abs(relPath); err != nil {
				return err
			} else {
				ok, err := util.PathExists(abs)

				if err != nil {
					return err
				}

				if !ok {
					return t.ErrNotFound()
				} else {
					output.InputDir = abs
				}
			}
		} else {
			output.InputDir = abs
		}
	}

	t.reportf(t.reportStack.S(), output.InputDir)

	return nil
}

func (t *Loader) processInput_Info_outputDir(input model.InputInfo, output *model.OutputInfo) error {
	t.reportStack.Push("outputDir")
	defer t.reportStack.Pop()

	if input.OutputDir == "" {
		output.OutputDir = filepath.Join(filepath.Dir(t.inputPath), "codegen-output")
	} else if filepath.IsAbs(input.OutputDir) {
		output.OutputDir = input.OutputDir
	} else {
		input.OutputDir = filepath.Join(t.OutputSpec.Info.OutputDir, input.OutputDir)
	}

	t.reportf(t.reportStack.S(), output.OutputDir)

	return nil
}
