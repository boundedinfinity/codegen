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

func (t floatEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.positive && t.negative {
		return ErrFloatEntityPosAndNegMutEx
	}

	if t.min > t.max {
		return ErrFloatEntityMinAboveMax.FormatFn("min: %v, max: v")(t.min, t.max)
	}

	if t.multipleOf < 1 {
		return ErrFloatEntityLessThan1.WithValue(t.multipleOf)
	}

	for i, rng := range t.ranges {
		return ErrFloatEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (t floatEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	iparam(data, "min", t.min)
	iparam(data, "min", t.max)
	iparam(data, "multiple-of", t.multipleOf)
	bparam(data, "positive", t.positive)
	bparam(data, "negative", t.negative)

	if len(t.oneOf) > 0 {
		data["one-of"] = t.oneOf
	}

	if len(t.ranges) > 0 {
		ranges := []map[string]any{}
		for _, irange := range t.ranges {
			ranges = append(ranges, map[string]any{
				"min": irange.Min,
				"max": irange.Max,
			})
		}
		data["ranges"] = ranges
	}

	if t.positive {
		data["positive"] = true
	}

	if t.negative {
		data["negative"] = true
	}

	return data, nil
}

func (t floatEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (t floatEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t floatEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t floatEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t floatEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t floatEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *floatEntity) QName(s string) *floatEntity            { t.qname = s; return t }
func (t *floatEntity) License(s License) *floatEntity         { t.license = s; return t }
func (t *floatEntity) Copyright(s string) *floatEntity        { t.copyright = s; return t }
func (t *floatEntity) Comments(s string) *floatEntity         { t.comments = s; return t }
func (t *floatEntity) LongDescription(s string) *floatEntity  { t.longDescription = s; return t }
func (t *floatEntity) ShortDescription(s string) *floatEntity { t.shortDescription = s; return t }
func (t *floatEntity) Serde(s string) *floatEntity            { t.serde = s; return t }
func (t *floatEntity) Json(s string) *floatEntity             { t.json = s; return t }
func (t *floatEntity) Yaml(s string) *floatEntity             { t.yaml = s; return t }
func (t *floatEntity) Sql(s string) *floatEntity              { t.sql = s; return t }

func (t *floatEntity) Required(b bool) *floatEntity             { t.required = b; return t }
func (t *floatEntity) Default(m map[string]any) *floatEntity    { t.defaultValue = m; return t }
func (t *floatEntity) AdditionalValidation(b bool) *floatEntity { t.additionalValidation = b; return t }

func (t *floatEntity) Positive() *floatEntity                  { t.positive = true; return t }
func (t *floatEntity) Negative() *floatEntity                  { t.negative = true; return t }
func (t *floatEntity) Min(n int) *floatEntity                  { t.min = n; return t }
func (t *floatEntity) Max(n int) *floatEntity                  { t.max = n; return t }
func (t *floatEntity) OneOf(n ...int) *floatEntity             { t.oneOf = n; return t }
func (t *floatEntity) Range(ranges ...FloatRange) *floatEntity { t.ranges = ranges; return t }

func (t *floatEntity) MultipleOf(n int) *floatEntity {
	t.multipleOf = n
	t.multipleOf = n
	return t
}
