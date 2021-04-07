package model

import "github.com/boundedinfinity/optional"

type OpenApiV310Extention struct {
	Module    optional.StringOptional `json:"module,omitempty" yaml:"module,omitempty" xml:"module,omitempty"`
	Version   optional.StringOptional `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	Input     optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output    optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
	Requires  []XBiGoRequires         `json:"requires,omitempty" yaml:"requires,omitempty" xml:"requires,omitempty"`
	Templates []XBiGoTemplate         `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type OpenApiV310ExtentionSchemas struct {
	Templates []XBiGoTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type OpenApiV310ExtentionSchema struct {
	Templates []XBiGoTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type XBiGoRequires struct {
	Package string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
}

type XBiGoTemplate struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
}

type GoLang struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Name    optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Schema  *JsonSchema_Draft07     `json:"schema,omitempty" yaml:"schema,omitempty" xml:"schema,omitempty"`
	Model   *OpenApiV310            `json:"model,omitempty" yaml:"model,omitempty" xml:"model,omitempty"`
}
