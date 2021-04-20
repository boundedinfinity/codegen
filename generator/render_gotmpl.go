package generator

import (
	"boundedinfinity/codegen/util"
	"bytes"
	"text/template"
)

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"base_path":         util.BasePath,
		"uc_first":          util.UcFirst,
		"uc":                util.Uc,
		"lc_first":          util.LcFirst,
		"lc":                util.Lc,
		"camel_case":        util.CamelCase,
		"path_2_camel_case": path2CamelCase,
		"line_prefix":       linePrefix,
		"to_json":           to_json,
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
