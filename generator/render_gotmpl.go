package generator

import (
	"boundedinfinity/codegen/util"
	"bytes"
	"strings"
	"text/template"
)

func (t *Generator) renderGoTemplate(s string, d interface{}) (string, error) {
	fnm := template.FuncMap{
		"path_base":         util.PathBase,
		"path_dir":          util.PathDir,
		"path_base_dir":     util.PathDirBase,
		"uc_first":          util.UcFirst,
		"uc":                util.Uc,
		"lc_first":          util.LcFirst,
		"lc":                util.Lc,
		"camel_case":        util.CamelCase,
		"path_2_camel_case": path2CamelCase,
		"is_int":            util.IsSchemaInt,
		"is_string":         util.IsSchemaString,
		"primitive":         t.schema2Primtive,
		"join":              strings.Join,
		"to_json":           util.ToJson,
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
