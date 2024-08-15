package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Object() *objectEntity {
	return &objectEntity{
		entityBase: entityBase{entityType: ObjectType},
	}
}

var _ Entity = &objectEntity{}

type objectEntity struct {
	entityBase
	props []Entity
}

var (
	ErrObjectEntityMissingProps = errorer.New("missing properties")
	ErrObjectEntityInvalidProps = errorer.New("invalid properties")
)

func (this objectEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	if this.props == nil {
		return ErrObjectEntityMissingProps
	}

	return nil
}

func (this objectEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	if this.props != nil {
		props := []map[string]any{}

		for _, prop := range this.props {
			pdata, err := prop.ToMap()
			if err != nil {
				return data, err
			}

			props = append(props, pdata)
		}

		data["props"] = props
	}

	return data, nil
}

func (this objectEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (this objectEntity) ToJson() ([]byte, error)             { return ToJson(this) }
func (this objectEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(this) }
func (this objectEntity) ToYaml() ([]byte, error)             { return ToYaml(this) }
func (this objectEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(this) }
func (this objectEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(this) }

func (this *objectEntity) QName(s string) *objectEntity     { this.qname = s; return this }
func (this *objectEntity) License(s License) *objectEntity  { this.license = s; return this }
func (this *objectEntity) Copyright(s string) *objectEntity { this.copyright = s; return this }
func (this *objectEntity) Comments(s string) *objectEntity  { this.comments = s; return this }
func (this *objectEntity) LongDescription(s string) *objectEntity {
	this.longDescription = s
	return this
}
func (this *objectEntity) Serde(s string) *objectEntity { this.serde = s; return this }
func (this *objectEntity) Json(s string) *objectEntity  { this.json = s; return this }
func (this *objectEntity) Yaml(s string) *objectEntity  { this.yaml = s; return this }
func (this *objectEntity) Sql(s string) *objectEntity   { this.sql = s; return this }

func (this *objectEntity) Required(b bool) *objectEntity          { this.required = b; return this }
func (this *objectEntity) Default(m map[string]any) *objectEntity { this.defaultValue = m; return this }
func (this *objectEntity) AdditionalValidation(b bool) *objectEntity {
	this.additionalValidation = b
	return this
}

func (this *objectEntity) Properties(elems ...Entity) *objectEntity { this.props = elems; return this }
