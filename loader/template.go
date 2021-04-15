package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path/filepath"
	"strings"
)

func (t *Loader) processTemplate(name string, input model.BiSpecTemplate) (model.BiGenTemplate, error) {
	var output model.BiGenTemplate

	if input.Input == "" {
		return output, fmt.Errorf("template.input cannot be empty")
	}

	if filepath.IsAbs(input.Input) {
		output.Input = input.Input
	}

	ok, err := util.PathExists(output.Input)

	if err != nil {
		return output, err
	}

	if !ok {
		relPath := filepath.Join(t.specDir, input.Input)

		ok, err := util.PathExists(relPath)

		if err != nil {
			return output, err
		}

		if !ok {
			return output, fmt.Errorf("template.input not found")
		} else {
			output.Input = relPath
		}
	}

	if name == "" {
		return output, fmt.Errorf("template.ouput cannot be empty")
	}

	var ext string
	var abs string

	ext = output.Input
	ext = filepath.Base(ext)
	ext = util.TrimTemplateExt(ext)
	ext = filepath.Ext(ext)
	ext = strings.TrimPrefix(ext, ".")

	abs = name
	abs = fmt.Sprintf("%v.%v", abs, ext)
	abs = filepath.Join(t.Gen.Info.OutputDir, abs)

	output.Output = abs
	return output, nil
}
