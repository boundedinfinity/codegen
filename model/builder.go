package model

import "github.com/boundedinfinity/go-commoner/functional/optioner"

func setAndReturn[B any, V comparable](builder B, opt *optioner.Option[V], value V) B {
	*opt = optioner.OfZero(value)
	return builder
}

func appendAndReturn[B any, V comparable](builder B, opt *optioner.Option[[]V], value []V) B {
	var new []V
	new = append(new, opt.Get()...)
	new = append(new, value...)
	*opt = optioner.OfSlice(new)
	return builder
}

var Build = build{}

type build struct{}

type SchemaBuilder interface {
	Schema() CodeGenSchema
}

///////////////////////////////////////////////////////////////////
// String
//////////////////////////////////////////////////////////////////

func (this build) String() *stringBuilder {
	return &stringBuilder{}
}

type stringBuilder struct {
	typ CodeGenString
}

var _ SchemaBuilder = &stringBuilder{}

func (this *stringBuilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *stringBuilder) Id(id string) *stringBuilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *stringBuilder) Name(name string) *stringBuilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *stringBuilder) Description(description string) *stringBuilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *stringBuilder) Required(required bool) *stringBuilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *stringBuilder) JsonName(jsonName string) *stringBuilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *stringBuilder) YamlName(yamlName string) *stringBuilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *stringBuilder) SqlName(sqlName string) *stringBuilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *stringBuilder) Regex(regex string) *stringBuilder {
	return setAndReturn(this, &this.typ.Regex, regex)
}

func (this *stringBuilder) Abnf(abnf string) *stringBuilder {
	return setAndReturn(this, &this.typ.Abnf, abnf)
}

func (this *stringBuilder) Includes(includes ...string) *stringBuilder {
	return appendAndReturn(this, &this.typ.Includes, includes)
}

func (this *stringBuilder) Excludes(excludes ...string) *stringBuilder {
	return appendAndReturn(this, &this.typ.Excludes, excludes)
}

func (this *stringBuilder) OneOf(oneOf ...string) *stringBuilder {
	return appendAndReturn(this, &this.typ.OneOf, oneOf)
}

func (this *stringBuilder) NoneOf(noneOf ...string) *stringBuilder {
	return appendAndReturn(this, &this.typ.NoneOf, noneOf)
}

func (this *stringBuilder) Min(min int) *stringBuilder {
	return setAndReturn(this, &this.typ.Min, min)
}

func (this *stringBuilder) Max(max int) *stringBuilder {
	return setAndReturn(this, &this.typ.Max, max)
}

func (this *stringBuilder) Range(min, max int) *stringBuilder {
	return setAndReturn(this, &this.typ.Range,
		Range[int]{Min: optioner.Some(min), Max: optioner.Some(max)})
}

///////////////////////////////////////////////////////////////////
// Integer
//////////////////////////////////////////////////////////////////

func (this build) Integer() *integerBuilder {
	return &integerBuilder{}
}

type integerBuilder struct {
	typ CodeGenInteger
}

var _ SchemaBuilder = &integerBuilder{}

func (this *integerBuilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *integerBuilder) Id(id string) *integerBuilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *integerBuilder) Name(name string) *integerBuilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *integerBuilder) Description(description string) *integerBuilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *integerBuilder) Required(required bool) *integerBuilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *integerBuilder) JsonName(jsonName string) *integerBuilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *integerBuilder) YamlName(yamlName string) *integerBuilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *integerBuilder) SqlName(sqlName string) *integerBuilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *integerBuilder) OneOf(oneOf ...int) *integerBuilder {
	return appendAndReturn(this, &this.typ.OneOf, oneOf)
}

func (this *integerBuilder) NoneOf(noneOf ...int) *integerBuilder {
	return appendAndReturn(this, &this.typ.NoneOf, noneOf)
}

func (this *integerBuilder) Max(max int) *integerBuilder {
	return setAndReturn(this, &this.typ.Max, max)
}

