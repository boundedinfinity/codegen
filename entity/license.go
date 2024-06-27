package entity

// https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository

type LicenseType string

func (l LicenseType) String() string {
	return licenseType2Name[l]
}

const (
	NoLicenseType LicenseType = "none"
	AFL30         LicenseType = "AFL-3.0"
	Apache20      LicenseType = "Apache-2.0"
)

var (
	licenseType2Name = map[LicenseType]string{
		NoLicenseType: "",
		AFL30:         "Academic Free License v3.0",
		Apache20:      "Apache license 2.0",
	}
)

type License struct {
	Type LicenseType
	Text string
}

func (t License) ToMap() (map[string]any, error) {
	data := map[string]any{}

	sparam(data, "text", t.Text)
	sparam(data, "type", t.Type.String())

	return data, nil
}

func NewLicenseText(s string) License {
	return License{
		Type: NoLicenseType,
		Text: s,
	}
}

func NewLicense(l LicenseType) License {
	return License{
		Type: l,
	}
}
