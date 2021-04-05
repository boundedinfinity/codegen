package model

type OpenApiV310Extention struct {
	Module    *string         `json:"module,omitempty" yaml:"module,omitempty" xml:"module,omitempty"`
	Version   *string         `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	Input     *string         `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Ouput     *string         `json:"ouput,omitempty" yaml:"ouput,omitempty" xml:"ouput,omitempty"`
	Requires  []XBiGoRequires `json:"requires,omitempty" yaml:"requires,omitempty" xml:"requires,omitempty"`
	Templates []XBiGoTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type OpenApiV310ExtentionSchemas struct {
	Templates []XBiGoTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type OpenApiV310ExtentionSchema struct {
	Package string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
}

type XBiGoRequires struct {
	Package string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
}

type XBiGoTemplate struct {
	Package *string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Input   *string `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  *string `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
}
