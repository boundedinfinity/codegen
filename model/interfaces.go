package model

type CodeGenType interface {
	Common() *CodeGenCommon
	GetType() string
	Validate() error
	HasValidation() bool
}

type ArrayBuilder interface {
	Build() *CodeGenArray
	Id(string) ArrayBuilder
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
	Id(string) BooleanBuilder
	Name(string) BooleanBuilder
	Package(string) BooleanBuilder
	Description(string) BooleanBuilder
	Required(bool) BooleanBuilder
}

type EnumBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenEnum
	Id(string) EnumBuilder
	Name(string) EnumBuilder
	Package(string) EnumBuilder
	Description(string) EnumBuilder
	Required(bool) EnumBuilder
	Values(...CodeGenEnumItem) EnumBuilder
}

type FloatBuilder interface {
	Ref() RefBuilder
	Build() *CodeGenFloat
	Id(string) FloatBuilder
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
	Id(string) IntBuilder
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
	Id(string) ObjectBuilder
	Name(string) ObjectBuilder
	Package(string) ObjectBuilder
	Description(string) ObjectBuilder
	Required(bool) ObjectBuilder
	Properties(...CodeGenType) ObjectBuilder
}

type RefBuilder interface {
	Build() *CodeGenRef
	Resolved(CodeGenType) RefBuilder
	Ref(string) RefBuilder
	Name(string) RefBuilder
	Package(string) RefBuilder
	Description(string) RefBuilder
	Required(bool) RefBuilder
}

type StringBuilder interface {
	Build() *CodeGenString
	Ref() RefBuilder
	Id(string) StringBuilder
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

type CodeGenOperationBuilder interface {
	Build() *CodeGenOperation
	Id(string) CodeGenOperationBuilder
	Name(string) CodeGenOperationBuilder
	Description(string) CodeGenOperationBuilder
	Inputs(...CodeGenType) CodeGenOperationBuilder
	Outputs(...CodeGenType) CodeGenOperationBuilder
}
