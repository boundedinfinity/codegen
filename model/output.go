package model

type OutputSpec struct {
	Models     []*OutputModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Namespaces []*OutputNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Operations []*OutputOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Info       OutputInfo         `json:"info,omitempty" yaml:"info,omitempty"`
}

func NewOutputSpec() *OutputSpec {
	return &OutputSpec{
		Models:     make([]*OutputModel, 0),
		Namespaces: make([]*OutputNamespace, 0),
	}
}

type OutputNamespace struct {
	SpecPath  string            `json:"specPath,omitempty" yaml:"specPath,omitempty"`
	Namespace string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Children  []string          `json:"children,omitempty" yaml:"children,omitempty"`
	Templates []*OutputTemplate `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewOutputNamespace() *OutputNamespace {
	return &OutputNamespace{
		Children:  make([]string, 0),
		Templates: make([]*OutputTemplate, 0),
	}
}

type OutputModel struct {
	SpecPath      string                 `json:"specPath,omitempty" yaml:"specPath,omitempty"`
	Namespace     string                 `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Name          string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Type          string                 `json:"type,omitempty" yaml:"type,omitempty"`
	Description   []string               `json:"description,omitempty" yaml:"description,omitempty"`
	Example       interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Imports       []string               `json:"imports,omitempty" yaml:"imports,omitempty"`
	Collection    bool                   `json:"collection,omitempty" yaml:"collection,omitempty"`
	JsonStructure map[string]interface{} `json:"jsonStructure,omitempty" yaml:"JsonStructure,omitempty"`
	Properties    []*OutputModel         `json:"properties,omitempty" yaml:"properties,omitempty"`
	Validations   []*OutputValidation    `json:"validations,omitempty" yaml:"validations,omitempty"`
	Templates     []*OutputTemplate      `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewOutputModel() *OutputModel {
	return &OutputModel{
		Description:   make([]string, 0),
		Imports:       make([]string, 0),
		JsonStructure: make(map[string]interface{}),
		Properties:    make([]*OutputModel, 0),
		Templates:     make([]*OutputTemplate, 0),
	}
}

type OutputValidation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type OutputInfo struct {
	InputDir       string `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir      string `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext    bool   `json:"dumpContext" yaml:"dumpContext"`
	FilenameMarker string `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
}

type OutputTemplate struct {
	Input          string   `json:"input,omitempty" yaml:"input,omitempty"`
	Output         string   `json:"output,omitempty" yaml:"output,omitempty"`
	InputLanguage  string   `json:"inputLanguage,omitempty" yaml:"inputLanguage,omitempty"`
	OutputLanguage string   `json:"outputLanguage,omitempty" yaml:"outputLanguage,omitempty"`
	Header         []string `json:"header,omitempty" yaml:"header,omitempty"`
}

func NewOutputTemplate() *OutputTemplate {
	return &OutputTemplate{
		Header: make([]string, 0),
	}
}

type OutputOperation struct {
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Namespace   string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	SpecPath    string            `json:"specPath,omitempty" yaml:"specPath,omitempty"`
	Description []string          `json:"description,omitempty" yaml:"description,omitempty"`
	Imports     []string          `json:"imports,omitempty" yaml:"imports,omitempty"`
	Input       OutputModel       `json:"input,omitempty" yaml:"input,omitempty"`
	Output      OutputModel       `json:"output,omitempty" yaml:"output,omitempty"`
	Templates   []*OutputTemplate `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewOutputOperation() *OutputOperation {
	return &OutputOperation{
		Description: make([]string, 0),
		Imports:     make([]string, 0),
		Templates:   make([]*OutputTemplate, 0),
	}
}

type OutputTemplateModelContext struct {
	Model OutputModel `json:"model,omitempty" yaml:"model,omitempty"`
	Spec  OutputSpec  `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type OutputTemplateOperationContext struct {
	Operation OutputOperation `json:"operation,omitempty" yaml:"operation,omitempty"`
	Spec      OutputSpec      `json:"spec,omitempty" yaml:"spec,omitempty"`
}

type OutputTemplateNamespaceContext struct {
	Namespace OutputNamespace `json:"namespace,omitempty" yaml:"namespace,omitempty"`
	Spec      OutputSpec      `json:"spec,omitempty" yaml:"spec,omitempty"`
}
