package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Ref() *refEntity {
	return &refEntity{
		entityBase: entityBase{entityType: RefType},
	}
}

var _ Entity = &refEntity{}

type refEntity struct {
	entityBase
	ref string
}

var (
	ErrRefEntityMissingRef = errorer.New("missing ref")
	ErrRefEntityInvalidRef = errorer.New("invalid ref")
)

func (this refEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.qname == "" {
		return ErrRefEntityMissingRef
	}

	return nil
}

func (this refEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	sparam(data, "ref", this.ref)

	return data, nil
}

func (this refEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this refEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this refEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this refEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this refEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this refEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *refEntity) License(s License) *refEntity        { this.license = s; return this }
func (this *refEntity) Copyright(s string) *refEntity       { this.copyright = s; return this }
func (this *refEntity) Comments(s string) *refEntity        { this.comments = s; return this }
func (this *refEntity) LongDescription(s string) *refEntity { this.longDescription = s; return this }
func (this *refEntity) Serde(s string) *refEntity           { this.serde = s; return this }
func (this *refEntity) Json(s string) *refEntity            { this.json = s; return this }
func (this *refEntity) Yaml(s string) *refEntity            { this.yaml = s; return this }
func (this *refEntity) Sql(s string) *refEntity             { this.sql = s; return this }

func (this *refEntity) Required(b bool) *refEntity          { this.required = b; return this }
func (this *refEntity) Default(m map[string]any) *refEntity { this.defaultValue = m; return this }
func (this *refEntity) AdditionalValidation(b bool) *refEntity {
	this.additionalValidation = b
	return this
}

func (this *refEntity) Ref(s string) *refEntity    { this.ref = s; return this }
func (this *refEntity) Entity(s Entity) *refEntity { return this.Ref(s.GetQName()) }
