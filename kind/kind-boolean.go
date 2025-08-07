package kind

import (
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Boolean() *booleanEntity {
	return &booleanEntity{
		entityBase: entityBase{entityType: BooleanType},
	}
}

var _ Kind = &booleanEntity{}

type booleanEntity struct {
	entityBase
}

func (this booleanEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this booleanEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this booleanEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this booleanEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this booleanEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this booleanEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *booleanEntity) QName(s string) *booleanEntity     { this.qname = s; return this }
func (this *booleanEntity) License(s License) *booleanEntity  { this.license = s; return this }
func (this *booleanEntity) Copyright(s string) *booleanEntity { this.copyright = s; return this }
func (this *booleanEntity) Comments(s string) *booleanEntity  { this.comments = s; return this }
func (this *booleanEntity) LongDescription(s string) *booleanEntity {
	this.longDescription = s
	return this
}
func (this *booleanEntity) ShortDescription(s string) *booleanEntity {
	this.shortDescription = s
	return this
}
func (this *booleanEntity) Serde(s string) *booleanEntity { this.serde = s; return this }
func (this *booleanEntity) Json(s string) *booleanEntity  { this.json = s; return this }
func (this *booleanEntity) Yaml(s string) *booleanEntity  { this.yaml = s; return this }
func (this *booleanEntity) Sql(s string) *booleanEntity   { this.sql = s; return this }

func (this *booleanEntity) Required(b bool) *booleanEntity { this.required = b; return this }
func (this *booleanEntity) AdditionalValidation(b bool) *booleanEntity {
	this.additionalValidation = b
	return this
}
func (this *booleanEntity) Default(m map[string]any) *booleanEntity {
	this.defaultValue = m
	return this
}
