package model

type BiSpec struct {
	Name       string          `json:"name,omitempty" yaml:"name,omitempty"`
	Version    string          `json:"version,omitempty" yaml:"version,omitempty"`
	Info       BiSpecInfo      `json:"info,omitempty" yaml:"info,omitempty"`
	Models     BiSpecModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations BiSpecOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
}

type BiSpecInfo struct {
	TemplateDir string            `json:"templateDir,omitempty" yaml:"templateDir,omitempty"`
	OutputDir   string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext bool              `json:"dumpContext" yaml:"dumpContext"`
	TypeMap     map[string]string `json:"typeMap" yaml:"typeMap"`
}

type BiSpecModel struct {
	Templates  []BiSpecTemplate  `json:"templates,omitempty" yaml:"templates,omitempty"`
	Namespaces []BiSpecNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiSpecOperation struct {
	Templates  []BiSpecTemplate  `json:"templates,omitempty" yaml:"templates,omitempty"`
	Namespaces []BiSpecNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiSpecTemplate struct {
	Input string `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
}

type BiSpecNamespace struct {
	Name       string            `json:"name,omitempty" yaml:"name,omitempty"`
	Types      []BiSpecType      `json:"types,omitempty" yaml:"types,omitempty"`
	Namespaces []BiSpecNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiSpecType struct {
	Name       string               `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string               `json:"type,omitempty" yaml:"type,omitempty"`
	Properties []BiSpecTypeProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type BiSpecTypeProperty struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	// Validations []BiSpecTypeValidation  `json:"validation,omitempty" yaml:"validation,omitempty" xml:"validation,omitempty"`
}
