package generator

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

import (
	"boundedinfinity/codegen/model"
	"errors"
	"strings"
	"text/template"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

var (
	ErrGeneratorIDEmpty             = errorer.New("id empty")
	ErrGeneratorLangNotSupported    = errorer.New("language not supported")
	ErrGeneratorPackageEmpty        = errorer.New("package empty")
	ErrGeneratorLangTypeUnsupported = errorer.New("unsupported lang")
)

func getHelpers(lang string) template.FuncMap {
	helpers := map[string]template.FuncMap{
		"go": {
			"sjoin":      strings.Join,
			"lowerFirst": stringer.LowercaseFirst[string],
			"dict":       dict,
		},
	}

	return helpers[lang]
}

func (this *Generator) resolve(typ model.CodeGenSchema) model.CodeGenSchema {
	var resolved model.CodeGenSchema

	switch x := typ.(type) {
	case *model.CodeGenArray:
		resolved = this.resolve(x.Items.Get())
	case *model.CodeGenRef:

	default:
		resolved = typ
	}
	return resolved
}

func dict(values ...interface{}) (map[string]interface{}, error) {
	if len(values)%2 != 0 {
		return nil, errors.New("invalid dict call")
	}
	dict := make(map[string]interface{}, len(values)/2)
	for i := 0; i < len(values); i += 2 {
		key, ok := values[i].(string)
		if !ok {
			return nil, errors.New("dict keys must be strings")
		}
		dict[key] = values[i+1]
	}
	return dict, nil
}