func (this *integerBuilder) Min(min int) *integerBuilder {
	return setAndReturn(this, &this.typ.Min, min)
}

func (this *integerBuilder) MultipleOf(multipleOf int) *integerBuilder {
	return setAndReturn(this, &this.typ.MultipleOf, multipleOf)
}

func (this *integerBuilder) Range(min, max int) *integerBuilder {
	return appendAndReturn(this, &this.typ.Ranges, []Range[int]{
		{Min: optioner.Some(min), Max: optioner.Some(max)},
	})
}

func (this *integerBuilder) Positive(positive bool) *integerBuilder {
	return setAndReturn(this, &this.typ.Positive, positive)
}

func (this *integerBuilder) Negative(negative bool) *integerBuilder {
	return setAndReturn(this, &this.typ.Positive, negative)
}

///////////////////////////////////////////////////////////////////
// Float
//////////////////////////////////////////////////////////////////

func (this build) Float() *floatBuilder {
	return &floatBuilder{}
}

type floatBuilder struct {
	typ CodeGenFloat
}

var _ SchemaBuilder = &floatBuilder{}

func (this *floatBuilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *floatBuilder) Id(id string) *floatBuilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *floatBuilder) Name(name string) *floatBuilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *floatBuilder) Description(description string) *floatBuilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *floatBuilder) Required(required bool) *floatBuilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *floatBuilder) JsonName(jsonName string) *floatBuilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *floatBuilder) YamlName(yamlName string) *floatBuilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *floatBuilder) SqlName(sqlName string) *floatBuilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *floatBuilder) OneOf(oneOf ...float64) *floatBuilder {
	return appendAndReturn(this, &this.typ.OneOf, oneOf)
}

func (this *floatBuilder) NoneOf(noneOf ...float64) *floatBuilder {
	return appendAndReturn(this, &this.typ.NoneOf, noneOf)
}

func (this *floatBuilder) Max(max float64) *floatBuilder {
	return setAndReturn(this, &this.typ.Max, max)
}

func (this *floatBuilder) Min(min float64) *floatBuilder {
	return setAndReturn(this, &this.typ.Min, min)
}

func (this *floatBuilder) MultipleOf(multipleOf float64) *floatBuilder {
	return setAndReturn(this, &this.typ.MultipleOf, multipleOf)
}

func (this *floatBuilder) Range(min, max float64) *floatBuilder {
	return appendAndReturn(this, &this.typ.Ranges, []Range[float64]{
		{Min: optioner.Some(min), Max: optioner.Some(max)},
	})
}

func (this *floatBuilder) Positive(positive bool) *floatBuilder {
	return setAndReturn(this, &this.typ.Positive, positive)
}

func (this *floatBuilder) Negative(negative bool) *floatBuilder {
	return setAndReturn(this, &this.typ.Positive, negative)
}

///////////////////////////////////////////////////////////////////
// Boolean
//////////////////////////////////////////////////////////////////

func (this build) Boolean() *booleanBuilder {
	return &booleanBuilder{}
}

type booleanBuilder struct {
	typ CodeGenBoolean
}

var _ SchemaBuilder = &booleanBuilder{}

func (this *booleanBuilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *booleanBuilder) Id(id string) *booleanBuilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *booleanBuilder) Name(name string) *booleanBuilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *booleanBuilder) Description(description string) *booleanBuilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *booleanBuilder) Required(required bool) *booleanBuilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *booleanBuilder) JsonName(jsonName string) *booleanBuilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *booleanBuilder) YamlName(yamlName string) *booleanBuilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *booleanBuilder) SqlName(sqlName string) *booleanBuilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

///////////////////////////////////////////////////////////////////
// Boolean
//////////////////////////////////////////////////////////////////

func (this build) Enum() *enumBuilder {
	return &enumBuilder{}
}

type enumBuilder struct {
	typ CodeGenEnum
}

var _ SchemaBuilder = &enumBuilder{}

