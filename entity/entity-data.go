package entity

func Data() *dataEntity {
	return &dataEntity{}
}

var _ Marshalable = &dataEntity{}
var _ Validatable = &dataEntity{}

type dataEntity struct {
	common
	entity Entity
	items  []map[string]any
}

func (f dataEntity) ToMap() (map[string]any, error) {
	data, err := f.common.ToMap()

	if err != nil {
		return data, err
	}

	entity, err := f.entity.ToMap()

	if err != nil {
		return data, err
	}

	mparam(data, "entity", entity)
	aparam(data, "items", f.items)

	return data, nil
}

func (t dataEntity) ToJson() ([]byte, error)       { return ToJson(t) }
func (t dataEntity) ToJsonIndent() ([]byte, error) { return ToJsonIndent(t) }
func (t dataEntity) ToYaml() ([]byte, error)       { return ToYaml(t) }

func (t dataEntity) Validate() error {
	if err := t.common.Validate(); err != nil {
		return err
	}

	return nil
}

func (t dataEntity) HasValidation() bool {
	return t.common.HasValidation()
}

func (t *dataEntity) QName(s string) *dataEntity            { t.qname = s; return t }
func (t *dataEntity) License(s License) *dataEntity         { t.license = s; return t }
func (t *dataEntity) Copyright(s string) *dataEntity        { t.copyright = s; return t }
func (t *dataEntity) Comments(s string) *dataEntity         { t.comments = s; return t }
func (t *dataEntity) LongDescription(s string) *dataEntity  { t.longDescription = s; return t }
func (t *dataEntity) ShortDescription(s string) *dataEntity { t.shortDescription = s; return t }
func (t *dataEntity) Serde(s string) *dataEntity            { t.serde = s; return t }
func (t *dataEntity) Json(s string) *dataEntity             { t.json = s; return t }
func (t *dataEntity) Yaml(s string) *dataEntity             { t.yaml = s; return t }
func (t *dataEntity) Sql(s string) *dataEntity              { t.sql = s; return t }

func (t *dataEntity) Item(v ...map[string]any) *dataEntity {
	t.items = append(t.items, v...)
	return t
}

func (t *dataEntity) Entity(v Entity) *dataEntity { t.entity = v; return t }
