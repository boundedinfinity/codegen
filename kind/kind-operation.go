package kind

func Operation() *operationEntity {
	return &operationEntity{}
}

var _ Marshalable = &operationEntity{}
var _ Validator = &operationEntity{}

type operationEntity struct {
	common
	entity  Kind
	inputs  []Kind
	outputs []Kind
}

func (this operationEntity) ToMap() (map[string]any, error) {
	data, err := this.common.ToMap()

	if err != nil {
		return data, err
	}

	var inputs []map[string]any
	for _, input := range this.inputs {
		idata, err := input.ToMap()
		if err != nil {
			return data, err
		}

		inputs = append(inputs, idata)
	}
	aparam(data, "inputs", inputs)

	var outputs []map[string]any
	for _, output := range this.outputs {
		odata, err := output.ToMap()
		if err != nil {
			return data, err
		}

		outputs = append(outputs, odata)
	}
	aparam(data, "outputs", outputs)

	return data, nil
}

func (this operationEntity) ToJson() ([]byte, error)       { return ToJson(this) }
func (this operationEntity) ToJsonIndent() ([]byte, error) { return ToJsonIndent(this) }
func (this operationEntity) ToYaml() ([]byte, error)       { return ToYaml(this) }

func (this operationEntity) Validate() error {
	if err := this.common.Validate(); err != nil {
		return err
	}

	return nil
}

func (this operationEntity) HasValidation() bool {
	return this.common.HasValidation()
}

func (this *operationEntity) QName(s string) *operationEntity     { this.qname = s; return this }
func (this *operationEntity) License(s License) *operationEntity  { this.license = s; return this }
func (this *operationEntity) Copyright(s string) *operationEntity { this.copyright = s; return this }
func (this *operationEntity) Comments(s string) *operationEntity  { this.comments = s; return this }
func (this *operationEntity) LongDescription(s string) *operationEntity {
	this.longDescription = s
	return this
}
func (this *operationEntity) Serde(s string) *operationEntity { this.serde = s; return this }
func (this *operationEntity) Json(s string) *operationEntity  { this.json = s; return this }
func (this *operationEntity) Yaml(s string) *operationEntity  { this.yaml = s; return this }
func (this *operationEntity) Sql(s string) *operationEntity   { this.sql = s; return this }

func (this *operationEntity) Input(v ...Kind) *operationEntity {
	this.inputs = append(this.inputs, v...)
	return this
}
func (this *operationEntity) Outputs(v ...Kind) *operationEntity {
	this.outputs = append(this.outputs, v...)
	return this
}
func (this *operationEntity) Entity(v Kind) *operationEntity { this.entity = v; return this }
