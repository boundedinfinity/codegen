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
                "properties": [
                    {
                        "base-type": "boolean",
                        "name": "A_BOOLEAN",
                        "description": "a bool description"
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

			assert.ErrorIs(tt, err, tc.err)
			assert.JSONEqf(tt, tc.expected, actual, actual)
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
			assert.ErrorIs(tt, err, tc.err, string(input))

			var actual model.CodeGenObject

			err = json.Unmarshal(input, &actual)

			assert.ErrorIs(tt, err, tc.err, string(input))
			assert.Equalf(tt, tc.obj, &actual, string(input))
		})
	}
}
