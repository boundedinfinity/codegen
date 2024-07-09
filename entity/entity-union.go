package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-jsonschema/idiomatic/json_schema"
)

func Union() *unionEntity {
	return &unionEntity{
		entityBase: entityBase{entityType: UnionType},
	}
}

var _ Marshalable = &unionEntity{}
var _ Validatable = &unionEntity{}

type unionEntity struct {
	entityBase
	entities []Entity
}

var (
	ErrUnionEntityMissingProps = errorer.New("missing properties")
	ErrUnionEntityInvalidProps = errorer.New("invalid properties")
)

func (t unionEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	return nil
}

func (t unionEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	entities := []map[string]any{}

	for _, entity := range t.entities {
		if edata, err := entity.ToMap(); err != nil {
			return data, err
		} else {
			entities = append(entities, edata)
		}
	}

	aparam(data, "entities", entities)

	return data, nil
}

func (t unionEntity) AsJsonSchema() (json_schema.JsonSchema, error) {
	schema := &json_schema.JsonSchemaArray{}
	return schema, nil
}

func (t unionEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t unionEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t unionEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t unionEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t unionEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *unionEntity) QName(s string) *unionEntity           { t.qname = s; return t }
func (t *unionEntity) License(s License) *unionEntity        { t.license = s; return t }
func (t *unionEntity) Copyright(s string) *unionEntity       { t.copyright = s; return t }
func (t *unionEntity) Comments(s string) *unionEntity        { t.comments = s; return t }
func (t *unionEntity) LongDescription(s string) *unionEntity { t.longDescription = s; return t }
func (t *unionEntity) Serde(s string) *unionEntity           { t.serde = s; return t }
func (t *unionEntity) Json(s string) *unionEntity            { t.json = s; return t }
func (t *unionEntity) Yaml(s string) *unionEntity            { t.yaml = s; return t }
func (t *unionEntity) Sql(s string) *unionEntity             { t.sql = s; return t }

func (t *unionEntity) Required(b bool) *unionEntity             { t.required = b; return t }
func (t *unionEntity) Default(m map[string]any) *unionEntity    { t.defaultValue = m; return t }
func (t *unionEntity) AdditionalValidation(b bool) *unionEntity { t.additionalValidation = b; return t }

func (t *unionEntity) Entity(elems ...Entity) *unionEntity {
	t.entities = append(t.entities, elems...)
	return t
}
