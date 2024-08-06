package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	o "github.com/boundedinfinity/go-commoner/functional/optioner"
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
			input: &model.CodeGenArray{},
			err:   nil,
			expected: `{
		        "type": "array"
		    }`,
		},
		{
			name: "Serialize boolean",
			input: &model.CodeGenArray{
				CodeGenCommon: model.CodeGenCommon{
					Name:        o.Some("AN_ARRAY"),
					Description: o.Some("an array description"),
				},
			},
			err: nil,
			expected: `{
		        "type": "array",
		        "name": "AN_ARRAY",
		        "description": "an array description"
		    }`,
		},
		{
			name: "Serialize array with boolean",
			err:  nil,
			input: &model.CodeGenArray{
				CodeGenCommon: model.CodeGenCommon{
					Name:        o.Some("AN_ARRAY"),
					Description: o.Some("an array description"),
				},
				Items: o.Some(model.CodeGenSchema(&model.CodeGenBoolean{
					CodeGenCommon: model.CodeGenCommon{
						Name:        o.Some("A_BOOLEAN"),
						Description: o.Some("a bool description"),
					},
				})),
			},
			expected: `{
		        "type": "array",
		        "name": "AN_ARRAY",
		        "description": "an array description",
		        "items": {
		            "type": "boolean",
		            "name": "A_BOOLEAN",
		            "description": "a bool description"
		        }
		    }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(&tc.input, "", "    ")
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err)
			assert.JSONEqf(tt, tc.expected, actual, actual)
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
			obj: &model.CodeGenArray{
				CodeGenCommon: model.CodeGenCommon{
					Name:        o.Some("AN_ARRAY"),
					Description: o.Some("an array description"),
				},
				Items: o.Some(model.CodeGenSchema(&model.CodeGenBoolean{
					CodeGenCommon: model.CodeGenCommon{
						Name:        o.Some("A_BOOLEAN"),
						Description: o.Some("a bool description"),
					},
				})),
			},
			err: nil,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			input, err := json.Marshal(&tc.obj)
			assert.ErrorIs(tt, err, tc.err, string(input))

			var actual model.CodeGenArray

			err = json.Unmarshal(input, &actual)

			assert.ErrorIs(tt, err, tc.err, string(input))
			assert.Equalf(tt, tc.obj, &actual, string(input))
		})
	}
}
