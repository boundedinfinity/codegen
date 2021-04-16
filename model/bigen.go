package model

type BiGenTypeMap struct {
	Types map[string]string `json:"types,omitempty" yaml:"types,omitempty"`
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
	DumpContext bool   `json:"dumpContext" yaml:"dumpContext"`
}

type BiGenModel struct {
	Namespaces []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenOperation struct {
	Namespaces []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenTemplate struct {
	Input          string `json:"input,omitempty" yaml:"input,omitempty"`
	Output         string `json:"output,omitempty" yaml:"output,omitempty"`
	InputLanguage  string `json:"inputLanguage,omitempty" yaml:"inputLanguage,omitempty"`
	OutputLanguage string `json:"outputLanguage,omitempty" yaml:"outputLanguage,omitempty"`
}

type BiGenNamespace struct {
	QualifiedName string           `json:"qualifiedName,omitempty" yaml:"qualifiedName,omitempty"`
	Types         []BiGenType      `json:"types,omitempty" yaml:"types,omitempty"`
	Namespaces    []BiGenNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Templates     []BiGenTemplate  `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiGenType struct {
	Name       string              `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace  string              `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Type       string              `json:"type,omitempty" yaml:"type,omitempty"`
	Imports    []string            `json:"imports,omitempty" yaml:"imports,omitempty"`
	Properties []BiGenTypeProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
	Templates  []BiGenTemplate     `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiGenTypeProperty struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Type      string `json:"type,omitempty" yaml:"type,omitempty"`
}
