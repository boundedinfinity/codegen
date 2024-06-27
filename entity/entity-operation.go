package entity

func Operation() *operationEntity {
	return &operationEntity{}
}

var _ Marshalable = &operationEntity{}
var _ Validatable = &operationEntity{}

type operationEntity struct {
	common
	entity  Entity
	inputs  []Entity
	outputs []Entity
}

func (f operationEntity) ToMap() (map[string]any, error) {
	data, err := f.common.ToMap()

	if err != nil {
		return data, err
	}

	var inputs []map[string]any
	for _, input := range f.inputs {
		idata, err := input.ToMap()
		if err != nil {
			return data, err
		}

		inputs = append(inputs, idata)
	}
	aparam(data, "inputs", inputs)

	var outputs []map[string]any
	for _, output := range f.outputs {
		odata, err := output.ToMap()
		if err != nil {
			return data, err
		}

		outputs = append(outputs, odata)
	}
	aparam(data, "outputs", outputs)

	return data, nil
}

func (t operationEntity) ToJson() ([]byte, error)       { return ToJson(t) }
func (t operationEntity) ToJsonIndent() ([]byte, error) { return ToJsonIndent(t) }
func (t operationEntity) ToYaml() ([]byte, error)       { return ToYaml(t) }

func (t operationEntity) Validate() error {
	if err := t.common.Validate(); err != nil {
		return err
	}

	return nil
}

func (t operationEntity) HasValidation() bool {
	return t.common.HasValidation()
}

func (t *operationEntity) QName(s string) *operationEntity       { t.qname = s; return t }
func (t *operationEntity) License(s License) *operationEntity    { t.license = s; return t }
func (t *operationEntity) Copyright(s string) *operationEntity   { t.copyright = s; return t }
func (t *operationEntity) Comments(s string) *operationEntity    { t.comments = s; return t }
func (t *operationEntity) Description(s string) *operationEntity { t.description = s; return t }
func (t *operationEntity) Serde(s string) *operationEntity       { t.serde = s; return t }
func (t *operationEntity) Json(s string) *operationEntity        { t.json = s; return t }
func (t *operationEntity) Yaml(s string) *operationEntity        { t.yaml = s; return t }
func (t *operationEntity) Sql(s string) *operationEntity         { t.sql = s; return t }

func (t *operationEntity) Input(v ...Entity) *operationEntity {
	t.inputs = append(t.inputs, v...)
	return t
}
func (t *operationEntity) Outputs(v ...Entity) *operationEntity {
	t.outputs = append(t.outputs, v...)
	return t
}
func (t *operationEntity) Entity(v Entity) *operationEntity { t.entity = v; return t }
