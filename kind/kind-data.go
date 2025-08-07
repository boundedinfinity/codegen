package kind

func Data() *dataEntity {
	return &dataEntity{}
}

var _ Marshalable = &dataEntity{}
var _ Validatable = &dataEntity{}

type dataEntity struct {
	common
	entity Kind
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

func (this dataEntity) ToJson() ([]byte, error)       { return ToJson(this) }
func (this dataEntity) ToJsonIndent() ([]byte, error) { return ToJsonIndent(this) }
func (this dataEntity) ToYaml() ([]byte, error)       { return ToYaml(this) }

func (this dataEntity) Validate() error {
	if err := this.common.Validate(); err != nil {
		return err
	}

	return nil
}

func (this dataEntity) HasValidation() bool {
	return this.common.HasValidation()
}

func (this *dataEntity) QName(s string) *dataEntity           { this.qname = s; return this }
func (this *dataEntity) License(s License) *dataEntity        { this.license = s; return this }
func (this *dataEntity) Copyright(s string) *dataEntity       { this.copyright = s; return this }
func (this *dataEntity) Comments(s string) *dataEntity        { this.comments = s; return this }
func (this *dataEntity) LongDescription(s string) *dataEntity { this.longDescription = s; return this }
func (this *dataEntity) ShortDescription(s string) *dataEntity {
	this.shortDescription = s
	return this
}
func (this *dataEntity) Serde(s string) *dataEntity { this.serde = s; return this }
func (this *dataEntity) Json(s string) *dataEntity  { this.json = s; return this }
func (this *dataEntity) Yaml(s string) *dataEntity  { this.yaml = s; return this }
func (this *dataEntity) Sql(s string) *dataEntity   { this.sql = s; return this }

func (this *dataEntity) Item(v ...map[string]any) *dataEntity {
	this.items = append(this.items, v...)
	return this
}

func (this *dataEntity) Entity(v Kind) *dataEntity { this.entity = v; return this }
