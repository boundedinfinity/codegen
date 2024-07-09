package entity

type common struct {
	qname            string
	shortDescription string
	longDescription  string
	comments         string
	copyright        string
	license          License
	serde            string
	json             string
	yaml             string
	sql              string
}

func (t common) Validate() error {
	return nil
}

func (t common) HasValidation() bool {
	return false
}

func (t common) ToMap() (map[string]any, error) {
	data := map[string]any{}

	sparam(data, "q-name", t.qname)
	sparam(data, "long-description", t.longDescription)
	sparam(data, "short-description", t.shortDescription)
	sparam(data, "serde", t.serde)
	sparam(data, "json", t.json)
	sparam(data, "yaml", t.yaml)
	sparam(data, "comments", t.comments)
	sparam(data, "copyright", t.copyright)

	if licenseData, err := t.license.ToMap(); err != nil {
		return data, err
	} else {
		mparam(data, "license", licenseData)
	}

	return data, nil
}

func (t common) GetQName() string { return t.qname }
