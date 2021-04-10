package model

import (
	"errors"

	"github.com/boundedinfinity/optional"
)

type RunContext struct {
	Input optional.StringOptional
	Model OpenApiV310
}

func (t RunContext) Validate() error {
	if t.Input.IsEmpty() {
		return errors.New("modelPath is empty")
	}

	return nil
}

type XBiGoGlobalContext struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Model   OpenApiV310             `json:"model,omitempty" yaml:"model,omitempty" xml:"model,omitempty"`
}

type XBiGoGlobalRuntime struct {
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
	Context XBiGoGlobalContext      `json:"context,omitempty" yaml:"context,omitempty" xml:"context,omitempty"`
}

type XBiGoSchemaContext struct {
	Package optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Name    optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	Schema  JsonSchema_Draft07      `json:"schema,omitempty" yaml:"schema,omitempty"`
	Model   OpenApiV310             `json:"model,omitempty" yaml:"model,omitempty" xml:"model,omitempty"`
}

type XBiGoSchemaRuntime struct {
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
	Context XBiGoSchemaContext      `json:"context,omitempty" yaml:"context,omitempty" xml:"context,omitempty"`
}

type XBiGoPathItemContext struct {
	Package  optional.StringOptional `json:"package,omitempty" yaml:"package,omitempty" xml:"package,omitempty"`
	Name     optional.StringOptional `json:"name,omitempty" yaml:"name,omitempty" xml:"name,omitempty"`
	PathItem OpenApiV310PathItem     `json:"pathItem,omitempty" yaml:"pathItem,omitempty" xml:"pathItem,omitempty"`
	Model    OpenApiV310             `json:"model,omitempty" yaml:"model,omitempty" xml:"model,omitempty"`
}

type XBiGoPathItemRuntime struct {
	Input   optional.StringOptional `json:"input,omitempty" yaml:"input,omitempty" xml:"input,omitempty"`
	Output  optional.StringOptional `json:"output,omitempty" yaml:"output,omitempty" xml:"output,omitempty"`
	Context XBiGoPathItemContext    `json:"context,omitempty" yaml:"context,omitempty" xml:"context,omitempty"`
}
