package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Loader) processTemplate1(namespace model.BiOutput_Namespace, input model.BiInput_Template, output *model.BiInput_Template) error {
	{
		t.reportStack.Push("input")

		if input.Path == "" {
			return t.CannotBeEmpty()
		}

		if filepath.IsAbs(input.Path) {
			if ok, err := util.PathExists(input.Path); err != nil {
				return err
			} else if !ok {
				return t.NotFound()
			} else {
				output.Path = input.Path
			}
		} else {
			relPath := filepath.Join(t.inputDir, input.Path)
			abs, err := filepath.Abs(relPath)

			if err != nil {
				return err
			}

			if ok, err := util.PathExists(abs); err != nil {
				return err
			} else if !ok {
				return t.NotFound()
			} else {
				output.Path = abs
			}
		}

		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("type")

		if input.Type == "" {
			return t.NotFound()
		}

		if !model.IsTemplateType(input.Type) {
			return t.MustBeOneOf(model.TemplateTypeStrings())
		}

		output.Type = input.Type
		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("header")
		output.Header = input.Header
		t.reportStack.Pop()
	}

	return nil
}

func (t *Loader) processTemplate2(namespace model.BiOutput_Namespace, name string, input model.BiInput_Template, output *model.BiOutput_Template) error {
	var tmplExt string
	var tmplName string
	var relPath string
	var abs string
	var fn string

	relPath = namespace.Namespace
	relPath = strings.TrimPrefix(relPath, t.rootNamespace())
	relPath = strings.TrimPrefix(relPath, "/")

	tmplName = input.Path
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
	abs = path.Join(t.Output.Info.OutputDir, relPath, fn)

	output.Input = input.Path
	output.Output = abs
	output.Header = t.splitDescription(input.Header)
	return nil
}
