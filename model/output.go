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
	Namespaces []BiOutput_Model_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiGenTemplate struct {
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
	Name      string   `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace string   `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Inputs    []string `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Output    []string `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

type BiOutput_Model struct {
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
