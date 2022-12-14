package template_manager

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/template_delimiter"
	"boundedinfinity/codegen/template_type"
	"boundedinfinity/codegen/util"
	"fmt"
	"io/ioutil"
	"text/template"
)

func (t *TemplateManager) Register(templates model.CodeGenSchemaTemplates) error {
	for _, file := range templates.Files {
		if err := t.registerFileFile(file); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateManager) registerFileFile(file model.CodeGenSchemaTemplateFile) error {
	if file.Path.Empty() {
		return nil
	}

	cached := t.cacher.Find(file.Path.Get())

	if cached.Empty() {
		return fmt.Errorf("cache not found for %v", file.Path.Get())
	}

	if t.pathMap.Has(cached.Get().DestPath) {
		return model.ErrCodeGenTemplateFilePathDuplicatev(file.Path.Get())
	}

	tc := TemplateContext{
		Path:      cached.Get().DestPath,
		ModelType: util.GetCanonicalType(cached.Get().DestPath),
	}

	if try := util.GetTemplateType(tc.Path); try.Failure() {
		return try.Error
	} else {
		tc.TemplateMimeType = try.Result
	}

	if try := util.GetOutputType(tc.Path); try.Failure() {
		return try.Error
	} else {
		tc.OutputMimeType = try.Result
	}

	if tt, err := template_type.FromUrl(tc.Path); err != nil {
		return err
	} else {
		tc.TemplateType = tt
	}

	if bs, err := ioutil.ReadFile(tc.Path); err != nil {
		return err
	} else {
		if t.codeGenSchema.Info.TemplateDump.Defined() && t.codeGenSchema.Info.TemplateDump.Get() {
			l, r := template_delimiter.Get(t.codeGenSchema.Info.Delimiter.Get())
			tmp := string(bs)
			tmp += fmt.Sprintf("\n\n%v DUMP . %v", l, r)
			bs = []byte(tmp)
		}

		l, r := template_delimiter.Get(t.codeGenSchema.Info.Delimiter.Get())

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
