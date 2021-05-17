package model

import (
	"fmt"
	"strings"

	"gopkg.in/yaml.v2"
)

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
	Models []InputModel `json:"models,omitempty" yaml:"models,omitempty"`
	// Operations []InputOperation          `json:"operations,omitempty" yaml:"operations,omitempty"`
	// Templates  []InputTemplate           `json:"templates,omitempty" yaml:"templates,omitempty"`
}

type InputModel struct {
	Name        string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type        SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Array       bool           `json:"array,omitempty" yaml:"array,omitempty"`
	Description string         `json:"description,omitempty" yaml:"description,omitempty"`
	String      StringInputModel
	StringArray StringArrayInputModel
	Int         IntInputModel
	IntArray    IntArrayInputModel
	Long        LongInputModel
	LongArray   LongArrayInputModel
	Bool        BoolInputModel
	BoolArray   BoolArrayInputModel
	Float       FloatInputModel
	FloatArray  FloatArrayInputModel
	Double      DoubleInputModel
	DoubleArray DoubleArrayInputModel
	Complex     ComplexInputModel
	Enum        EnumInputModel
	Ref         RefInputModel
}

func (t *InputModel) MarshalYAML() (interface{}, error) {
	switch t.Type {
	case SchemaType_Boolean:
		return yaml.Marshal(t.Bool)
	default:
		return struct{}{}, fmt.Errorf("invalid type '%v'", t.Type)
	}
}

func (t *InputModel) UnmarshalYAML(unmarshal func(interface{}) error) error {
	var d inputModelDiscriminator

	if err := unmarshal(&d); err != nil {
		return err
	}

	isArray := strings.HasSuffix(d.Type, COLLECTION_SUFFIX)
	singularType := strings.TrimSuffix(d.Type, COLLECTION_SUFFIX)

	t.Name = d.Name
	t.Array = isArray

	switch singularType {
	case SchemaType_Boolean.String():
		t.Type = SchemaType_Boolean
		if isArray {
			if err := unmarshal(&t.BoolArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.Bool); err != nil {
				return err
			}
		}
	case SchemaType_String.String():
		t.Type = SchemaType_String
		if isArray {
			if err := unmarshal(&t.StringArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.String); err != nil {
				return err
			}
		}
	case SchemaType_Int.String():
		t.Type = SchemaType_Int
		if isArray {
			if err := unmarshal(&t.IntArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.Int); err != nil {
				return err
			}
		}
	case SchemaType_Long.String():
		t.Type = SchemaType_Long
		if isArray {
			if err := unmarshal(&t.LongArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.Long); err != nil {
				return err
			}
		}
	case SchemaType_Float.String():
		t.Type = SchemaType_Double
		if isArray {
			if err := unmarshal(&t.FloatArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.Float); err != nil {
				return err
			}
		}
	case SchemaType_Double.String():
		t.Type = SchemaType_Double
		if isArray {
			if err := unmarshal(&t.DoubleArray); err != nil {
				return err
			}
		} else {
			if err := unmarshal(&t.Double); err != nil {
				return err
			}
		}
	case SchemaType_Complex.String():
		t.Type = SchemaType_Complex
		if err := unmarshal(&t.Complex); err != nil {
			return err
		}
	case SchemaType_Enum.String():
		t.Type = SchemaType_Enum
		if err := unmarshal(&t.Enum); err != nil {
			return err
		}
	case SchemaType_Ref.String():
		t.Type = SchemaType_Ref
		if err := unmarshal(&t.Ref); err != nil {
			return err
		}
	default:
		return fmt.Errorf("invalid type '%v' (name: %v)", d.Type, d.Name)
	}

	return nil
}

type inputModelDiscriminator struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	Type string `json:"type,omitempty" yaml:"type,omitempty"`
}

type BoolInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example bool           `json:"example,omitempty" yaml:"example,omitempty"`
}

type BoolArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []bool         `json:"example,omitempty" yaml:"example,omitempty"`
}

type StringInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example string         `json:"example,omitempty" yaml:"example,omitempty"`
}

type StringArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []string       `json:"example,omitempty" yaml:"example,omitempty"`
}

type IntInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example int32          `json:"example,omitempty" yaml:"example,omitempty"`
}

type IntArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []int32        `json:"example,omitempty" yaml:"example,omitempty"`
}

type LongInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example int64          `json:"example,omitempty" yaml:"example,omitempty"`
}

type LongArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []int64        `json:"example,omitempty" yaml:"example,omitempty"`
}

type FloatInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example float32        `json:"example,omitempty" yaml:"example,omitempty"`
}

type FloatArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []float32      `json:"example,omitempty" yaml:"example,omitempty"`
}

type DoubleInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example float64        `json:"example,omitempty" yaml:"example,omitempty"`
}

type DoubleArrayInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Example []float64      `json:"example,omitempty" yaml:"example,omitempty"`
}

type ComplexInputModel struct {
	Name       string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type       SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Properties []InputModel   `json:"properties,omitempty" yaml:"properties,omitempty"`
}

type EnumInputModel struct {
	Name    string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type    SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Symbols []string       `json:"symbols,omitempty" yaml:"symbols,omitempty"`
	Example string         `json:"example,omitempty" yaml:"example,omitempty"`
}

type RefInputModel struct {
	Name string         `json:"name,omitempty" yaml:"name,omitempty"`
	Type SchemaTypeEnum `json:"type,omitempty" yaml:"type,omitempty"`
	Ref  string         `json:"ref,omitempty" yaml:"ref,omitempty"`
}

type InputValidation struct {
	Minimum  int  `json:"minimum,omitempty" yaml:"minimum,omitempty"`
	Maximum  int  `json:"maximum,omitempty" yaml:"maximum,omitempty"`
	Required bool `json:"required,omitempty" yaml:"required,omitempty"`
}

type InputOperation struct {
	Name string `json:"name,omitempty" yaml:"name,omitempty"`
	// Input       InputModel `json:"input,omitempty" yaml:"input,omitempty"`
	// Output      InputModel `json:"output,omitempty" yaml:"output,omitempty"`
	Description string `json:"description,omitempty" yaml:"description,omitempty"`
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
	Header string `json:"header,omitempty" yaml:"header,omitempty"`
	Path   string `json:"path,omitempty" yaml:"path,omitempty"`
	Type   string `json:"type,omitempty" yaml:"type,omitempty"`
}
