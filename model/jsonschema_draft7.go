package model

type JsonSchema_Draft07 struct {
	Type       *string                       `json:"type,omitempty" yaml:"type,omitempty"`
	Required   []string                      `json:"required,omitempty" yaml:"required,omitempty"`
	Properties map[string]JsonSchema_Draft07 `json:"properties,omitempty" yaml:"properties,omitempty"`
	Ref        *string                       `json:"$ref,omitempty" yaml:"$ref,omitempty"`
	Items      *JsonSchema_Draft07           `json:"items,omitempty" yaml:"items,omitempty"`
	X_Bi_Go    *OpenApiV310ExtentionSchema   `json:"x-bi-go,omitempty" yaml:"x-bi-go,omitempty" xml:"x-bi-go,omitempty"`
}
