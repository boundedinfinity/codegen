package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Boolean(t *testing.T) {
	tcs := []struct {
		name     string
		input    model.CodeGenType
		expected string
		err      error
	}{
		{
			name:  "Marshal boolean",
			input: model.NewBoolean(),
			err:   nil,
			expected: `{
		        "base-type": "boolean",
                "name": null,
                "description": null,
                "required": null,
                "default": null,
                "inherit": null,
                "links": null
		    }`,
		},
		{
			name:  "Marshal boolean with name",
			input: model.NewBoolean().WithName("A_BOOLEAN"),
			err:   nil,
			expected: `{
		        "base-type": "boolean",
                "name": "A_BOOLEAN",
                "description": null,
                "required": null,
                "default": null,
                "inherit": null,
                "links": null
		    }`,
		},
		{
			name:  "Marshal boolean with name and required",
			input: model.NewBoolean().WithName("A_BOOLEAN").WithRequired(true),
			err:   nil,
			expected: `{
		        "base-type": "boolean",
                "name": "A_BOOLEAN",
                "description": null,
                "required": true,
                "default": null,
                "inherit": null,
                "links": null
		    }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(tc.input, "", "    ")
			actual := string(bs)

			if tc.err != nil {
				assert.Equalf(t, tc.err, err, "%v : %v", tc.name, actual)
			} else {
				assert.Nil(t, err, tc.name, actual)
			}

			assert.JSONEqf(t, tc.expected, actual, "%v : %v", tc.name, actual)
		})
	}
}

func Test_Unmarshal_Boolean(t *testing.T) {
	tcs := []struct {
		name string
		obj  *model.CodeGenBoolean
		err  error
	}{
		{
			name: "Marshal boolean",
			obj:  model.NewBoolean(),
			err:  nil,
		},
		{
			name: "Marshal boolean",
			obj:  model.NewBoolean().WithName("A_BOOLEAN"),
			err:  nil,
		},
		{
			name: "Marshal boolean",
			obj:  model.NewBoolean().WithName("A_BOOLEAN").WithRequired(true),
			err:  nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			data, err := json.MarshalIndent(tc.obj, "", "    ")
			assert.Nil(t, err, tc.name, string(data))

			var actual model.CodeGenBoolean

			err = json.Unmarshal(data, &actual)

			if tc.err != nil {
				assert.Equal(t, tc.err, err, tc.name, string(data))
			} else {
				assert.Nil(t, err, tc.name, string(data))
			}

			assert.EqualValuesf(t, tc.obj, &actual, "%v : %v", tc.name, string(data))
		})
	}
}
