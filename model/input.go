package model

type BiInput struct {
	Name          string                `json:"name,omitempty" yaml:"name,omitempty"`
	Version       string                `json:"version,omitempty" yaml:"version,omitempty"`
	Info          BiInput_Info          `json:"info,omitempty" yaml:"info,omitempty"`
	Specification BiInput_Specification `json:"specification,omitempty" yaml:"specification,omitempty"`
	Template      BiInput_Template      `json:"template,omitempty" yaml:"template,omitempty"`
}

type BiInput_Specification struct {
	Model     BiInput_Specification_Model     `json:"model,omitempty" yaml:"model,omitempty"`
	Operation BiInput_Specification_Operation `json:"operation,omitempty" yaml:"operation,omitempty"`
}

type BiInput_Specification_Operation struct {
	Namespaces []BiInput_Specification_Namespace_Operation `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiInput_Specification_Model struct {
	Namespaces []BiInput_Specification_Namespace_Model `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiInput_Specification_Namespace_Model struct {
	Name       string                                  `json:"name,omitempty" yaml:"name,omitempty"`
	Models     []BiInput_Model                         `json:"models,omitempty" yaml:"models,omitempty"`
	Namespaces []BiInput_Specification_Namespace_Model `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiInput_Model struct {
	Name       string `json:"name,omitempty" yaml:"name,omitempty"`
	Type       string `json:"type,omitempty" yaml:"type,omitempty"`
	Properties []BiInput_Model_Property
}

type BiInput_Model_Property struct {
	Name        string                              `json:"name,omitempty" yaml:"name,omitempty"`
	Type        string                              `json:"type,omitempty" yaml:"type,omitempty"`
	Validations []BiInput_Model_Property_Validation `json:"validations,omitempty" yaml:"validations,omitempty"`
}

type BiInput_Model_Property_Validation struct {
}

type BiInput_Specification_Namespace_Operation struct {
	Name       string                                      `json:"name,omitempty" yaml:"name,omitempty"`
	Operations []BiInput_Operation                         `json:"operations,omitempty" yaml:"operations,omitempty"`
	Namespaces []BiInput_Specification_Namespace_Operation `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type BiInput_Operation struct {
	Name    string   `json:"name,omitempty" yaml:"name,omitempty"`
	Inputs  []string `json:"inputs,omitempty" yaml:"inputs,omitempty"`
	Outputs []string `json:"outputs,omitempty" yaml:"outputs,omitempty"`
}

type BiInput_Info struct {
	InputDir    string            `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir   string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext bool              `json:"dumpContext" yaml:"dumpContext"`
	TypeMap     map[string]string `json:"typeMap" yaml:"typeMap"`
}

type BiInput_Template struct {
	Models     BiInput_Template_Models     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations BiInput_Template_Operations `json:"operations,omitempty" yaml:"operations,omitempty"`
}

type BiInput_Template_Models struct {
	Namespace []BiInput_Template_Info `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Model     []BiInput_Template_Info `json:"model,omitempty" yaml:"model,omitempty"`
}

type BiInput_Template_Operations struct {
	Namespace []BiInput_Template_Info `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Operation []BiInput_Template_Info `json:"operation,omitempty" yaml:"operation,omitempty"`
}

type BiInput_Template_Info struct {
	Input string `json:"input,omitempty" yaml:"input,omitempty"`
}
