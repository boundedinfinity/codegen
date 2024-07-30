package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Project(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenProject
		expected string
		err      error
	}{
		{
			name: "Marshal project",
			input: model.BuildProject().
				Name("A_PROJECT").
				Operations(
					model.BuildOperation().Inputs(
						model.BuildRef().Ref("label").Name("A_LABEL").Build(),
					).Build(),
				).
				Types(
					model.BuildObject().
						Id("label").
						Description("A simple label").
						Properties(
							model.BuildString().
								Name("name").
								Description("The label name").
								Min(0).
								Max(50).
								Build(),
						).Build(),
				).Build(),
			err: nil,
			expected: `{
		        "type": "boolean",
                "name": null,
                "description": null,
                "required": null,
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

			assert.ErrorIs(tt, err, tc.err)
			assert.JSONEqf(tt, tc.expected, actual, actual)
		})
	}
}

func Test_Unmarshal_Project(t *testing.T) {
	tcs := []struct {
		name string
		obj  *model.CodeGenBoolean
		err  error
	}{
		{
			name: "Marshal boolean",
			obj:  model.BuildBoolean().Name("A_BOOLEAN").Required(true).Build(),
			err:  nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			data, err := json.MarshalIndent(tc.obj, "", "    ")
			assert.Nil(tt, err, string(data))

			var actual model.CodeGenBoolean

			err = json.Unmarshal(data, &actual)

			assert.ErrorIs(tt, err, tc.err)
			assert.EqualValuesf(tt, tc.obj, &actual, "%v : %v", string(data))
		})
	}
}
