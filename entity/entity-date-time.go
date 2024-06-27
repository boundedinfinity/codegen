package entity

import (
	"time"

	"github.com/boundedinfinity/go-commoner/errorer"
)

func DateTime() *dateTimeEntity {
	return &dateTimeEntity{
		entityBase: entityBase{entityType: DateTimeType},
	}
}

var _ Entity = &dateTimeEntity{}

type dateTimeEntity struct {
	entityBase
	min    time.Time
	max    time.Time
	oneOf  []time.Time
	ranges []DateTimeRange
}

var (
	ErrDateTimeEntityMinAboveMax = errorer.New("min above max")
	ErrDateTimeEntityLessThan1   = errorer.New("less than one")
)

func (t dateTimeEntity) Validate() error {

	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	var zero time.Time

	if t.min != zero && t.max != zero && t.min.After(t.max) {
		return ErrDateTimeEntityMinAboveMax.FormatFn("min: %v, max: v")(t.min, t.max)
	}

	for i, rng := range t.ranges {
		return ErrDateTimeEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (t dateTimeEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	tparam(data, "min", t.min)
	tparam(data, "max", t.max)
	tparams(data, "one-of", t.oneOf...)

	if len(t.ranges) > 0 {
		ranges := []map[string]any{}

		for _, rng := range t.ranges {
			rm := map[string]any{}
			tparam(rm, "min", rng.Min)
			tparam(rm, "max", rng.Max)
			ranges = append(ranges, rm)
		}

		if len(ranges) > 0 {
			data["ranges"] = ranges
		}
	}

	return data, nil
}

func (t dateTimeEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t dateTimeEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t dateTimeEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t dateTimeEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t dateTimeEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *dateTimeEntity) License(s License) *dateTimeEntity    { t.license = s; return t }
func (t *dateTimeEntity) Copyright(s string) *dateTimeEntity   { t.copyright = s; return t }
func (t *dateTimeEntity) Comments(s string) *dateTimeEntity    { t.comments = s; return t }
func (t *dateTimeEntity) Description(s string) *dateTimeEntity { t.description = s; return t }
func (t *dateTimeEntity) Serde(s string) *dateTimeEntity       { t.serde = s; return t }
func (t *dateTimeEntity) Json(s string) *dateTimeEntity        { t.json = s; return t }
func (t *dateTimeEntity) Yaml(s string) *dateTimeEntity        { t.yaml = s; return t }
func (t *dateTimeEntity) Sql(s string) *dateTimeEntity         { t.sql = s; return t }

func (t *dateTimeEntity) Required(b bool) *dateTimeEntity          { t.required = b; return t }
func (t *dateTimeEntity) Default(m map[string]any) *dateTimeEntity { t.defaultValue = m; return t }
func (t *dateTimeEntity) AdditionalValidation(b bool) *dateTimeEntity {
	t.additionalValidation = b
	return t
}

func (t *dateTimeEntity) Min(n time.Time) *dateTimeEntity      { t.min = n; return t }
func (t *dateTimeEntity) Max(n time.Time) *dateTimeEntity      { t.max = n; return t }
func (t *dateTimeEntity) OneOf(n ...time.Time) *dateTimeEntity { t.oneOf = n; return t }
func (t *dateTimeEntity) Range(ranges ...DateTimeRange) *dateTimeEntity {
	t.ranges = append(t.ranges, ranges...)
	return t
}

// ========================================================================================

type DateTimeRange struct {
	Min time.Time
	Max time.Time
}

var DateTimeRanges = dateTimeRanges{}

type dateTimeRanges struct {
}

func (t dateTimeRanges) NextWeekFrom(d time.Time) DateTimeRange {
	return DateTimeRange{
		Min: d,
		Max: d.Add(7 * 24 * time.Hour),
	}
}

func (t dateTimeRanges) NextWeekFromNow() DateTimeRange {
	return t.NextWeekFrom(time.Now())
}

func (t dateTimeRanges) PrevWeekFrom(d time.Time) DateTimeRange {
	return DateTimeRange{
		Min: d,
		Max: d.Add(-7 * 24 * time.Hour),
	}
}

func (t dateTimeRanges) PrevWeekFromNow() DateTimeRange {
	return t.PrevWeekFrom(time.Now())
}
