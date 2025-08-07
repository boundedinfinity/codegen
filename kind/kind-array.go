package kind

import (
	"github.com/boundedinfinity/go-commoner/errorer"
)

func Array() *arrayEntity {
	return &arrayEntity{
		entityBase: entityBase{entityType: ArrayType},
	}
}

var _ Kind = &arrayEntity{}

type arrayEntity struct {
	entityBase
	items    Kind
	min      int
	max      int
	length   int
	notEmpty bool
}

var (
	ErrArrayEntityMissingItems   = errorer.New("missing items")
	ErrArrayEntityInvalidItems   = errorer.New("invalid items")
	ErrArrayEntityMinAboveMax    = errorer.New("min above max")
	ErrArrayEntityMinNegative    = errorer.New("min negative")
	ErrArrayEntityLengthNegative = errorer.New("length negative")
)

func (this arrayEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.items == nil {
		return ErrArrayEntityMissingItems
	}

	if this.min > this.max {
		return ErrArrayEntityMinAboveMax.FormatFn("min: %v, max: v")(this.min, this.max)
	}

	if this.min < 0 {
		return ErrArrayEntityMinNegative.WithValue(this.min)
	}

	if this.length < 0 {
		return ErrArrayEntityLengthNegative.WithValue(this.min)
	}

	return nil
}

func (this arrayEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	if this.items != nil {
		idata, err := this.items.ToMap()
		if err != nil {
			return data, err
		}

		data["items"] = idata
	}

	if this.max != 0 {
		data["max"] = this.max
	}

	if this.max != 0 {
		data["max"] = this.max
	}

	return data, nil
}

func (this arrayEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this arrayEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this arrayEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this arrayEntity) ToJsonSchema() ([]byte, error)       { return ToJsonSchema(this) }
func (this arrayEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *arrayEntity) QName(s string) *arrayEntity     { this.qname = s; return this }
func (this *arrayEntity) License(s License) *arrayEntity  { this.license = s; return this }
func (this *arrayEntity) Copyright(s string) *arrayEntity { this.copyright = s; return this }
func (this *arrayEntity) Comments(s string) *arrayEntity  { this.comments = s; return this }
func (this *arrayEntity) LongDescription(s string) *arrayEntity {
	this.longDescription = s
	return this
}
func (this *arrayEntity) Serde(s string) *arrayEntity { this.serde = s; return this }
func (this *arrayEntity) Json(s string) *arrayEntity  { this.json = s; return this }
func (this *arrayEntity) Yaml(s string) *arrayEntity  { this.yaml = s; return this }
func (this *arrayEntity) Sql(s string) *arrayEntity   { this.sql = s; return this }

func (this *arrayEntity) Required(b bool) *arrayEntity { this.required = b; return this }
func (this *arrayEntity) AdditionalValidation(b bool) *arrayEntity {
	this.additionalValidation = b
	return this
}
func (this *arrayEntity) Default(m map[string]any) *arrayEntity { this.defaultValue = m; return this }

func (this *arrayEntity) Items(s Kind) *arrayEntity    { this.items = s; return this }
func (this *arrayEntity) NotEmpty(b bool) *arrayEntity { this.notEmpty = b; return this }
func (this *arrayEntity) Min(n int) *arrayEntity       { this.min = n; return this }
func (this *arrayEntity) Max(n int) *arrayEntity       { this.max = n; return this }
func (this *arrayEntity) Length(n int) *arrayEntity    { this.length = n; return this }
