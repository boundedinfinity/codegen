package generator

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

import (
	"boundedinfinity/codegen/model"
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

func (this *Generator) getHelpers(lang string) template.FuncMap {
	helpers := map[string]template.FuncMap{
		"go": {
			"sjoin":      strings.Join,
			"lowerFirst": stringer.LowercaseFirst[string],
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
