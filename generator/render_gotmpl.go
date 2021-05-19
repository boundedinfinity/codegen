package generator

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"bytes"
	"fmt"
	"text/template"
)

func (t *Generator) schema2Primtive(s string) string {
	p, _ := t.spec.Info.Primitives[s]
	return p
}

func (t *Generator) langType(m *model.OutputModel) string {
	typ := "<unkown type>"

	switch m.Type {
	case model.SchemaType_String, model.SchemaType_Int, model.SchemaType_Long, model.SchemaType_Float, model.SchemaType_Double, model.SchemaType_Boolean:
		typ = t.schema2Primtive(m.Type.String())
	case model.SchemaType_Ref:
		if ref, ok := t.spec.ModelMap[m.Ref]; ok {
			if util.SameNamespace(*m, *ref) {
				typ = m.Ref
			} else {
				n := util.UcFirst(util.NsBase(m.Ref))
				r := util.NsBase(util.NsDir(m.Ref))
				typ = fmt.Sprintf("%v.%v", r, n)
			}
		}
	}

	return typ
}

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {

	fnm := template.FuncMap{
		"ns_base":           util.NsBase,
		"ns_dir":            util.NsDir,
		"uc_first":          util.UcFirst,
		"uc":                util.Uc,
		"lc_first":          util.LcFirst,
		"lc":                util.Lc,
		"camel_case":        util.CamelCase,
		"path_2_camel_case": path2CamelCase,
		"is_int":            util.IsSchemaInt,
		"is_string":         util.IsSchemaString,
		"primitive":         t.schema2Primtive,
		"lang_type":         t.langType,
	}

	tmpl, err := template.New("template").Funcs(fnm).Parse(s)

	if err != nil {
		return "", err
	}

	var o bytes.Buffer

	if err := tmpl.Execute(&o, d); err != nil {
		return "", err
	}

	return o.String(), nil
}
