package entity

import (
	"time"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
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

func (this dateTimeEntity) Validate() error {

	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	var zero time.Time

	if this.min != zero && this.max != zero && this.min.After(this.max) {
		return ErrDateTimeEntityMinAboveMax.FormatFn("min: %v, max: v")(this.min, this.max)
	}

	for i, rng := range this.ranges {
		return ErrDateTimeEntityMinAboveMax.FormatFn("ranges[%v] min: %v, max: v")(i, rng.Min, rng.Max)
	}

	return nil
}

func (this dateTimeEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	tparam(data, "min", this.min)
	tparam(data, "max", this.max)
	tparams(data, "one-of", this.oneOf...)

	if len(this.ranges) > 0 {
		ranges := []map[string]any{}

		for _, rng := range this.ranges {
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

func (this dateTimeEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this dateTimeEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this dateTimeEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this dateTimeEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this dateTimeEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this dateTimeEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *dateTimeEntity) License(s License) *dateTimeEntity  { this.license = s; return this }
func (this *dateTimeEntity) Copyright(s string) *dateTimeEntity { this.copyright = s; return this }
func (this *dateTimeEntity) Comments(s string) *dateTimeEntity  { this.comments = s; return this }
func (this *dateTimeEntity) LongDescription(s string) *dateTimeEntity {
	this.longDescription = s
	return this
}
func (this *dateTimeEntity) ShortDescription(s string) *dateTimeEntity {
	this.shortDescription = s
	return this
}
func (this *dateTimeEntity) Serde(s string) *dateTimeEntity { this.serde = s; return this }
func (this *dateTimeEntity) Json(s string) *dateTimeEntity  { this.json = s; return this }
func (this *dateTimeEntity) Yaml(s string) *dateTimeEntity  { this.yaml = s; return this }
func (this *dateTimeEntity) Sql(s string) *dateTimeEntity   { this.sql = s; return this }

func (this *dateTimeEntity) Required(b bool) *dateTimeEntity { this.required = b; return this }
func (this *dateTimeEntity) Default(m map[string]any) *dateTimeEntity {
	this.defaultValue = m
	return this
}
func (this *dateTimeEntity) AdditionalValidation(b bool) *dateTimeEntity {
	this.additionalValidation = b
	return this
}

func (this *dateTimeEntity) Min(n time.Time) *dateTimeEntity      { this.min = n; return this }
func (this *dateTimeEntity) Max(n time.Time) *dateTimeEntity      { this.max = n; return this }
func (this *dateTimeEntity) OneOf(n ...time.Time) *dateTimeEntity { this.oneOf = n; return this }
func (this *dateTimeEntity) Range(ranges ...DateTimeRange) *dateTimeEntity {
	this.ranges = append(this.ranges, ranges...)
	return this
}

// ========================================================================================

type DateTimeRange struct {
	Min time.Time
	Max time.Time
}

var DateTimeRanges = dateTimeRanges{}

type dateTimeRanges struct {
}

func (this dateTimeRanges) NextWeekFrom(d time.Time) DateTimeRange {
	return DateTimeRange{
		Min: d,
		Max: d.Add(7 * 24 * time.Hour),
	}
}

func (this dateTimeRanges) NextWeekFromNow() DateTimeRange {
	return this.NextWeekFrom(time.Now())
}

func (this dateTimeRanges) PrevWeekFrom(d time.Time) DateTimeRange {
	return DateTimeRange{
		Min: d,
		Max: d.Add(-7 * 24 * time.Hour),
	}
}

func (this dateTimeRanges) PrevWeekFromNow() DateTimeRange {
	return this.PrevWeekFrom(time.Now())
}
