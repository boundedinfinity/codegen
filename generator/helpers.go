package generator

import (
	"boundedinfinity/codegen/model"
	"text/template"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/langer"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

var (
	ErrGeneratorQNameEmpty       = errorer.New("q-name empty")
	ErrGeneratorLangNotSupported = errorer.New("language not supported")
	ErrGeneratorPackageEmpty     = errorer.New("package empty")
)

func (t *Generator) getHelpers(lang string) template.FuncMap {
	helpers := map[string]template.FuncMap{
		"go": template.FuncMap{
			"typeName":    t.typeName,
			"typePackage": t.typePackage,
		},
	}

	return helpers[lang]
}

func (t *Generator) typeName(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	if typ.QName().Empty() {
		return result, ErrGeneratorQNameEmpty
	}

	result = typ.QName().Get()
	result = pather.Paths.Base(result)
	result, err = langer.Go.WithCaserConversion(t.caserConversion).Identifier(result)

	return result, err
}

func (t *Generator) typePackage(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	if typ.QName().Empty() {
		return result, ErrGeneratorPackageEmpty
	}

	result = typ.Common().QName().Get()
	result = pather.Paths.Dir(result)
	result = pather.Paths.Base(result)

	return result, err
}
