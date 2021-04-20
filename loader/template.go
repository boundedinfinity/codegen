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
			return t.CannotBeEmpty()
		}

		if filepath.IsAbs(v.Path) {
			ok, err := util.PathExists(v.Path)

			if err != nil {
				return err
			}

			if !ok {
				return t.NotFound()
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
				return t.NotFound()
			} else {
				v.Path = abs
			}
		}

		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("type")

		if v.Type == "" {
			return t.NotFound()
		}

		if !model.IsTemplateType(v.Type) {
			return t.MustBeOneOf(model.TemplateTypeStrings())
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

	var tmplExt string
	var tmplName string
	var abs string
	var fn string

	tmplName = output.Input
	tmplName = filepath.Base(tmplName)
	tmplName = util.TrimTemplateExt(tmplName)
	tmplExt = filepath.Ext(tmplName)
	tmplName = strings.TrimSuffix(tmplName, tmplExt)
	tmplExt = strings.TrimPrefix(tmplExt, ".")

	if name == "" {
		fn = tmplName
	} else {
		fn = name
	}

	if t.input.Info.FilenameMarker != "" {
		fn = fmt.Sprintf("%v.%v", fn, t.input.Info.FilenameMarker)
	}

	fn = fmt.Sprintf("%v.%v", fn, tmplExt)

	abs = t.Output.Info.OutputDir
	abs = path.Join(abs, t.relativeNamespace(ns))
	abs = path.Join(abs, fn)

	output.Output = abs
	return output, nil
}
