package template_manager

import (
	cp "boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/template_delimiter"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"text/template"

	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *TemplateManager) Register(root o.Option[string], file cp.CodeGenProjectTemplateFile) error {
	if file.Path.Empty() {
		return nil
	}

	tc := TemplateContext{
		Root:      root.OrElse(filepath.Dir(file.Path.Get())),
		Source:    file.Path.Get(),
		ModelType: util.GetCanonicalType(file.Path.Get()),
	}

	if try := util.GetTemplateType(tc.Source); try.Failure() {
		return try.Error
	} else {
		tc.TemplateMimeType = try.Result
	}

	if try := util.GetOutputType(tc.Source); try.Failure() {
		return try.Error
	} else {
		tc.OutputMimeType = try.Result
	}

	if tt, err := template_type.FromUrl(tc.Source); err != nil {
		return err
	} else {
		tc.TemplateType = tt
	}

	if bs, err := ioutil.ReadFile(tc.Source); err != nil {
		return err
	} else {
		l, r := template_delimiter.Get(t.projectManager.Merged.Info.Delimiter.Get())

		if t.projectManager.Merged.Info.TemplateDump.Defined() && t.projectManager.Merged.Info.TemplateDump.Get() {
			tmp := string(bs)
			tmp += fmt.Sprintf("\n\n%v DUMP . %v", l, r)
			bs = []byte(tmp)
		}

		if tmpl, err := template.New("").Funcs(t.funcs).Delims(l, r).Parse(string(bs)); err != nil {
			return err
		} else {
			tc.Template = tmpl
		}
	}

	t.pathMap[file.Path.Get()] = tc
	t.AppendTemplateContext(tc)

	return nil
}
