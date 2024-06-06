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
		        "type-id": "boolean",
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
			expected: model.NewBoolean(),
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := model.UnmarshalCodeGenType([]byte(tc.input))

			if tc.err != nil {
				assert.Equalf(t, tc.err, err, "%v : %v", tc.name, actual)
			} else {
				assert.Nil(t, err, tc.name, actual)
			}

			assert.Equalf(t, tc.expected, actual, "%v : %v", tc.name, actual)
		})
	}
}
