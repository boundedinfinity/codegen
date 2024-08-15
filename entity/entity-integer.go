package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Integer() *integerEntity {
	return &integerEntity{
		entityBase: entityBase{entityType: IntegerType},
	}
}

var _ Entity = &integerEntity{}

type integerEntity struct {
	entityBase
	min        int
	max        int
	multipleOf int
	oneOf      []int
	positive   bool
	negative   bool
	ranges     []IntegerRange
}

type IntegerRange struct {
	Min int
	Max int
}

var (
	ErrIntegerEntityMinAboveMax    = errorer.New("min above max")
	ErrIntegerEntityLessThan1      = errorer.New("less than one")
	ErrIntegerEntityPosAndNegMutEx = errorer.New("positive and negative are mutually exclusive")
)

func (this integerEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.positive && this.negative {
		return ErrIntegerEntityPosAndNegMutEx
	}

	if this.min > this.max {
		return ErrIntegerEntityMinAboveMax.FormatFn("min: %v, max: v")(this.min, this.max)
	}

	if this.multipleOf < 1 {
		return ErrIntegerEntityLessThan1.WithValue(this.multipleOf)
	}

	for i, rng := range this.ranges {
		return ErrIntegerEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (this integerEntity) ToMap() (map[string]any, error) {
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

func (this integerEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaInteger{
		JsonSchemaCore: json_schema.JsonSchemaCore{
			Schema:      json_schema.SCHEMA_VERSION_2020_12,
			Id:          this.qname,
			Comment:     this.comments,
			Title:       this.shortDescription,
			Description: this.longDescription,
		},
	}
	return schema, nil
}

func (this integerEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this integerEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this integerEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this integerEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this integerEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *integerEntity) QName(s string) *integerEntity     { this.qname = s; return this }
func (this *integerEntity) License(s License) *integerEntity  { this.license = s; return this }
func (this *integerEntity) Copyright(s string) *integerEntity { this.copyright = s; return this }
func (this *integerEntity) Comments(s string) *integerEntity  { this.comments = s; return this }
func (this *integerEntity) LongDescription(s string) *integerEntity {
	this.longDescription = s
	return this
}
func (this *integerEntity) Serde(s string) *integerEntity { this.serde = s; return this }
func (this *integerEntity) Json(s string) *integerEntity  { this.json = s; return this }
func (this *integerEntity) Yaml(s string) *integerEntity  { this.yaml = s; return this }
func (this *integerEntity) Sql(s string) *integerEntity   { this.sql = s; return this }

func (this *integerEntity) Required(b bool) *integerEntity { this.required = b; return this }
func (this *integerEntity) Default(m map[string]any) *integerEntity {
	this.defaultValue = m
	return this
}
func (this *integerEntity) AdditionalValidation(b bool) *integerEntity {
	this.additionalValidation = b
	return this
}

func (this *integerEntity) Positive() *integerEntity      { this.positive = true; return this }
func (this *integerEntity) Negative() *integerEntity      { this.negative = true; return this }
func (this *integerEntity) Min(n int) *integerEntity      { this.min = n; return this }
func (this *integerEntity) Max(n int) *integerEntity      { this.max = n; return this }
func (this *integerEntity) OneOf(n ...int) *integerEntity { this.oneOf = n; return this }
func (this *integerEntity) Range(ranges ...IntegerRange) *integerEntity {
	this.ranges = ranges
	return this
}

func (this *integerEntity) MultipleOf(n int) *integerEntity {
	this.multipleOf = n
	this.multipleOf = n
	return this
}
