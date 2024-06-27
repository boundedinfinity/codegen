package entity

import (
	"github.com/boundedinfinity/go-commoner/errorer"
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

func (t *refEntity) Validate() error {
	if err := t.entityBase.Validate(); err != nil {
		return err
	}

	if t.qname == "" {
		return ErrRefEntityMissingRef
	}

	return nil
}

func (t refEntity) ToMap() (map[string]any, error) {
	data, err := t.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	sparam(data, "ref", t.ref)

	return data, nil
}

func (t refEntity) ToJson() ([]byte, error)             { return ToJson(t) }
func (t refEntity) ToJsonIndent() ([]byte, error)       { return ToJsonIndent(t) }
func (t refEntity) ToYaml() ([]byte, error)             { return ToYaml(t) }
func (t refEntity) ToJsonSchema() ([]byte, error)       { return ToJsonIndent(t) }
func (t refEntity) ToJsonSchemaIndent() ([]byte, error) { return ToJsonSchemaIndent(t) }

func (t *refEntity) License(s License) *refEntity    { t.license = s; return t }
func (t *refEntity) Copyright(s string) *refEntity   { t.copyright = s; return t }
func (t *refEntity) Comments(s string) *refEntity    { t.comments = s; return t }
func (t *refEntity) Description(s string) *refEntity { t.description = s; return t }
func (t *refEntity) Serde(s string) *refEntity       { t.serde = s; return t }
func (t *refEntity) Json(s string) *refEntity        { t.json = s; return t }
func (t *refEntity) Yaml(s string) *refEntity        { t.yaml = s; return t }
func (t *refEntity) Sql(s string) *refEntity         { t.sql = s; return t }

func (t *refEntity) Required(b bool) *refEntity             { t.required = b; return t }
func (t *refEntity) Default(m map[string]any) *refEntity    { t.defaultValue = m; return t }
func (t *refEntity) AdditionalValidation(b bool) *refEntity { t.additionalValidation = b; return t }

func (t *refEntity) Ref(s string) *refEntity    { t.ref = s; return t }
func (t *refEntity) Entity(s Entity) *refEntity { return t.Ref(s.GetQName()) }
