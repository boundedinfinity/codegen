package generator

import (
	"bytes"
	"fmt"
	"path"
	"strings"
	"text/template"
)

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"base_path":    basePath,
		"uc_first":     ucFirst,
		"uc":           uc,
		"lc_first":     lcFirst,
		"lc":           lc,
		"custom_type":  t.typeCustom,
		"lang_type":    t.typeGo,
		"lang_rel_pkg": t.goType2RelativePkg,
		"lang_abs_pkg": t.goType2AbsolutePkg,
		"operationId":  operationId,
		"filename":     filename,
		"camel_case":   camelCase,
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

const (
	UNKNOWN_TYPE = "<UNKNOWN_TYPE>"
)

func (t *Generator) typeCustom(v string) bool {
	if lang, ok := t.mapper.Language["go"]; ok {
		for k := range lang {
			v = strings.TrimSuffix(v, "[]")
			if k == v {
				return false
			}
		}
	}

	return true
}

func (t *Generator) typeGo(v string) string {
	x := UNKNOWN_TYPE
	isArrary := false

	if strings.HasSuffix(v, "[]") {
		isArrary = true
		v = strings.TrimSuffix(v, "[]")
	}

	if lang, ok := t.mapper.Language["go"]; ok {

		if typ, ok := lang[v]; ok {
			x = typ
			if isArrary {
				x = fmt.Sprintf("[]%v", x)
			}
		}
	}

	if x == UNKNOWN_TYPE {
		if typ, ok := t.spec.Lookup[v]; ok {
			x = typ.Name
		}
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
