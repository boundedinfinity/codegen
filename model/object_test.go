package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Object(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenObject
		expected string
		err      error
	}{
		{
			name:  "Serialize boolean",
			input: model.NewObject(),
			err:   nil,
			expected: `{
		        "base-type": "object",
                "name": null,
                "description": null,
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "properties": []
		    }`,
		},
		{
			name: "Serialize boolean",
			input: model.NewObject().
				WithName("AN_OBJECT").
				WithDescription("an object description"),
			err: nil,
			expected: `{
		        "base-type": "object",
                "name": "AN_OBJECT",
                "description": "an object description",
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "properties": []
		    }`,
		},
		{
			name: "Serialize object with boolean",
			err:  nil,
			input: model.NewObject().
				WithName("AN_OBJECT").
				WithDescription("an object description").
				WithProperties(model.NewBoolean().
					WithName("A_BOOLEAN").
					WithDescription("a bool description")),
			expected: `{
		        "base-type": "object",
                "name": "AN_OBJECT",
                "description": "an object description",
                "required": null,
                "default": null,
                "inherit": null,
                "links": null,
                "properties": [
                    {
                        "base-type": "boolean",
                        "name": "A_BOOLEAN",
                        "description": "a bool description",
                        "required": null,
                        "default": null,
                        "inherit": null,
                        "links": null
                    }
                ]
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

func Test_Unmarshal_Object(t *testing.T) {
	tcs := []struct {
		name string
		obj  *model.CodeGenObject
		err  error
	}{
		{
			name: "Unmarshal object and boolean item",
			obj: model.NewObject().
				WithName("AN_OBJECT").
				WithDescription("an object description").
				WithProperties(
					model.NewBoolean().WithName("A_BOOLEAN").WithDescription("a bool description"),
					// model.NewInteger().WithName("A_INT").WithDescription("a int description"),
				),
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			input, err := json.Marshal(&tc.obj)
			assert.Nilf(t, err, tc.name, "%v : %v", tc.name, string(input))

			var actual model.CodeGenObject

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
