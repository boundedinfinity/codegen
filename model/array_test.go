package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Array(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenArray
		expected string
		err      error
	}{
		{
			name:  "Serialize boolean",
			input: model.NewArray(),
			err:   nil,
			expected: `{
		        "base-type": "array",
                "name": null,
                "description": null,
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "min": null,
                "max": null
		    }`,
		},
		{
			name: "Serialize boolean",
			input: model.NewArray().
				WithName("AN_ARRAY").
				WithDescription("an array description"),
			err: nil,
			expected: `{
		        "base-type": "array",
                "name": "AN_ARRAY",
                "description": "an array description",
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "min": null,
                "max": null
		    }`,
		},
		{
			name: "Serialize array with boolean",
			err:  nil,
			input: model.NewArray().
				WithName("AN_ARRAY").
				WithDescription("an array description").
				WithItems(model.NewBoolean().
					WithName("A_BOOLEAN").
					WithDescription("a bool description")),
			expected: `{
		        "base-type": "array",
                "name": "AN_ARRAY",
                "description": "an array description",
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "items": {
                    "base-type": "boolean",
                    "name": "A_BOOLEAN",
                    "description": "a bool description",
                    "required": null,
                    "default": null,
                    "inherit": null,
                    "links": null
                },
                "min": null,
                "max": null
		    }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// bs, err := model.MarshalCodeGenObject(tc.input)
			bs, err := json.MarshalIndent(&tc.input, "", "    ")
			actual := string(bs)

			if tc.err != nil {
				assert.Equal(t, tc.err, err, tc.name, actual)
			} else {
				assert.Nil(t, err, tc.name, actual)
			}

			assert.JSONEqf(t, tc.expected, actual, "%v = %v", tc.name, actual)
		})
	}
}

func Test_Unmarshal_Array(t *testing.T) {
	tcs := []struct {
		name string
		obj  *model.CodeGenArray
		err  error
	}{
		{
			name: "Unmarshal array and boolean item",
			obj: model.NewArray().
				WithName("AN_ARRAY").
				WithDescription("an array description").
				WithItems(model.NewBoolean().
					WithName("A_BOOLEAN").
					WithDescription("a bool description")),
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			input, err := json.Marshal(&tc.obj)
			assert.Nilf(t, err, tc.name, "%v : %v", tc.name, string(input))

			var actual model.CodeGenArray

			err = json.Unmarshal(input, &actual)

			if tc.err != nil {
				assert.Equalf(t, tc.err, err, "%v : %v", tc.name, string(input))
			} else {
				assert.Nilf(t, err, tc.name, "%v : %v", tc.name, string(input))
			}

			assert.Equalf(t, tc.obj, &actual, "%v : %v", tc.name, string(input))
		})
	}
}
