package entity

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func String() *stringEntity {
	return &stringEntity{
		entityBase: entityBase{entityType: StringType},
	}
}

var _ Entity = &stringEntity{}

type stringEntity struct {
	entityBase
	min        int
	max        int
	length     int
	regex      string
	includes   string
	startsWith string
	endsWith   string
}

var (
	ErrStringEntityMinAboveMax  = errorer.New("min above max")
	ErrStringEntityMinNegative  = errorer.New("min negative")
	ErrStringEntityInvalidRegex = errorer.New("invalid regex")
)

func (t stringEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.min > t.max {
		return ErrStringEntityMinAboveMax.FormatFn("min: %v, max: v")(t.min, t.max)
	}

	if t.min < 0 {
		return ErrStringEntityMinNegative.WithValue(t.min)
	}

	if t.regex != "" {
		_, err := regexp.Compile(t.regex)

		if err != nil {
			return ErrStringEntityInvalidRegex.Sub(err)
		}
	}

	return nil
}

func (t stringEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	iparam(data, "min", t.min)
	iparam(data, "min", t.max)
	iparam(data, "length", t.length)
	sparam(data, "regex", t.regex)
	sparam(data, "includes", t.includes)
	sparam(data, "starts-with", t.startsWith)
	sparam(data, "ends-with", t.endsWith)

	return data, nil
}

func (t stringEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (t stringEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t stringEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t stringEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t stringEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t stringEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *stringEntity) QName(s string) *stringEntity           { t.qname = s; return t }
func (t *stringEntity) License(s License) *stringEntity        { t.license = s; return t }
func (t *stringEntity) Copyright(s string) *stringEntity       { t.copyright = s; return t }
func (t *stringEntity) Comments(s string) *stringEntity        { t.comments = s; return t }
func (t *stringEntity) LongDescription(s string) *stringEntity { t.longDescription = s; return t }
func (t *stringEntity) Serde(s string) *stringEntity           { t.serde = s; return t }
func (t *stringEntity) Json(s string) *stringEntity            { t.json = s; return t }
func (t *stringEntity) Yaml(s string) *stringEntity            { t.yaml = s; return t }
func (t *stringEntity) Sql(s string) *stringEntity             { t.sql = s; return t }

func (t *stringEntity) Required(b bool) *stringEntity          { t.required = b; return t }
func (t *stringEntity) Default(m map[string]any) *stringEntity { t.defaultValue = m; return t }
func (t *stringEntity) AdditionalValidation(b bool) *stringEntity {
	t.additionalValidation = b
	return t
}

func (t *stringEntity) Min(n int) *stringEntity           { t.min = n; return t }
func (t *stringEntity) Max(n int) *stringEntity           { t.max = n; return t }
func (t *stringEntity) Length(n int) *stringEntity        { t.length = n; return t }
func (t *stringEntity) Regex(s string) *stringEntity      { t.regex = s; return t }
func (t *stringEntity) Includes(s string) *stringEntity   { t.includes = s; return t }
func (t *stringEntity) StartsWith(r string) *stringEntity { t.startsWith = r; return t }
func (t *stringEntity) EndsWith(r string) *stringEntity   { t.endsWith = r; return t }
