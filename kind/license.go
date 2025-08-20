package kind

// https://docs.github.com/en/repositories/managing-your-repositorys-settings-and-features/customizing-your-repository/licensing-a-repository

// //////////////////////////////////////////////////////////////////////////
// License Type
// //////////////////////////////////////////////////////////////////////////

type LicenseType string

func (this LicenseType) String() string {
	return string(this)
}

func (this LicenseType) Name() string {
	return licenseType2Name[this]
}

const (
	None     LicenseType = "none"
	AFL30    LicenseType = "AFL-3.0"
	Apache20 LicenseType = "Apache-2.0"
)

var (
	licenseType2Name = map[LicenseType]string{
		None:     "",
		AFL30:    "Academic Free License v3.0",
		Apache20: "Apache license 2.0",
	}
)

// //////////////////////////////////////////////////////////////////////////
// License
// //////////////////////////////////////////////////////////////////////////

type License struct {
	Type LicenseType
	Text string
}

func NewLicenseText(s string) License {
	return License{
		Type: None,
		Text: s,
	}
}

func NewLicense(l LicenseType) License {
	return License{
		Type: l,
	}
}
