package renderer

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/template_delimiter"
	"fmt"
	"io/ioutil"
	"text/template"
)

func (t *Renderer) Load(lc *ct.TemplateLoaderContext) error {
	bs, err := ioutil.ReadFile(lc.FileInfo.Source)

	if err != nil {
		return err
	}

	l, r := template_delimiter.Get(t.projectManager.Merged.Info.Delimiter.Get())

	if t.projectManager.Merged.Info.TemplateDump.Defined() && t.projectManager.Merged.Info.TemplateDump.Get() {
		tmp := string(bs)
		tmp += fmt.Sprintf("\n\n%v DUMP . %v", l, r)
		bs = []byte(tmp)
	}

	if tmpl, err := template.New("").Funcs(t.funcs).Delims(l, r).Parse(string(bs)); err != nil {
		return err
	} else {
		lc.Template = tmpl
	}

	return nil
}
