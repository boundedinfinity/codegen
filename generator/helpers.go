package generator

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"strings"
	"text/template"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/langer"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

var (
	ErrGeneratorQNameEmpty          = errorer.New("q-name empty")
	ErrGeneratorLangNotSupported    = errorer.New("language not supported")
	ErrGeneratorPackageEmpty        = errorer.New("package empty")
	ErrGeneratorLangTypeUnsupported = errorer.New("unsupported lang")
)

func (t *Generator) getHelpers(lang string) template.FuncMap {
	helpers := map[string]template.FuncMap{
		"go": {
			"optPtr":      t.optPtr,
			"typeName":    t.typeName,
			"typePackage": t.typePackage,
			"langType":    t.langType,
			"sjoin":       strings.Join,
		},
	}

	return helpers[lang]
}

func (t *Generator) optPtr(opt optioner.Option[int]) (string, error) {
	var result string
	var err error

	if opt.Defined() {
		result = fmt.Sprintf("%v", opt.Get())
	} else {
		result = "nil"
	}

	return result, err
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

func (t *Generator) langType(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	switch typ.BaseType() {
	case model.CodeGenString{}.BaseType():
		result = "string"
	case model.CodeGenInteger{}.BaseType():
		result = "int"
	case model.CodeGenFloat{}.BaseType():
		result = "float64"
	case model.CodeGenBoolean{}.BaseType():
		result = "bool"
	default:
		return result, ErrGeneratorLangTypeUnsupported.WithValue(typ.BaseType())
	}

	return result, err
}
