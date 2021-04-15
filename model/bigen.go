package model

type BiGenLangTypeMapper struct {
	Language map[string]map[string]string `json:"languages,omitempty" yaml:"languages,omitempty"`
}

type BiGenTemplateTypeContext struct {
	Type BiGenType `json:"type,omitempty" yaml:"type,omitempty"`
	Spec BiGen     `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiGen struct {
	Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
	Info       BiGenInfo      `json:"info,omitempty" yaml:"info,omitempty"`
	Models     BiGenModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations BiGenOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
}

type BiGenInfo struct {
	TemplateDir string `json:"templateDir,omitempty" yaml:"templateDir,omitempty"`
	OutputDir   string `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
}

type BiGenModel struct {
	Namespaces []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenOperation struct {
	Namespaces []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenTemplate struct {
	Input  string `json:"input,omitempty" yaml:"input,omitempty"`
	Output string `json:"output,omitempty" yaml:"output,omitempty"`
}

type BiGenNamespace struct {
	Name       string           `json:"name,omitempty" yaml:"name,omitempty"`
	Types      []BiGenType      `json:"types,omitempty" yaml:"types,omitempty"`
	Namespaces []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Templates  []BiGenTemplate  `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiGenType struct {
	Name       string              `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string              `json:"type,omitempty" yaml:"type,omitempty"`
	Properties []BiGenTypeProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
	Templates  []BiGenTemplate     `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiGenTypeProperty struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
	// Validations []BiSpecTypeValidation  `json:"validation,omitempty" yaml:"validation,omitempty" xml:"validation,omitempty"`
}
