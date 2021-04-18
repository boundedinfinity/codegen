package model

type BiInput struct {
	Name          string            `json:"name,omitempty" yaml:"name,omitempty"`
	Version       string            `json:"version,omitempty" yaml:"version,omitempty"`
	Info          BiInput_Info      `json:"info,omitempty" yaml:"info,omitempty"`
	Specification BiInput_Namespace `json:"specification,omitempty" yaml:"specification,omitempty"`
}

type BiInput_Namespace struct {
	Name       string              `json:"name,omitempty" yaml:"name,omitempty"`
	Models     []BiInput_Model     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations []BiInput_Operation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Namespaces []BiInput_Namespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
	Templates  []BiInput_Template  `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type BiInput_Model struct {
	Name        string `json:"name,omitempty" yaml:"name,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
	Properties  []BiInput_Property
}

type BiInput_Property struct {
	Name        string               `json:"name,omitempty" yaml:"name,omitempty"`
	Type        string               `json:"type,omitempty" yaml:"type,omitempty"`
	Description string               `json:"description,omitempty" yaml:"description,omitempty"`
	Validations []BiInput_Validation `json:"validations,omitempty" yaml:"validations,omitempty"`
}

type BiInput_Validation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type BiInput_Operation struct {
	Name        string           `json:"name,omitempty" yaml:"name,omitempty"`
	Input       BiInput_Property `json:"input,omitempty" yaml:"input,omitempty"`
	Output      BiInput_Property `json:"output,omitempty" yaml:"output,omitempty"`
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
}

type BiInput_Info struct {
	InputDir       string            `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir      string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext    bool              `json:"dumpContext" yaml:"dumpContext"`
	FilenameMarker string            `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
	TemplateHeader string            `json:"templateHeader,omitempty" yaml:"templateHeader,omitempty"`
	TypeMap        map[string]string `json:"typeMap" yaml:"typeMap"`
}

type BiInput_Template struct {
	Header string `json:"header,omitempty" yaml:"header,omitempty"`
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}
