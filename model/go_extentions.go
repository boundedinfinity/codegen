package model

import "github.com/boundedinfinity/optional"

type X_Bi_Go_Extention struct {
	TemplateRoot optional.StringOptional    `json:"templateRoot,omitempty" yaml:"templateRoot,omitempty" xml:"templateRoot,omitempty"`
	GenRoot      optional.StringOptional    `json:"genRoot,omitempty" yaml:"genRoot,omitempty" xml:"genRoot,omitempty"`
	Module       *X_Bi_Go_Module            `json:"module,omitempty" yaml:"module,omitempty" xml:"module,omitempty"`
	Templates    []*X_Bi_Go_Template        `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
	Components   *X_Bi_Go_Global_Components `json:"components,omitempty" yaml:"components,omitempty"`
	Paths        *X_Bi_Go_Global_Paths      `json:"paths,omitempty" yaml:"paths,omitempty" xml:"paths,omitempty"`
}

type X_Bi_Go_Module struct {
	Name     optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Version  optional.StringOptional `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	Requires []X_Bi_Go_Requires      `json:"requires,omitempty" yaml:"requires,omitempty" xml:"requires,omitempty"`
}

type X_Bi_Go_Requires struct {
	Package string `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Version string `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
}

type X_Bi_Go_Template struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
}

type X_Bi_Go_Global_Components struct {
	Schemas *X_Bi_Go_Global_Components_Schema `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

type X_Bi_Go_Global_Components_Schema struct {
	Package   optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Templates []*X_Bi_Go_Template     `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type X_Bi_Go_Components struct {
	Schemas *X_Bi_Go_Components_Schemas `json:"schemas,omitempty" yaml:"schemas,omitempty"`
}

type X_Bi_Go_Global_Paths struct {
	Package   optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Templates []*X_Bi_Go_Template     `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type X_Bi_Go_Components_Schemas struct {
	Package   optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Templates []*X_Bi_Go_Template     `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type X_Bi_Schema struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
}
