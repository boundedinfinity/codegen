package kind

import (
	"github.com/boundedinfinity/go-commoner/errorer"
)

func Union() *unionEntity {
	return &unionEntity{
		entityBase: entityBase{entityType: UnionType},
	}
}

var _ Validator = &unionEntity{}

type unionEntity struct {
	entityBase
	entities []Kind
}

var (
	ErrUnionEntityMissingProps = errorer.New("missing properties")
	ErrUnionEntityInvalidProps = errorer.New("invalid properties")
)

func (this unionEntity) Validate() error {
	if err := this.entityBase.Validate(); err != nil {
		return err
	}

	return nil
}

func (this unionEntity) ToMap() (map[string]any, error) {
	data, err := this.entityBase.ToMap()

	if err != nil {
		return data, err
	}

	entities := []map[string]any{}

	for _, entity := range this.entities {
		if edata, err := entity.ToMap(); err != nil {
			return data, err
		} else {
			entities = append(entities, edata)
		}
	}

	aparam(data, "entities", entities)

	return data, nil
}

func (this *unionEntity) License(s License) *unionEntity  { this.license = s; return this }
func (this *unionEntity) Copyright(s string) *unionEntity { this.copyright = s; return this }
func (this *unionEntity) Comments(s string) *unionEntity  { this.comments = s; return this }
func (this *unionEntity) LongDescription(s string) *unionEntity {
	this.longDescription = s
	return this
}
func (this *unionEntity) Serde(s string) *unionEntity { this.serde = s; return this }
func (this *unionEntity) Json(s string) *unionEntity  { this.json = s; return this }
func (this *unionEntity) Yaml(s string) *unionEntity  { this.yaml = s; return this }
func (this *unionEntity) Sql(s string) *unionEntity   { this.sql = s; return this }

func (this *unionEntity) Required(b bool) *unionEntity          { this.required = b; return this }
func (this *unionEntity) Default(m map[string]any) *unionEntity { this.defaultValue = m; return this }
func (this *unionEntity) AdditionalValidation(b bool) *unionEntity {
	this.additionalValidation = b
	return this
}

func (this *unionEntity) Entity(elems ...Kind) *unionEntity {
	this.entities = append(this.entities, elems...)
	return this
}
