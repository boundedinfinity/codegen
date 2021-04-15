package generator

import (
	"bytes"
	"path"
	"strings"
	"text/template"
)

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"basePath":     basePath,
		"ucFirst":      ucFirst,
		"uc":           uc,
		"lcFirst":      lcFirst,
		"lc":           lc,
		"peq":          peq,
		"lang_type":    t.typeGo,
		"lang_rel_pkg": t.goType2RelativePkg,
		"lang_abs_pkg": t.goType2AbsolutePkg,
		"operationId":  operationId,
		"filename":     filename,
		"camelCase":    camelCase,
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

func (t *Generator) typeGo(v string) string {
	var x string

	if lang, ok := t.mapper.Language["go"]; ok {
		if strings.HasSuffix(v, "[]") {

		}

		if typ, ok := lang[v]; ok {
			x = typ
		} else {
			x = "<UNKNOWN_TYPE>"
		}
	} else {
		x = "<UNKNOWN_LANG>"
	}

	return x
}

func (t *Generator) goType2RelativePkg(v string) string {
	var x string

	x = v
	x = path.Dir(x)
	x = path.Base(x)

	return x
}

func (t *Generator) goType2AbsolutePkg(v string) string {
	var x string

	x = v
	x = path.Join(t.spec.Name, x)
	x = path.Dir(x)

	return x
}
