package model

import "github.com/boundedinfinity/optional"

type BiSpecGenerationTemplateGlobalContext struct {
	Spec BiSpec `json:"spec,omitempty" yaml:"spec,omitempty" xml:"spec,omitempty"`
}

type BiSpecGenerationTemplateTypeContext struct {
	Type BiSpecType `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	Spec BiSpec     `json:"spec,omitempty" yaml:"spec,omitempty" xml:"spec,omitempty"`
}

type BiSpecGenerationTemplateOperationContext struct {
	Operation BiSpecOperation `json:"operation,omitempty" yaml:"operation,omitempty" xml:"operation,omitempty"`
	Spec      BiSpec          `json:"spec,omitempty" yaml:"spec,omitempty" xml:"spec,omitempty"`
}

type BiSpec struct {
	Name       optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Version    optional.StringOptional `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	Types      []BiSpecType            `json:"types,omitempty" yaml:"types,omitempty" xml:"types,omitempty"`
	Operations []BiSpecOperation       `json:"operations,omitempty" yaml:"operations,omitempty" xml:"operations,omitempty"`
	Generation *BiSpecGeneration       `json:"generation,omitempty" yaml:"generation,omitempty" xml:"generation,omitempty"`
}

type BiSpecGeneration struct {
	Version     optional.StringOptional       `json:"version,omitempty" yaml:"version,omitempty" xml:"version,omitempty"`
	TemplateDir optional.StringOptional       `json:"templateDir,omitempty" yaml:"templateDir,omitempty" xml:"templateDir,omitempty"`
	OutputDir   optional.StringOptional       `json:"outputDir,omitempty" yaml:"outputDir,omitempty" xml:"outputDir,omitempty"`
	Global      *BiSpecGenerationTemplateInfo `json:"global,omitempty" yaml:"global,omitempty" xml:"global,omitempty"`
	Types       *BiSpecGenerationTemplateInfo `json:"types,omitempty" yaml:"types,omitempty" xml:"types,omitempty"`
	Operations  *BiSpecGenerationTemplateInfo `json:"operations,omitempty" yaml:"operations,omitempty" xml:"operations,omitempty"`
}

type BiSpecGenerationTemplateInfo struct {
	Templates []BiSpecGenerationTemplate `json:"templates,omitempty" yaml:"templates,omitempty" xml:"templates,omitempty"`
}

type BiSpecGenerationTemplate struct {
	Input optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
}

type BiSpecOperation struct {
	Name   optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Input  BiSpecOperationInput    `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output BiSpecOperationOutput   `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
}

type BiSpecOperationInput struct {
	Name        optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Type        optional.StringOptional `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	Description optional.StringOptional `json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
}

type BiSpecOperationOutput struct {
	Name        optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Type        optional.StringOptional `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	Description optional.StringOptional `json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
}

type BiSpecType struct {
	Name        optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Type        optional.StringOptional `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	Description optional.StringOptional `json:"description,omitempty" yaml:"description,omitempty" xml:"description,omitempty"`
	Properties  []BiSpecTypeProperty    `json:"properties,omitempty" yaml:"properties,omitempty" xml:"properties,omitempty"`
}

type BiSpecTypeProperty struct {
	Name        optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Type        optional.StringOptional `json:"type,omitempty" yaml:"type,omitempty" xml:"type,omitempty"`
	Validations []BiSpecTypeValidation  `json:"validation,omitempty" yaml:"validation,omitempty" xml:"validation,omitempty"`
}

type BiSpecTypeValidation struct {
	Minimum optional.StringOptional `json:"minimum,omitempty" yaml:"minimum,omitempty" xml:"minimum,omitempty"`
	Maximum optional.StringOptional `json:"maximum,omitempty" yaml:"maximum,omitempty" xml:"maximum,omitempty"`
}

type BiSpecLangTypeMapper struct {
	Language map[string]map[string]string
}
