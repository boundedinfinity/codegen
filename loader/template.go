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
						outTmpls = append(outTmpls, tmpl)
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

	{
		t.reportStack.Push("input")

		if v.Path == "" {
			return model.CannotBeEmptyErr
		}

		if filepath.IsAbs(v.Path) {
			ok, err := util.PathExists(v.Path)

			if err != nil {
				return err
			}

			if !ok {
				return model.NotFoundErr
			}
		} else {
			relPath := filepath.Join(t.inputDir, v.Path)
			abs, err := filepath.Abs(relPath)

			if err != nil {
				return err
			}

			ok, err := util.PathExists(abs)

			if err != nil {
				return err
			}

			if !ok {
				return model.NotFoundErr
			} else {
				v.Path = abs
			}
		}

		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("type")

		if v.Type == "" {
			return model.CannotBeEmptyErr
		}

		if !model.IsTemplateType(v.Type) {
			return model.InvalidateType(model.TemplateTypeStrings())
		}

		t.reportStack.Pop()
	}

	ns := t.currentNamespace()

	if _, ok := t.templateMap[ns]; !ok {
		t.templateMap[ns] = make([]model.BiInput_Template, 0)
	}

	t.templateMap[ns] = append(t.templateMap[ns], v)

	t.reportStack.Pop()
	return nil
}

func (t *Loader) processTemplate2(ns, name string, input model.BiInput_Template) (model.BiOutput_Template, error) {
	output := model.BiOutput_Template{
		Input: input.Path,
	}

	var ext string
	var abs string
	var fn string

	ext = output.Input
	ext = filepath.Base(ext)
	ext = util.TrimTemplateExt(ext)
	ext = filepath.Ext(ext)
	ext = strings.TrimPrefix(ext, ".")

	if name == "" {
		var tmplNameOnly string
		tmplNameOnly = output.Input
		tmplNameOnly = filepath.Base(tmplNameOnly)
		tmplNameOnly = util.TrimTemplateExt(tmplNameOnly)
		fn = tmplNameOnly
	} else {
		fn = name
		fn = fmt.Sprintf("%v.%v", fn, ext)
	}

	abs = t.Output.Info.OutputDir
	abs = path.Join(abs, t.relativeNamespace(ns))
	abs = path.Join(abs, fn)

	output.Output = abs
	return output, nil
}
