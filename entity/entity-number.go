package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
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

func (t *integerEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.positive && t.negative {
		return ErrIntegerEntityPosAndNegMutEx
	}

	if t.min > t.max {
		return ErrIntegerEntityMinAboveMax.FormatFn("min: %v, max: v")(t.min, t.max)
	}

	if t.multipleOf < 1 {
		return ErrIntegerEntityLessThan1.WithValue(t.multipleOf)
	}

	for i, rng := range t.ranges {
		return ErrIntegerEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (t integerEntity) ToMap() (map[string]any, error) {
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

func (t integerEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t integerEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t integerEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t integerEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t integerEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *integerEntity) QName(s string) *integerEntity       { t.qname = s; return t }
func (t *integerEntity) License(s License) *integerEntity    { t.license = s; return t }
func (t *integerEntity) Copyright(s string) *integerEntity   { t.copyright = s; return t }
func (t *integerEntity) Comments(s string) *integerEntity    { t.comments = s; return t }
func (t *integerEntity) Description(s string) *integerEntity { t.description = s; return t }
func (t *integerEntity) Serde(s string) *integerEntity       { t.serde = s; return t }
func (t *integerEntity) Json(s string) *integerEntity        { t.json = s; return t }
func (t *integerEntity) Yaml(s string) *integerEntity        { t.yaml = s; return t }
func (t *integerEntity) Sql(s string) *integerEntity         { t.sql = s; return t }

func (t *integerEntity) Required(b bool) *integerEntity          { t.required = b; return t }
func (t *integerEntity) Default(m map[string]any) *integerEntity { t.defaultValue = m; return t }
func (t *integerEntity) AdditionalValidation(b bool) *integerEntity {
	t.additionalValidation = b
	return t
}

func (t *integerEntity) Positive() *integerEntity                    { t.positive = true; return t }
func (t *integerEntity) Negative() *integerEntity                    { t.negative = true; return t }
func (t *integerEntity) Min(n int) *integerEntity                    { t.min = n; return t }
func (t *integerEntity) Max(n int) *integerEntity                    { t.max = n; return t }
func (t *integerEntity) OneOf(n ...int) *integerEntity               { t.oneOf = n; return t }
func (t *integerEntity) Range(ranges ...IntegerRange) *integerEntity { t.ranges = ranges; return t }

func (t *integerEntity) MultipleOf(n int) *integerEntity {
	t.multipleOf = n
	t.multipleOf = n
	return t
}
