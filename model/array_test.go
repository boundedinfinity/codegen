package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Array(t *testing.T) {
	tcs := []struct {
		name     string
		input    model.CodeGenType
		expected string
		err      error
	}{
		// {
		// 	name:  "Serialize boolean",
		// 	input: model.BuildArray().Value(),
		// 	err:   nil,
		// 	expected: `{
		//         "type-id": "array",
		//         "value": {}
		//     }`,
		// },
		// {
		// 	name:  "Serialize boolean",
		// 	input: model.BuildArray().Name("A_NAME").Value(),
		// 	err:   nil,
		// 	expected: `{
		//         "type-id": "array",
		//         "value": {
		//             "name": "A_NAME"
		//         }
		//     }`,
		// },
		{
			name: "Serialize array with boolean",
			err:  nil,
			input: &model.CodeGenArray{
				CodeGenCommon: model.CodeGenCommon{
					Name: optioner.Some("AN_ARRAY"),
				},
				Items: &model.CodeGenBoolean{
					CodeGenCommon: model.CodeGenCommon{
						Name: optioner.Some("A_BOOLEAN"),
					},
				},
			},
			expected: `{}`,
		},
		// {
		// 	name: "Serialize array with boolean",
		// 	err:  nil,
		// 	input: &model.CodeGenArray{
		// 		CodeGenCommon: model.CodeGenCommon{
		// 			Name: optioner.Some("AN_ARRY"),
		// 		},
		// 		Items: &model.CodeGenBoolean{
		// 			Name:     optioner.Some("A_BOOL"),
		// 			Required: optioner.Some(true),
		// 		},
		// 	},
		// 	expected: `{}`,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			// bs, err := model.MarshalCodeGenObject(tc.input)
			bs, err := json.MarshalIndent(tc.input, "", "    ")
			actual := string(bs)

			if tc.err != nil {
				assert.Equal(t, tc.err, err, tc.name, actual)
			} else {
				assert.Nil(t, err, tc.name, actual)
			}

			assert.JSONEq(t, tc.expected, actual, tc.name, actual)
		})
	}
}

// func Test_Unmarshal_Array(t *testing.T) {
// 	tcs := []struct {
// 		name     string
// 		input    string
// 		expected model.CodeGenType
// 		err      error
// 	}{
// 		{
// 			name:     "Serialize boolean",
// 			input:    `{"type-id":"boolean","value":{}}`,
// 			expected: model.BuildBoolean().Value(),
// 			err:      nil,
// 		},
// 		{
// 			name:     "Serialize boolean",
// 			input:    `{"type-id":"boolean","value":{"name":"A_NAME"}}`,
// 			expected: model.BuildBoolean().Name("A_NAME").Value(),
// 			err:      nil,
// 		},
// 		{
// 			name:     "Serialize boolean",
// 			input:    `{"type-id":"boolean","value":{"name":"A_NAME", "required":true}}`,
// 			expected: model.BuildBoolean().Name("A_NAME").Required(true).Value(),
// 			err:      nil,
// 		},
// 	}

// 	for _, tc := range tcs {
// 		t.Run(tc.name, func(tt *testing.T) {
// 			actual, err := model.UnmarshalCodeGenObject([]byte(tc.input))

// 			if tc.err != nil {
// 				assert.Equal(t, tc.err, err, tc.name, tc.input)
// 			} else {
// 				assert.Nil(t, err, tc.name, tc.input)
// 			}

// 			assert.Equal(t, tc.expected, actual, tc.name, tc.input)
// 		})
// 	}
// }
