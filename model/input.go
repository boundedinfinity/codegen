package model

type InputSpec struct {
	Name          string         `json:"name,omitempty" yaml:"name,omitempty"`
	Version       string         `json:"version,omitempty" yaml:"version,omitempty"`
	Specification InputNamespace `json:"specification,omitempty" yaml:"specification,omitempty"`
}

func (t InputSpec) RootPackage() string {
	return t.Name
}

type InputModel struct {
	Name        string            `json:"name,omitempty" yaml:"name,omitempty"`
	Type        string            `json:"type,omitempty" yaml:"type,omitempty"`
	Description string            `json:"description,omitempty" yaml:"description,omitempty"`
	Example     interface{}       `json:"example,omitempty" yaml:"example,omitempty"`
	Properties  []InputModel      `json:"properties,omitempty" yaml:"properties,omitempty"`
	Validations []InputValidation `json:"validations,omitempty" yaml:"validations,omitempty"`
}

type InputValidation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type InputNamespace struct {
	Name       string           `json:"name,omitempty" yaml:"name,omitempty"`
	Models     []InputModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations []InputOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Namespaces []InputNamespace `json:"namespaces,omitempty" yaml:"namespaces,omitempty"`
}

type InputOperation struct {
	Name        string     `json:"name,omitempty" yaml:"name,omitempty"`
	Input       InputModel `json:"input,omitempty" yaml:"input,omitempty"`
	Output      InputModel `json:"output,omitempty" yaml:"output,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
}

type InputInfo struct {
	InputDir                  string `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir                 string `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext               bool   `json:"dumpContext" yaml:"dumpContext"`
	FilenameMarker            string `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
	TemplateHeader            string `json:"templateHeader,omitempty" yaml:"templateHeader,omitempty"`
	DescriptionSplitCharacter string `json:"descriptionSplitCharacter,omitempty" yaml:"descriptionSplitCharacter,omitempty"`
}

type InputTemplate struct {
	Header string `json:"header,omitempty" yaml:"header,omitempty"`
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}
