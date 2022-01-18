package util

import (
	"boundedinfinity/codegen/lang_ext"
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
	"strings"
	"unicode"

	"github.com/boundedinfinity/jsonschema"
	"github.com/boundedinfinity/optioner"
)

func GetTemplateSourcePath(lang lang_ext.LanguageExt, schema jsonschema.JsonSchmea) (string, error) {
	if schema.Type.IsEmpty() {
		return "", model.ErrMissingName
	}

	typ := schema.Type.Get()
	f := fmt.Sprintf("%v.gotmpl", typ)
	p := path.Join("templates", lang.String(), f)

	return p, nil
}

func GetPackageName(name optioner.StringOption, schema jsonschema.JsonSchmea, rs map[string]string) (string, error) {
	var n string

	if name.IsDefined() && len(name.Get()) > 1 {
		n = name.Get()
	} else if schema.Id.IsDefined() && len(schema.Id.Get()) > 1 {
		n = schema.Id.Get()
	} else {
		return "", model.ErrMissingName
	}

	n = strings.TrimSpace(n)
	n = ReplaceMap(n, rs)
	n = path.Dir(n)

	return n, nil
}

func GetTypeName(name optioner.StringOption, schema jsonschema.JsonSchmea, rs map[string]string) (string, error) {
	var n string

	if name.IsDefined() && len(name.Get()) > 1 {
		n = name.Get()
	} else if schema.Id.IsDefined() && len(schema.Id.Get()) > 1 {
		n = schema.Id.Get()
	} else {
		return "", model.ErrMissingName
	}

	n = path.Base(n)
	n = strings.TrimSpace(n)
	n = strings.ToLower(n)
	n = strings.Map(sanitize, n)

	return n, nil
}

func sanitize(r rune) rune {
	if r == '_' || unicode.IsLetter(r) || unicode.IsNumber(r) {
		return r
	} else {
		return '_'
	}
}
