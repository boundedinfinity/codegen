package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Loader) getTemplates(ns string, typ model.TemplateType) ([]model.BiInput_Template, error) {
	var outTmpls []model.BiInput_Template
	cns := ns

	for cns != "." {
		if tmpls, ok := t.templateMap[cns]; ok {
			if tmpls != nil {
				for _, tmpl := range tmpls {
					if tmpl.Type == string(typ) {
						outTmpls = append(outTmpls, tmpls...)
					}
				}
			}
		}

		cns = path.Dir(cns)
	}

	return outTmpls, nil
}

func (t *Loader) processTemplate1(si int, v model.BiInput_Template) error {
	t.reportStack.Push("template[%v]", si)

	if v.Path == "" {
		return nil
	}

	ns := t.currentNamespace()

	if _, ok := t.templateMap[ns]; !ok {
		t.templateMap[ns] = make([]model.BiInput_Template, 0)
	}

	var found bool

	if !found {
		t.templateMap[ns] = append(t.templateMap[ns], v)
	}

	t.reportStack.Pop()
	return nil
}

func (t *Loader) processTemplate2(ns, name string, input model.BiInput_Template) (model.BiOutput_Template, error) {
	var output model.BiOutput_Template

	if input.Path == "" {
		return output, fmt.Errorf("template.input cannot be empty")
	}

	if filepath.IsAbs(input.Path) {
		output.Input = input.Path
	}

	ok, err := util.PathExists(output.Input)

	if err != nil {
		return output, err
	}

	if !ok {
		relPath := filepath.Join(t.inputDir, input.Path)
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
