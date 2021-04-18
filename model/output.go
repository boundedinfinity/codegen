package model

type BiOutput_TypeMap struct {
	Types map[string]string `json:"types,omitempty" yaml:"types,omitempty"`
}

type BiOutput_TemplateModelContext struct {
	Model BiOutput_Model `json:"model,omitempty" yaml:"model,omitempty"`
	Spec  BiOutput       `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiOutput struct {
	Name       string               `json:"name,omitempty" yaml:"name,omitempty"`
	Info       BiOutput_Info        `json:"info,omitempty" yaml:"info,omitempty"`
	Models     []BiOutput_Model     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations []BiOutput_Operation `json:"operations,omitempty" yaml:"operations,omitempty"`
}

type BiOutput_Info struct {
	InputDir    string `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir   string `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext bool   `json:"dumpContext" yaml:"dumpContext"`
}

type BiOutput_Template struct {
	Input          string `json:"input,omitempty" yaml:"input,omitempty"`
	Output         string `json:"output,omitempty" yaml:"output,omitempty"`
	InputLanguage  string `json:"inputLanguage,omitempty" yaml:"inputLanguage,omitempty"`
	OutputLanguage string `json:"outputLanguage,omitempty" yaml:"outputLanguage,omitempty"`
}

type BiOutput_Operation struct {
	Name      string                `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string                `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Imports   []string              `json:"imports,omitempty" yaml:"imports,omitempty"`
	Input     BiOutput_TypeProperty `json:"input,omitempty" yaml:"input,omitempty"`
	Output    BiOutput_TypeProperty `json:"output,omitempty" yaml:"output,omitempty"`
	Templates []BiOutput_Template   `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiOutput_Model struct {
	Name       string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace  string                  `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Imports    []string                `json:"imports,omitempty" yaml:"imports,omitempty"`
	Properties []BiOutput_TypeProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
	Templates  []BiOutput_Template     `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiOutput_TypeProperty struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Type      string `json:"type,omitempty" yaml:"type,omitempty"`
}