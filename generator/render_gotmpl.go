package generator

import (
	"bytes"
	"text/template"
)

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"base_path":         basePath,
		"uc_first":          ucFirst,
		"uc":                uc,
		"lc_first":          lcFirst,
		"lc":                lc,
		"camel_case":        camelCase,
		"path_2_camel_case": path2CamelCase,
		"line_prefix":       linePrefix,
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
