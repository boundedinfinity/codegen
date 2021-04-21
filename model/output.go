package model

type BiOutput_TypeMap struct {
	Types map[string]string `json:"types,omitempty" yaml:"types,omitempty"`
}

type BiOutput_TemplateModelContext struct {
	Model BiOutput_Model `json:"model,omitempty" yaml:"model,omitempty"`
	Spec  BiOutput       `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiOutput_TemplateOperationContext struct {
	Operation BiOutput_Operation `json:"operation,omitempty" yaml:"operation,omitempty"`
	Spec      BiOutput           `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiOutput_TemplateNamespaceContext struct {
	Namespace         string   `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	RelativeNamespace string   `json:"relativeNamespace,omitempty" yaml:"relativeNamespace,omitempty"`
	Spec              BiOutput `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type BiOutput struct {
	Name       string                `json:"name,omitempty" yaml:"name,omitempty"`
	Info       BiOutput_Info         `json:"info,omitempty" yaml:"info,omitempty"`
	Models     []*BiOutput_Model     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations []*BiOutput_Operation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Namespaces []*BiOutput_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

func New_BiOutput() BiOutput {
	return BiOutput{
		Models:     make([]*BiOutput_Model, 0),
		Operations: make([]*BiOutput_Operation, 0),
		Namespaces: make([]*BiOutput_Namespace, 0),
	}
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
	Header         string `json:"header,omitempty" yaml:"header,omitempty"`
}

type BiOutput_Namespace struct {
	Namespace string               `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Children  []string             `json:"children,omitempty" yaml:"children,omitempty"`
	Templates []*BiOutput_Template `json:"templates,omitempty" yaml:"templates,omitempty"`
	// RelativeNamespace string              `json:"relativeNamespace,omitempty" yaml:"relativeNamespace,omitempty"`
}

func New_BiOutput_Namespace() *BiOutput_Namespace {
	return &BiOutput_Namespace{
		Children:  make([]string, 0),
		Templates: make([]*BiOutput_Template, 0),
	}
}

type BiOutput_Operation struct {
	Name        string               `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace   string               `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	SpecType    string               `json:"specType,omitempty" yaml:"specType,omitempty"`
	Description []string             `json:"description,omitempty" yaml:"description,omitempty"`
	Imports     []string             `json:"imports,omitempty" yaml:"imports,omitempty"`
	Input       BiOutput_Property    `json:"input,omitempty" yaml:"input,omitempty"`
	Output      BiOutput_Property    `json:"output,omitempty" yaml:"output,omitempty"`
	Templates   []*BiOutput_Template `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func New_BiOutput_Operation() BiOutput_Operation {
	return BiOutput_Operation{
		Description: make([]string, 0),
		Imports:     make([]string, 0),
		Templates:   make([]*BiOutput_Template, 0),
	}
}

type BiOutput_Model struct {
	Name         string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace    string                 `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	SpecName     string                 `json:"specName,omitempty" yaml:"specName,omitempty"`
	Description  []string               `json:"description,omitempty" yaml:"description,omitempty"`
	Imports      []string               `json:"imports,omitempty" yaml:"imports,omitempty"`
	JsonStruture map[string]interface{} `json:"jsonStruture,omitempty" yaml:"jsonStruture,omitempty"`
	Properties   []*BiOutput_Property   `json:"properties,omitempty" yaml:"properties,omitempty"`
	Templates    []*BiOutput_Template   `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func New_BiOutput_Model() *BiOutput_Model {
	return &BiOutput_Model{
		Description:  make([]string, 0),
		Imports:      make([]string, 0),
		JsonStruture: make(map[string]interface{}),
		Properties:   make([]*BiOutput_Property, 0),
		Templates:    make([]*BiOutput_Template, 0),
	}
}

type BiOutput_Property struct {
	Name         string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace    string                 `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	SpecName     string                 `json:"specName,omitempty" yaml:"specName,omitempty"`
	Type         string                 `json:"type,omitempty" yaml:"type,omitempty"`
	SpecType     string                 `json:"specType,omitempty" yaml:"specType,omitempty"`
	JsonPath     string                 `json:"jsonPath,omitempty" yaml:"jsonPath,omitempty"`
	JsonStruture map[string]interface{} `json:"jsonStruture,omitempty" yaml:"jsonStruture,omitempty"`
	Example      string                 `json:"example,omitempty" yaml:"example,omitempty"`
	Description  []string               `json:"description,omitempty" yaml:"description,omitempty"`
	Validations  []*BiOutput_Validation `json:"validations,omitempty" yaml:"validations,omitempty"`
}

func New_BiOutput_Property() *BiOutput_Property {
	return &BiOutput_Property{
		JsonStruture: make(map[string]interface{}),
		Description:  make([]string, 0),
		Validations:  make([]*BiOutput_Validation, 0),
	}
}

type BiOutput_Validation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}
