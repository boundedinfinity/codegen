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
	ErrGeneratorIDEmpty             = errorer.New("id empty")
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

	if typ.Common().Id.Empty() {
		return result, ErrGeneratorIDEmpty
	}

	result = typ.Common().Id.Get()
	result = pather.Paths.Base(result)
	result, err = langer.Go.WithCaserConversion(t.caserConversion).Identifier(result)

	return result, err
}

func (t *Generator) typePackage(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	if typ.Common().Id.Empty() {
		return result, ErrGeneratorPackageEmpty
	}

	result = typ.Common().Id.Get()
	result = pather.Paths.Dir(result)
	result = pather.Paths.Base(result)

	return result, err
}

func (t *Generator) langType(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	switch typ.GetType() {
	case model.CodeGenString{}.GetType():
		result = "string"
	case model.CodeGenInteger{}.GetType():
		result = "int"
	case model.CodeGenFloat{}.GetType():
		result = "float64"
	case model.CodeGenBoolean{}.GetType():
		result = "bool"
	default:
		return result, ErrGeneratorLangTypeUnsupported.WithValue(typ.GetType())
	}

	return result, err
}
