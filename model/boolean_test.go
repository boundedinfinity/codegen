package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	o "github.com/boundedinfinity/go-commoner/functional/optioner"
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
			input: &model.CodeGenBoolean{},
			err:   nil,
			expected: ` {
                "type": "boolean"
            }`,
		},
		{
			name: "Marshal boolean with name",
			input: &model.CodeGenBoolean{
				CodeGenCommon: model.CodeGenCommon{
					Name: o.Some("A_BOOLEAN"),
				},
			},
			err: nil,
			expected: `{
		        "type": "boolean",
		        "name": "A_BOOLEAN"
		    }`,
		},
		{
			name: "Marshal boolean with name and required",
			input: &model.CodeGenBoolean{
				CodeGenCommon: model.CodeGenCommon{
					Name:     o.Some("A_BOOLEAN"),
					Required: o.Some(true),
				},
			},
			err: nil,
			expected: `{
		        "type": "boolean",
		        "name": "A_BOOLEAN",
		        "required": true
		    }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(tc.input, "", "    ")
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err, string(actual))
			assert.JSONEqf(tt, tc.expected, actual, actual)
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
			obj:  &model.CodeGenBoolean{},
			err:  nil,
		},
		{
			name: "Marshal boolean",
			obj: &model.CodeGenBoolean{
				CodeGenCommon: model.CodeGenCommon{
					Name: o.Some("A_BOOLEAN"),
				},
			},
			err: nil,
		},
		{
			name: "Marshal boolean",
			obj: &model.CodeGenBoolean{
				CodeGenCommon: model.CodeGenCommon{
					Name:     o.Some("A_BOOLEAN"),
					Required: o.Some(true),
				},
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			data, err := json.MarshalIndent(tc.obj, "", "    ")
			assert.ErrorIs(tt, err, tc.err)

			var actual model.CodeGenBoolean

			err = json.Unmarshal(data, &actual)

			assert.ErrorIs(tt, err, tc.err)
			assert.EqualValuesf(tt, tc.obj, &actual, string(data))
		})
	}
}
