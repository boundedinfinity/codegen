package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Loader) processTemplate(ns, name string, input model.BiInput_Template_Info) (model.BiOutput_Template, error) {
	var output model.BiOutput_Template

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
		relPath := filepath.Join(t.inputDir, input.Input)
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
	var fn string

	ext = output.Input
	ext = filepath.Base(ext)
	ext = util.TrimTemplateExt(ext)
	ext = filepath.Ext(ext)
	ext = strings.TrimPrefix(ext, ".")

	fn = name
	fn = fmt.Sprintf("%v.%v", fn, ext)

	abs = t.Output.Info.OutputDir
	abs = path.Join(abs, t.relativeNamespace(ns))
	abs = path.Join(abs, fn)

	output.Output = abs
	return output, nil
}
