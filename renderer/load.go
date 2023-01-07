package renderer

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/template_delimiter"
	"fmt"
	"io/ioutil"
	"text/template"
)

func (t *Renderer) Load(meta *ct.TemplateMeta) error {
	merged := t.projectManager.Merged
	bs, err := ioutil.ReadFile(meta.SourcePath.Get())

	if err != nil {
		return err
	}

	l, r := template_delimiter.Get(merged.Info.Delimiter.Get())

	if merged.Info.TemplateDump.Defined() && merged.Info.TemplateDump.Get() {
		tmp := string(bs)
		tmp += fmt.Sprintf("\n\n%v DUMP . %v", l, r)
		bs = []byte(tmp)
	}

	if tmpl, err := template.New("").Funcs(t.funcs).Delims(l, r).Parse(string(bs)); err != nil {
		return err
	} else {
		meta.Template = tmpl
	}

	return nil
}
