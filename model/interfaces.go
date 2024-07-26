package model

import "github.com/boundedinfinity/go-commoner/functional/optioner"

type CodeGenType interface {
	GetType() string
	GetQName() optioner.Option[string]
	GetName() optioner.Option[string]
	GetPackage() optioner.Option[string]
	Validate() error
	HasValidation() bool
}

type ArrayBuilder interface {
	Build() *CodeGenArray
	QName(string) ArrayBuilder
	Name(string) ArrayBuilder
	Package(string) ArrayBuilder
	Description(string) ArrayBuilder
	Required(bool) ArrayBuilder
	Min(int) ArrayBuilder
	Max(int) ArrayBuilder
	Items(CodeGenType) ArrayBuilder
}

type BooleanBuilder interface {
	Build() *CodeGenBoolean
	QName(string) BooleanBuilder
	Name(string) BooleanBuilder
	Package(string) BooleanBuilder
	Description(string) BooleanBuilder
	Required(bool) BooleanBuilder
}

type EnumBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenEnum
	QName(string) EnumBuilder
	Name(string) EnumBuilder
	Package(string) EnumBuilder
	Description(string) EnumBuilder
	Required(bool) EnumBuilder
	Values(...CodeGenEnumItem) EnumBuilder
}

type FloatBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenFloat
	QName(string) FloatBuilder
	Name(string) FloatBuilder
	Package(string) FloatBuilder
	Description(string) FloatBuilder
	Required(bool) FloatBuilder
	Ranges(...NumberRange[float64]) FloatBuilder
	MultipleOf(float64) FloatBuilder
	Negative() FloatBuilder
	Positive() FloatBuilder
	OneOf(...float64) FloatBuilder
	NoneOf(...float64) FloatBuilder
	Tolerance(float64) FloatBuilder
	Precision(int) FloatBuilder
}

type IntBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenInteger
	QName(string) IntBuilder
	Name(string) IntBuilder
	Package(string) IntBuilder
	Description(string) IntBuilder
	Required(bool) IntBuilder
	Ranges(...NumberRange[int]) IntBuilder
	MultipleOf(int) IntBuilder
	Negative() IntBuilder
	Positive() IntBuilder
	OneOf(...int) IntBuilder
	NoneOf(...int) IntBuilder
}

type ObjectBuilder interface {
	Build() *CodeGenObject
	QName(string) ObjectBuilder
	Name(string) ObjectBuilder
	Package(string) ObjectBuilder
	Description(string) ObjectBuilder
	Required(bool) ObjectBuilder
	Properties(...CodeGenType) ObjectBuilder
}

type RefBuilder interface {
	Build() CodeGenRef
	QName(string) RefBuilder
	Ref(string) RefBuilder
	Name(string) RefBuilder
	Package(string) RefBuilder
	Description(string) RefBuilder
	Required(bool) RefBuilder
}

type StringBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenString
	QName(string) StringBuilder
	Name(string) StringBuilder
	Package(string) StringBuilder
	Description(string) StringBuilder
	Required(bool) StringBuilder
	Min(int) StringBuilder
	Max(int) StringBuilder
	Regex(string) StringBuilder
	Abnf(string) StringBuilder
	Includes(...string) StringBuilder
	Excludes(...string) StringBuilder
	OneOf(...string) StringBuilder
	NoneOf(...string) StringBuilder
}
