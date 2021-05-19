package model

import (
	"github.com/iancoleman/orderedmap"
)

type OutputSpec struct {
	Models     []*OutputModel          `json:"models,omitempty" yaml:"models,omitempty"`
	Namespaces []*OutputNamespace      `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Operations []*OutputOperation      `json:"operations,omitempty" yaml:"operations,omitempty"`
	Info       OutputInfo              `json:"info,omitempty" yaml:"info,omitempty"`
	ModelMap   map[string]*OutputModel `json:"-" yaml:"-"`
}

func NewOutputSpec() *OutputSpec {
	return &OutputSpec{
		Models:     make([]*OutputModel, 0),
		Namespaces: make([]*OutputNamespace, 0),
	}
}

type OutputNamespace struct {
	Name      string            `json:"name,omitempty" yaml:"name,omitempty"`
	Templates []*OutputTemplate `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewOutputNamespace() *OutputNamespace {
	return &OutputNamespace{
		Templates: make([]*OutputTemplate, 0),
	}
}

type OutputModel struct {
	Name        string                 `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum         `json:"type,omitempty" yaml:"type,omitempty"`
	Description []string               `json:"description,omitempty" yaml:"description,omitempty"`
	Items       *OutputModel           `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  []OutputModel          `json:"properties,omitempty" yaml:"properties,omitempty"`
	Symbols     []string               `json:"symbols,omitempty" yaml:"symbols,omitempty"`
	Ref         string                 `json:"ref,omitempty" yaml:"ref,omitempty"`
	Example     interface{}            `json:"example,omitempty" yaml:"example,omitempty"`
	Json        *orderedmap.OrderedMap `json:"json,omitempty" yaml:"json,omitempty"`
	Imports     []string               `json:"imports,omitempty" yaml:"imports,omitempty"`
	Templates   []*OutputTemplate      `json:"templates,omitempty" yaml:"templates,omitempty"`
}

func NewOutputModel() *OutputModel {
	return &OutputModel{
		Description: make([]string, 0),
		Properties:  make([]OutputModel, 0),
		Templates:   make([]*OutputTemplate, 0),
		Symbols:     make([]string, 0),
		Imports:     make([]string, 0),
		Json:        orderedmap.New(),
	}
}

func NewOutputModelWithInput(input *InputModel) *OutputModel {
	m := NewOutputModel()

	if input == nil {
		return m
	}

	for _, property := range input.Properties {
		m.Properties = append(m.Properties, *NewOutputModelWithInput(&property))
	}

	m.Name = input.Name
	m.Type = input.Type
	m.Description = splitDescription(input.Description)
	m.Items = NewOutputModelWithInput(input.Items)
	m.Symbols = append(m.Symbols, input.Symbols...)
	m.Example = input.Example
	m.Ref = input.Ref

	return m
}

type OutputValidation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type OutputInfo struct {
	InputDir       string            `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir      string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext    bool              `json:"dumpContext" yaml:"dumpContext"`
	FilenameMarker string            `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
	Primitives     map[string]string `json:"primitives,omitempty" yaml:"primitives,omitempty"`
	Namespace      string            `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}

type OutputTemplate struct {
	Type           TemplateTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Input          string           `json:"input,omitempty" yaml:"input,omitempty"`
	Output         string           `json:"output,omitempty" yaml:"output,omitempty"`
	InputLanguage  string           `json:"inputLanguage,omitempty" yaml:"inputLanguage,omitempty"`
	OutputLanguage string           `json:"outputLanguage,omitempty" yaml:"outputLanguage,omitempty"`
	Header         []string         `json:"header,omitempty" yaml:"header,omitempty"`
}

func NewOutputTemplate() *OutputTemplate {
	return &OutputTemplate{
		Header: make([]string, 0),
	}
}

func NewOutputTemplateWithInput() *OutputTemplate {
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
