package model

import "github.com/boundedinfinity/optional"

type X_Bi_GlobalExtention struct {
	TemplateRoot optional.StringOptional `json:"templateRoot,omitempty" yaml:"templateRoot,omitempty" xml:"templateRoot,omitempty"`
	GenRoot      optional.StringOptional `json:"genRoot,omitempty" yaml:"genRoot,omitempty" xml:"genRoot,omitempty"`
	Module       *X_Bi_Module            `json:"module,omitempty" yaml:"module,omitempty" xml:"module,omitempty"`
	Templates    []*X_Bi_GoTemplate      `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type X_Bi_Module struct {
	Name     optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Version  optional.StringOptional `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	Requires []X_Bi_GoRequires       `json:"requires,omitempty" yaml:"requires,omitempty" xml:"requires,omitempty"`
}

type X_Bi_GoRequires struct {
	Package string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
}

type X_Bi_GoTemplate struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
}

type X_Bi_GoComponents struct {
	Schemas *X_Bi_GoComponents_Schemas `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

type X_Bi_GoComponents_Schemas struct {
	Templates []*X_Bi_GoTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type X_Bi_Schema struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
}
