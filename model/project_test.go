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
			input: model.NewProject().
				WithName("A_PROJECT").
				WithOperations(
					model.NewOperation().WithInputs(
						model.NewRef().WithRef("label").WithName("A_LABEL"),
					),
				).
				WithTypes(
					model.NewObject().
						WithName("label").
						WithDescription("A simple label").
						WithProperties(
							model.NewString().
								WithName("name").
								WithDescription("The label name").
								WithMin(0).
								WithMax(50),
						),
				),
			err: nil,
			expected: `{
		        "codegen-id": "boolean",
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

			if tc.err != nil {
				assert.Equalf(t, tc.err, err, "%v : %v", tc.name, actual)
			} else {
				assert.Nil(t, err, tc.name, actual)
			}

			assert.JSONEqf(t, tc.expected, actual, "%v : %v", tc.name, actual)
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
