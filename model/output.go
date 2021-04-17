package model

type BiOutput_TypeMap struct {
	Types map[string]string `json:"types,omitempty" yaml:"types,omitempty"`
}

type BiOutput_TemplateModelContext struct {
	Model BiOutput_Model `json:"model,omitempty" yaml:"model,omitempty"`
	Spec  BiOutput       `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiOutput struct {
	Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
	Info       BiOutput_Info  `json:"info,omitempty" yaml:"info,omitempty"`
	Models     BiGenModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations BiGenOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
}

type BiOutput_Info struct {
	InputDir    string `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir   string `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext bool   `json:"dumpContext" yaml:"dumpContext"`
}

type BiGenModel struct {
	Namespaces []BiOutput_Model_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenOperation struct {
	Namespaces []BiOutput_Operation_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiOutput_Template struct {
	Input          string `json:"input,omitempty" yaml:"input,omitempty"`
	Output         string `json:"output,omitempty" yaml:"output,omitempty"`
	InputLanguage  string `json:"inputLanguage,omitempty" yaml:"inputLanguage,omitempty"`
	OutputLanguage string `json:"outputLanguage,omitempty" yaml:"outputLanguage,omitempty"`
}

type BiOutput_Model_Namespace struct {
	Name       string                     `json:"name,omitempty" yaml:"name,omitempty"`
	Models     []BiOutput_Model           `json:"models,omitempty" yaml:"models,omitempty"`
	Namespaces []BiOutput_Model_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiOutput_Operation_Namespace struct {
	Name       string                         `json:"name,omitempty" yaml:"name,omitempty"`
	Operations []BiOutput_Operation           `json:"operations,omitempty" yaml:"operations,omitempty"`
	Namespaces []BiOutput_Operation_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiOutput_Operation struct {
	Name      string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string                  `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Imports   []string                `json:"imports,omitempty" yaml:"imports,omitempty"`
	Inputs    []BiOutput_TypeProperty `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs   []BiOutput_TypeProperty `json:"outputs,omitempty" yaml:"outputs,omitempty"`
	Templates []BiOutput_Template     `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiOutput_Model struct {
	Name       string                  `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace  string                  `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Type       string                  `json:"type,omitempty" yaml:"type,omitempty"`
	Imports    []string                `json:"imports,omitempty" yaml:"imports,omitempty"`
	Properties []BiOutput_TypeProperty `json:"properties,omitempty" yaml:"properties,omitempty"`
	Templates  []BiOutput_Template     `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiOutput_TypeProperty struct {
	Name      string `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Type      string `json:"type,omitempty" yaml:"type,omitempty"`
}
