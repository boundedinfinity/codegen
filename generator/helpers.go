package generator

import (
	"boundedinfinity/codegen/model"
	"text/template"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/idiomatic/langer"
)

var (
	ErrGeneratorLangTypeNameEmpty = errorer.New("type name empty")
	ErrGeneratorLangNotSupported  = errorer.New("lang not supported")
	ErrGeneratorPackageEmpty      = errorer.New("package empty")
)

func (t *Generator) getHelpers() template.FuncMap {
	return template.FuncMap{
		"typeName":    t.typeName,
		"typePackage": t.typePackage,
	}
}

func (t *Generator) typeName(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	if typ.TypeId().Empty() {
		return result, ErrGeneratorLangTypeNameEmpty
	}

	result = typ.TypeId().Get()

	switch t.lang {
	case "go":
		result, err = langer.Go.Identifier(result)
	default:
		return result, ErrGeneratorLangNotSupported.WithValue(t.lang)
	}

	return result, err
}

func (t *Generator) typePackage(typ model.CodeGenType) (string, error) {
	var result string
	var err error

	if typ.TypeId().Empty() {
		return result, ErrGeneratorPackageEmpty
	}

	result = typ.TypeId().Get()

	switch t.lang {
	case "go":
		result, err = langer.Go.Identifier(result)
	default:
		return result, ErrGeneratorLangNotSupported.WithValue(t.lang)
	}

	return result, err
}
