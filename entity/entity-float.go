package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Float() *floatEntity {
	return &floatEntity{
		entityBase: entityBase{entityType: FloatType},
	}
}

var _ Entity = &floatEntity{}

type floatEntity struct {
	entityBase
	min        int
	max        int
	multipleOf int
	oneOf      []int
	positive   bool
	negative   bool
	ranges     []FloatRange
}

type FloatRange struct {
	Min int
	Max int
}

var (
	ErrFloatEntityMinAboveMax    = errorer.New("min above max")
	ErrFloatEntityLessThan1      = errorer.New("less than one")
	ErrFloatEntityPosAndNegMutEx = errorer.New("positive and negative are mutually exclusive")
)

func (this floatEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.positive && this.negative {
		return ErrFloatEntityPosAndNegMutEx
	}

	if this.min > this.max {
		return ErrFloatEntityMinAboveMax.FormatFn("min: %v, max: v")(this.min, this.max)
	}

	if this.multipleOf < 1 {
		return ErrFloatEntityLessThan1.WithValue(this.multipleOf)
	}

	for i, rng := range this.ranges {
		return ErrFloatEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (this floatEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	iparam(data, "min", this.min)
	iparam(data, "min", this.max)
	iparam(data, "multiple-of", this.multipleOf)
	bparam(data, "positive", this.positive)
	bparam(data, "negative", this.negative)

	if len(this.oneOf) > 0 {
		data["one-of"] = this.oneOf
	}

	if len(this.ranges) > 0 {
		ranges := []map[string]any{}
		for _, irange := range this.ranges {
			ranges = append(ranges, map[string]any{
				"min": irange.Min,
				"max": irange.Max,
			})
		}
		data["ranges"] = ranges
	}

	if this.positive {
		data["positive"] = true
	}

	if this.negative {
		data["negative"] = true
	}

	return data, nil
}

func (this floatEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this floatEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this floatEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this floatEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this floatEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this floatEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *floatEntity) QName(s string) *floatEntity     { this.qname = s; return this }
func (this *floatEntity) License(s License) *floatEntity  { this.license = s; return this }
func (this *floatEntity) Copyright(s string) *floatEntity { this.copyright = s; return this }
func (this *floatEntity) Comments(s string) *floatEntity  { this.comments = s; return this }
func (this *floatEntity) LongDescription(s string) *floatEntity {
	this.longDescription = s
	return this
}
func (this *floatEntity) ShortDescription(s string) *floatEntity {
	this.shortDescription = s
	return this
}
func (this *floatEntity) Serde(s string) *floatEntity { this.serde = s; return this }
func (this *floatEntity) Json(s string) *floatEntity  { this.json = s; return this }
func (this *floatEntity) Yaml(s string) *floatEntity  { this.yaml = s; return this }
func (this *floatEntity) Sql(s string) *floatEntity   { this.sql = s; return this }

func (this *floatEntity) Required(b bool) *floatEntity          { this.required = b; return this }
func (this *floatEntity) Default(m map[string]any) *floatEntity { this.defaultValue = m; return this }
func (this *floatEntity) AdditionalValidation(b bool) *floatEntity {
	this.additionalValidation = b
	return this
}

func (this *floatEntity) Positive() *floatEntity                  { this.positive = true; return this }
func (this *floatEntity) Negative() *floatEntity                  { this.negative = true; return this }
func (this *floatEntity) Min(n int) *floatEntity                  { this.min = n; return this }
func (this *floatEntity) Max(n int) *floatEntity                  { this.max = n; return this }
func (this *floatEntity) OneOf(n ...int) *floatEntity             { this.oneOf = n; return this }
func (this *floatEntity) Range(ranges ...FloatRange) *floatEntity { this.ranges = ranges; return this }

func (this *floatEntity) MultipleOf(n int) *floatEntity {
	this.multipleOf = n
	this.multipleOf = n
	return this
}
