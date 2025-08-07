package kind

import (
	"regexp"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

type StringKindConfiguration struct {
	Name       string
	Min        int
	Max        int
	Regex      string
	StartsWith string
	EndsWith   string
}

type StringKind struct {
	Config StringKindConfiguration
	Source string
	QName  string
}

func String() *stringEntity {
	return &stringEntity{
		entityBase: entityBase{entityType: StringType},
	}
}

var _ Kind = &stringEntity{}

type stringEntity struct {
	entityBase
	min        int
	max        int
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

func (this stringEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.min > this.max {
		return ErrStringEntityMinAboveMax.FormatFn("min: %v, max: v")(this.min, this.max)
	}

	if this.min < 0 {
		return ErrStringEntityMinNegative.WithValue(this.min)
	}

	if this.regex != "" {
		_, err := regexp.Compile(this.regex)

		if err != nil {
			return ErrStringEntityInvalidRegex.Sub(err)
		}
	}

	return nil
}

func (this stringEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	iparam(data, "min", this.min)
	iparam(data, "min", this.max)
	iparam(data, "length", this.length)
	sparam(data, "regex", this.regex)
	sparam(data, "includes", this.includes)
	sparam(data, "starts-with", this.startsWith)
	sparam(data, "ends-with", this.endsWith)

	return data, nil
}

func (this stringEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this stringEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this stringEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this stringEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this stringEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this stringEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *stringEntity) QName(s string) *stringEntity     { this.qname = s; return this }
func (this *stringEntity) License(s License) *stringEntity  { this.license = s; return this }
func (this *stringEntity) Copyright(s string) *stringEntity { this.copyright = s; return this }
func (this *stringEntity) Comments(s string) *stringEntity  { this.comments = s; return this }
func (this *stringEntity) LongDescription(s string) *stringEntity {
	this.longDescription = s
	return this
}
func (this *stringEntity) Serde(s string) *stringEntity { this.serde = s; return this }
func (this *stringEntity) Json(s string) *stringEntity  { this.json = s; return this }
func (this *stringEntity) Yaml(s string) *stringEntity  { this.yaml = s; return this }
func (this *stringEntity) Sql(s string) *stringEntity   { this.sql = s; return this }

func (this *stringEntity) Required(b bool) *stringEntity          { this.required = b; return this }
func (this *stringEntity) Default(m map[string]any) *stringEntity { this.defaultValue = m; return this }
func (this *stringEntity) AdditionalValidation(b bool) *stringEntity {
	this.additionalValidation = b
	return this
}

func (this *stringEntity) Min(n int) *stringEntity           { this.min = n; return this }
func (this *stringEntity) Max(n int) *stringEntity           { this.max = n; return this }
func (this *stringEntity) Length(n int) *stringEntity        { this.length = n; return this }
func (this *stringEntity) Regex(s string) *stringEntity      { this.regex = s; return this }
func (this *stringEntity) Includes(s string) *stringEntity   { this.includes = s; return this }
func (this *stringEntity) StartsWith(r string) *stringEntity { this.startsWith = r; return this }
func (this *stringEntity) EndsWith(r string) *stringEntity   { this.endsWith = r; return this }
