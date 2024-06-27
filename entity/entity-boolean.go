package entity

func Boolean() *booleanEntity {
	return &booleanEntity{
		entityBase: entityBase{entityType: BooleanType},
	}
}

var _ Entity = &booleanEntity{}

type booleanEntity struct {
	entityBase
}

func (t booleanEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t booleanEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t booleanEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t booleanEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t booleanEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *booleanEntity) QName(s string) *booleanEntity       { t.qname = s; return t }
func (t *booleanEntity) License(s License) *booleanEntity    { t.license = s; return t }
func (t *booleanEntity) Copyright(s string) *booleanEntity   { t.copyright = s; return t }
func (t *booleanEntity) Comments(s string) *booleanEntity    { t.comments = s; return t }
func (t *booleanEntity) Description(s string) *booleanEntity { t.description = s; return t }
func (t *booleanEntity) Serde(s string) *booleanEntity       { t.serde = s; return t }
func (t *booleanEntity) Json(s string) *booleanEntity        { t.json = s; return t }
func (t *booleanEntity) Yaml(s string) *booleanEntity        { t.yaml = s; return t }
func (t *booleanEntity) Sql(s string) *booleanEntity         { t.sql = s; return t }

func (t *booleanEntity) Required(b bool) *booleanEntity { t.required = b; return t }
func (t *booleanEntity) AdditionalValidation(b bool) *booleanEntity {
	t.additionalValidation = b
	return t
}
func (t *booleanEntity) Default(m map[string]any) *booleanEntity { t.defaultValue = m; return t }
