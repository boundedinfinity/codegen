package kind

type common struct {
	qname            string
	shortDescription string
	longDescription  string
	comments         string
	copyright        string
	license          License
	serde            KindSerializationConfig
	json             KindSerializationConfig
	yaml             KindSerializationConfig
	sql              KindSerializationConfig
}

func (this common) Validate() error {
	return nil
}

func (this common) HasValidation() bool {
	return false
}

func (this common) ToMap() (map[string]any, error) {
	data := map[string]any{}

	sparam(data, "q-name", this.qname)
	sparam(data, "long-description", this.longDescription)
	sparam(data, "short-description", this.shortDescription)
	sparam(data, "serde", this.serde)
	sparam(data, "json", this.json)
	sparam(data, "yaml", this.yaml)
	sparam(data, "comments", this.comments)
	sparam(data, "copyright", this.copyright)

	if licenseData, err := this.license.ToMap(); err != nil {
		return data, err
	} else {
		mparam(data, "license", licenseData)
	}

	return data, nil
}

func (this common) GetQName() string { return this.qname }
