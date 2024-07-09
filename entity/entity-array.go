package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Array() *arrayEntity {
	return &arrayEntity{
		entityBase: entityBase{entityType: ArrayType},
	}
}

var _ Entity = &arrayEntity{}

type arrayEntity struct {
	entityBase
	items    Entity
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

func (t arrayEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.items == nil {
		return ErrArrayEntityMissingItems
	}

	if t.min > t.max {
		return ErrArrayEntityMinAboveMax.FormatFn("min: %v, max: v")(t.min, t.max)
	}

	if t.min < 0 {
		return ErrArrayEntityMinNegative.WithValue(t.min)
	}

	if t.length < 0 {
		return ErrArrayEntityLengthNegative.WithValue(t.min)
	}

	return nil
}

func (t arrayEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	if t.items != nil {
		idata, err := t.items.ToMap()
		if err != nil {
			return data, err
		}

		data["items"] = idata
	}

	if t.max != 0 {
		data["max"] = t.max
	}

	if t.max != 0 {
		data["max"] = t.max
	}

	return data, nil
}

func (t arrayEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (t arrayEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t arrayEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t arrayEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t arrayEntity) ToJsonSchema() ([]byte, error)       { return ToJsonSchema(t) }
func (t arrayEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *arrayEntity) QName(s string) *arrayEntity           { t.qname = s; return t }
func (t *arrayEntity) License(s License) *arrayEntity        { t.license = s; return t }
func (t *arrayEntity) Copyright(s string) *arrayEntity       { t.copyright = s; return t }
func (t *arrayEntity) Comments(s string) *arrayEntity        { t.comments = s; return t }
func (t *arrayEntity) LongDescription(s string) *arrayEntity { t.longDescription = s; return t }
func (t *arrayEntity) Serde(s string) *arrayEntity           { t.serde = s; return t }
func (t *arrayEntity) Json(s string) *arrayEntity            { t.json = s; return t }
func (t *arrayEntity) Yaml(s string) *arrayEntity            { t.yaml = s; return t }
func (t *arrayEntity) Sql(s string) *arrayEntity             { t.sql = s; return t }

func (t *arrayEntity) Required(b bool) *arrayEntity             { t.required = b; return t }
func (t *arrayEntity) AdditionalValidation(b bool) *arrayEntity { t.additionalValidation = b; return t }
func (t *arrayEntity) Default(m map[string]any) *arrayEntity    { t.defaultValue = m; return t }

func (t *arrayEntity) Items(s Entity) *arrayEntity  { t.items = s; return t }
func (t *arrayEntity) NotEmpty(b bool) *arrayEntity { t.notEmpty = b; return t }
func (t *arrayEntity) Min(n int) *arrayEntity       { t.min = n; return t }
func (t *arrayEntity) Max(n int) *arrayEntity       { t.max = n; return t }
func (t *arrayEntity) Length(n int) *arrayEntity    { t.length = n; return t }
