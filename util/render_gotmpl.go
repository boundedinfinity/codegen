package util

import (
	"bytes"
	"text/template"
)

func renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"basePath":    basePath,
		"ucFirst":     ucFirst,
		"uc":          uc,
		"lcFirst":     lcFirst,
		"lc":          lc,
		"peq":         peq,
		"type_go":     typeGo,
		"operationId": operationId,
		"filename":    filename,
	}

	t, err := template.New("template").Funcs(fnm).Parse(s)

	if err != nil {
		return "", err
	}

	var o bytes.Buffer

	if err := t.Execute(&o, d); err != nil {
		return "", err
	}

	return o.String(), nil
}
