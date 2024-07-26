package model_test

import (
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Unmarshal_CodeGenType(t *testing.T) {
	tcs := []struct {
		name     string
		input    string
		expected model.CodeGenType
		err      error
	}{
		{
			name: "Unmarshal boolean",
			input: `{
		        "base-type": "boolean",
		        "value":{
                    "name": null,
                    "description": null,
                    "required": null,
                    "default": null,
                    "inherit": null,
                    "links": null
                }
		    }`,
			err:      nil,
			expected: model.BuildBoolean().Build(),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := model.UnmarshalCodeGenType([]byte(tc.input))

			assert.ErrorIs(tt, err, tc.err)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
