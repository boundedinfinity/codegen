package model

// Reference
// https://medium.com/@nate510/dynamic-json-umarshalling-in-go-88095561d6a0

type InputFile struct {
	Name          string             `json:"name,omitempty" yaml:"name,omitempty"`
	Version       string             `json:"version,omitempty" yaml:"version,omitempty"`
	Info          InputInfo          `json:"info,omitempty" yaml:"info,omitempty"`
	Specification InputSpecification `json:"specification,omitempty" yaml:"specification,omitempty"`
	Source        LanguageExtEnum
}

type InputSpecification struct {
	Models     []InputModel     `json:"models,omitempty" yaml:"models,omitempty"`
	Operations []InputOperation `json:"operations,omitempty" yaml:"operations,omitempty"`
	Templates  []InputTemplate  `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type InputModel struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Items       *InputModel    `json:"items,omitempty" yaml:"items,omitempty"`
	Properties  []InputModel   `json:"properties,omitempty" yaml:"properties,omitempty"`
	Symbols     []string       `json:"symbols,omitempty" yaml:"symbols,omitempty"`
	Ref         string         `json:"ref,omitempty" yaml:"ref,omitempty"`
	Example     interface{}    `json:"example,omitempty" yaml:"example,omitempty"`
}

type StringExample struct {
	Example string `json:"example,omitempty" yaml:"example,omitempty"`
}

type StringArrayExample struct {
	Example []string `json:"example,omitempty" yaml:"example,omitempty"`
}

type EnumExample struct {
	Example string `json:"example,omitempty" yaml:"example,omitempty"`
}

type EnumArrayExample struct {
	Example []string `json:"example,omitempty" yaml:"example,omitempty"`
}

type InputValidation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type InputOperation struct {
	Name        string     `json:"name,omitempty" yaml:"name,omitempty"`
	Input       InputModel `json:"input,omitempty" yaml:"input,omitempty"`
	Output      InputModel `json:"output,omitempty" yaml:"output,omitempty"`
	Description string     `json:"description,omitempty" yaml:"description,omitempty"`
}

type InputInfo struct {
	InputDir                  string            `json:"inputDir,omitempty" yaml:"inputDir,omitempty"`
	OutputDir                 string            `json:"outputDir,omitempty" yaml:"outputDir,omitempty"`
	DumpContext               bool              `json:"dumpContext" yaml:"dumpContext"`
	FilenameMarker            string            `json:"filenameMarker,omitempty" yaml:"filenameMarker,omitempty"`
	TemplateHeader            string            `json:"templateHeader,omitempty" yaml:"templateHeader,omitempty"`
	DescriptionSplitCharacter string            `json:"descriptionSplitCharacter,omitempty" yaml:"descriptionSplitCharacter,omitempty"`
	Primitives                map[string]string `json:"primitives,omitempty" yaml:"primitives,omitempty"`
}

type InputTemplate struct {
	Header    string           `json:"header,omitempty" yaml:"header,omitempty"`
	Path      string           `json:"path,omitempty" yaml:"path,omitempty"`
	Type      TemplateTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Namespace string           `json:"namespace,omitempty" yaml:"namespace,omitempty"`
}
