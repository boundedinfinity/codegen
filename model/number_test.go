package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Marshal_Integer(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenInteger
		expected string
		err      error
	}{
		// {
		// 	name:  "case 1",
		// 	input: model.BuildInteger().Build(),
		// 	err:   nil,
		// 	expected: `{
		//         "type": "integer"
		//     }`,
		// },
		// {
		// 	name: "case 2",
		// 	input: model.BuildInteger().
		// 		Name("AN_INTEGER").
		// 		Description("an integer description").
		// 		Build(),
		// 	err: nil,
		// 	expected: `{
		//         "type": "integer",
		//         "name": "AN_INTEGER",
		//         "description": "an integer description"
		//     }`,
		// },
		// {
		// 	name: "case 3",
		// 	input: model.BuildInteger().
		// 		Name("AN_INTEGER").
		// 		Description("an integer description").
		// 		MultipleOf(5).
		// 		Build(),
		// 	err: nil,
		// 	expected: `{
		//         "type": "integer",
		//         "name": "AN_INTEGER",
		//         "description": "an integer description",
		//         "multiple-of": 5
		//     }`,
		// },
		// {
		// 	name: "case 4",
		// 	input: model.BuildInteger().
		// 		Name("AN_INTEGER").
		// 		Description("an integer description").
		// 		MultipleOf(5).
		// 		Ranges(
		// 			model.BuildRange[int]().ExclusiveMax(10).ExclusiveMin(1).Build(),
		// 		).Build(),
		// 	err: nil,
		// 	expected: `{
		//         "type": "integer",
		//         "description": "an integer description",
		//         "multiple-of": 5,
		//         "name": "AN_INTEGER",
		//         "ranges": [
		//             {
		//                 "exclusive-max": 10,
		//                 "exclusive-min": 1
		//             }
		//         ]
		//     }`,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := json.MarshalIndent(&tc.input, "", "    ")
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err)
			assert.JSONEqf(tt, tc.expected, actual, "%v = %v", actual)
		})
	}
}

func Test_Marshal_Validate(t *testing.T) {
	tcs := []struct {
		name     string
		input    *model.CodeGenInteger
		expected error
	}{
		// {
		// 	name:     "case 01",
		// 	input:    model.BuildInteger().Build(),
		// 	expected: nil,
		// },

		// {
		// 	name:     "case 02",
		// 	input:    model.BuildInteger().MultipleOf(5).Build(),
		// 	expected: nil,
		// },
		// {
		// 	name:     "case 03",
		// 	input:    model.BuildInteger().MultipleOf(0).Build(),
		// 	expected: model.ErrNumberMultipleOfBelow1,
		// },
		// {
		// 	name:     "case 04",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Max(4).Build()).Build(),
		// 	expected: model.ErrNumberRangeMinOrExclusiveMinRequired,
		// },
		// {
		// 	name:     "case 05",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(4).Build()).Build(),
		// 	expected: model.ErrNumberRangeMaxOrExclusiveMaxRequired,
		// },
		// {
		// 	name:     "case 06",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(1).ExclusiveMin(1).Build()).Build(),
		// 	expected: model.ErrNumberRangeMinAndExclusiveMinMutuallyExclusive,
		// },
		// {
		// 	name:     "case 07",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(1).Max(10).ExclusiveMax(10).Build()).Build(),
		// 	expected: model.ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive,
		// },
		// {
		// 	name:     "case 08",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(1).Max(1).Build()).Build(),
		// 	expected: nil,
		// },
		// {
		// 	name:     "case 09",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(2).Max(1).Build()).Build(),
		// 	expected: model.ErrNumberRangeMaxLessThanMin,
		// },
		// {
		// 	name:     "case 10",
		// 	input:    model.BuildInteger().Ranges(model.BuildRange[int]().Min(1).ExclusiveMax(1).Build()).Build(),
		// 	expected: model.ErrNumberRangeMaxLessThanMin,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Validate()
			assert.ErrorIs(tt, actual, tc.expected)
		})
	}
}
