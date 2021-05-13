package model

type TypeDescriptor struct {
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type BooleanTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     bool           `json:"example,omitempty" yaml:"example,omitempty"`
}

type StringTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     string         `json:"example,omitempty" yaml:"example,omitempty"`
}

type IntTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     int32          `json:"example,omitempty" yaml:"example,omitempty"`
}

type LongTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     int64          `json:"example,omitempty" yaml:"example,omitempty"`
}

type FloatTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     float32        `json:"example,omitempty" yaml:"example,omitempty"`
}

type DoubleTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     float64        `json:"example,omitempty" yaml:"example,omitempty"`
}

type EnumTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Symbols     []string       `json:"symbols,omitempty" yaml:"symbols,omitempty"`
	Example     string         `json:"example,omitempty" yaml:"example,omitempty"`
}

type ArrayStringTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     []string       `json:"example,omitempty" yaml:"example,omitempty"`
}

type ArrayLongTypeDescriptor struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	Example     []int32        `json:"example,omitempty" yaml:"example,omitempty"`
}

type RecordTypeDescriptor struct {
	Name        string           `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum   `json:"type,omitempty" yaml:"type,omitempty"`
	Description string           `json:"description,omitempty" yaml:"description,omitempty"`
	Fields      []TypeDescriptor `json:"fields,omitempty" yaml:"fields,omitempty"`
}