func (this *enumBuilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *enumBuilder) Id(id string) *enumBuilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *enumBuilder) Name(name string) *enumBuilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *enumBuilder) Description(description string) *enumBuilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *enumBuilder) Required(required bool) *enumBuilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *enumBuilder) JsonName(jsonName string) *enumBuilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *enumBuilder) YamlName(yamlName string) *enumBuilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *enumBuilder) SqlName(sqlName string) *enumBuilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *enumBuilder) Values(name, description string, items ...string) *enumBuilder {
	item := CodeGenEnumItem{
		Name:        optioner.OfZero(name),
		Description: optioner.OfZero(description),
		Items:       optioner.OfSlice(items),
	}
	return appendAndReturn(this, &this.typ.Values, []CodeGenEnumItem{item})
}

///////////////////////////////////////////////////////////////////
// Object
//////////////////////////////////////////////////////////////////

func (this build) Object() *objectBilder {
	return &objectBilder{}
}

type objectBilder struct {
	typ CodeGenObject
}

var _ SchemaBuilder = &objectBilder{}

func (this *objectBilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *objectBilder) Id(id string) *objectBilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *objectBilder) Name(name string) *objectBilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *objectBilder) Description(description string) *objectBilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *objectBilder) Required(required bool) *objectBilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *objectBilder) JsonName(jsonName string) *objectBilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *objectBilder) YamlName(yamlName string) *objectBilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *objectBilder) SqlName(sqlName string) *objectBilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *objectBilder) Property(property SchemaBuilder) *objectBilder {
	return appendAndReturn(this, &this.typ.Properties, []CodeGenSchema{property.Schema()})
}

///////////////////////////////////////////////////////////////////
// Array
//////////////////////////////////////////////////////////////////

func (this build) Array() *arrayBilder {
	return &arrayBilder{}
}

type arrayBilder struct {
	typ CodeGenArray
}

var _ SchemaBuilder = &arrayBilder{}

func (this *arrayBilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *arrayBilder) Id(id string) *arrayBilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *arrayBilder) Name(name string) *arrayBilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *arrayBilder) Description(description string) *arrayBilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *arrayBilder) Required(required bool) *arrayBilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *arrayBilder) JsonName(jsonName string) *arrayBilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *arrayBilder) YamlName(yamlName string) *arrayBilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *arrayBilder) SqlName(sqlName string) *arrayBilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *arrayBilder) Items(property SchemaBuilder) *arrayBilder {
	return setAndReturn(this, &this.typ.Items, property.Schema())
}

///////////////////////////////////////////////////////////////////
// Ref
//////////////////////////////////////////////////////////////////

func (this build) Ref() *refBilder {
	return &refBilder{}
}

type refBilder struct {
	typ CodeGenRef
}

var _ SchemaBuilder = &refBilder{}

func (this *refBilder) Schema() CodeGenSchema {
	return &this.typ
}

func (this *refBilder) Id(id string) *refBilder {
	return setAndReturn(this, &this.typ.Id, id)
}

func (this *refBilder) Name(name string) *refBilder {
	return setAndReturn(this, &this.typ.Name, name)
}

func (this *refBilder) Description(description string) *refBilder {
	return setAndReturn(this, &this.typ.Description, description)
}

func (this *refBilder) Required(required bool) *refBilder {
	return setAndReturn(this, &this.typ.Required, required)
}

func (this *refBilder) JsonName(jsonName string) *refBilder {
	return setAndReturn(this, &this.typ.JsonName, jsonName)
}

func (this *refBilder) YamlName(yamlName string) *refBilder {
	return setAndReturn(this, &this.typ.YamlName, yamlName)
}

func (this *refBilder) SqlName(sqlName string) *refBilder {
	return setAndReturn(this, &this.typ.SqlName, sqlName)
}

func (this *refBilder) Ref(id string) *refBilder {
	return setAndReturn(this, &this.typ.Ref, id)
}

func (this *refBilder) RefSchema(ref SchemaBuilder) *refBilder {
	return this.Ref(ref.Schema().Common().Id.Get())
}
