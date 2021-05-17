package model

import "github.com/boundedinfinity/optional"

//go:generate enumeration -package=model -name=JsonSchemaType -items=string,integer,number,object,array,boolean,null
//go:generate enumeration -package=model -name=JsonSchemaStringFormat -items=date-time,time,date,email,idn-hostname,ipv4,ipv6,uri,uri-reference,iri,iri-reference,uri-template,json-pointer,relative-json-pointer,regex

type JsonSchema_Draft07_Model struct {
	Type       JsonSchemaType                      `json:"type,omitempty" yaml:"type,omitempty"`
	Required   []string                            `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]JsonSchema_Draft07_Model `json:"properties,omitempty" yaml:"properties,omitempty"`
	Ref        optional.StringOptional             `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	X_Bi_Go    *X_Bi_Schema                        `json:"x-bi-go,omitempty" yaml:"x-bi-go,omitempty" xml:"x-bi-go,omitempty"`
}

type JsonSchema_Draft07_Items struct {
	Type            JsonSchemaType             `json:"type,omitempty" yaml:"type,omitempty"`
	Enum            []string                   `json:"enum,omitempty" yaml:"enum,omitempty"`
	AdditionalItems []JsonSchema_Draft07_Items `json:"additionalItems,omitempty" yaml:"additionalItems,omitempty"`
	Contains        []JsonSchema_Draft07_Items `json:"contains,omitempty" yaml:"contains,omitempty"`
}

type JsonSchema_Draft07_String struct {
	Type      JsonSchemaType         `json:"type,omitempty" yaml:"type,omitempty"`
	MinLength int                    `json:"minLength,omitempty" yaml:"minLength,omitempty"`
	MaxLength int                    `json:"maxLength,omitempty" yaml:"maxLength,omitempty"`
	Pattern   string                 `json:"pattern,omitempty" yaml:"pattern,omitempty"`
	Format    JsonSchemaStringFormat `json:"format,omitempty" yaml:"format,omitempty"`
}

type JsonSchema_Draft07_Integer struct {
	Type             JsonSchemaType         `json:"type,omitempty" yaml:"type,omitempty"`
	Minimum          optional.Int64Optional `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum          optional.Int64Optional `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	ExclusiveMinimum optional.BoolOptional  `json:"exclusiveMinimum,omitempty" yaml:"exclusiveMinimum,omitempty"`
	ExclusiveMaximum optional.BoolOptional  `json:"exclusiveMaximum,omitempty" yaml:"exclusiveMaximum,omitempty"`
	MultipleOf       int                    `json:"multipleOf,omitempty" yaml:"multipleOf,omitempty"`
}
