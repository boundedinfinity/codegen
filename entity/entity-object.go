package entity

import "github.com/boundedinfinity/go-commoner/errorer"

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

func (t *objectEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.props == nil {
		return ErrObjectEntityMissingProps
	}

	return nil
}

func (t objectEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	if t.props != nil {
		props := []map[string]any{}

		for _, prop := range t.props {
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

func (t objectEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t objectEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t objectEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t objectEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t objectEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *objectEntity) QName(s string) *objectEntity       { t.qname = s; return t }
func (t *objectEntity) License(s License) *objectEntity    { t.license = s; return t }
func (t *objectEntity) Copyright(s string) *objectEntity   { t.copyright = s; return t }
func (t *objectEntity) Comments(s string) *objectEntity    { t.comments = s; return t }
func (t *objectEntity) Description(s string) *objectEntity { t.description = s; return t }
func (t *objectEntity) Serde(s string) *objectEntity       { t.serde = s; return t }
func (t *objectEntity) Json(s string) *objectEntity        { t.json = s; return t }
func (t *objectEntity) Yaml(s string) *objectEntity        { t.yaml = s; return t }
func (t *objectEntity) Sql(s string) *objectEntity         { t.sql = s; return t }

func (t *objectEntity) Required(b bool) *objectEntity          { t.required = b; return t }
func (t *objectEntity) Default(m map[string]any) *objectEntity { t.defaultValue = m; return t }
func (t *objectEntity) AdditionalValidation(b bool) *objectEntity {
	t.additionalValidation = b
	return t
}

func (t *objectEntity) Properties(elems ...Entity) *objectEntity { t.props = elems; return t }
