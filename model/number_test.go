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
		{
			name:  "case 1",
			input: model.NewInteger(),
			err:   nil,
			expected: `{
		        "base-type": "integer"
		    }`,
		},
		{
			name: "case 2",
			input: model.NewInteger().
				WithName("AN_INTEGER").
				WithDescription("an integer description"),
			err: nil,
			expected: `{
		        "base-type": "integer",
		        "name": "AN_INTEGER",
		        "description": "an integer description"
		    }`,
		},
		{
			name: "case 3",
			input: model.NewInteger().
				WithName("AN_INTEGER").
				WithDescription("an integer description").
				WithMultipleOf(5),
			err: nil,
			expected: `{
		        "base-type": "integer",
		        "name": "AN_INTEGER",
		        "description": "an integer description",
                "multiple-of": 5
		    }`,
		},
		{
			name: "case 4",
			input: model.NewInteger().
				WithName("AN_INTEGER").
				WithDescription("an integer description").
				WithMultipleOf(5).
				WithRange(model.NewRange[int]().
					WithExclusiveMax(10).WithExclusiveMin(1),
				),
			err: nil,
			expected: `{
                "base-type": "integer",
                "description": "an integer description",
                "multiple-of": 5,
                "name": "AN_INTEGER",
                "ranges": [
                    {
                        "exclusive-max": 10,
                        "exclusive-min": 1
                    }
                ]
            }`,
		},
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
		{
			name:     "case 01",
			input:    model.NewInteger(),
			expected: nil,
		},

		{
			name:     "case 02",
			input:    model.NewInteger().WithMultipleOf(5),
			expected: nil,
		},
		{
			name:     "case 03",
			input:    model.NewInteger().WithMultipleOf(0),
			expected: model.ErrNumberMultipleOfBelow1,
		},
		{
			name:     "case 04",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMax(4)),
			expected: model.ErrNumberRangeMinOrExclusiveMinRequired,
		},
		{
			name:     "case 05",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(4)),
			expected: model.ErrNumberRangeMaxOrExclusiveMaxRequired,
		},
		{
			name:     "case 06",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(1).WithExclusiveMin(1)),
			expected: model.ErrNumberRangeMinAndExclusiveMinMutuallyExclusive,
		},
		{
			name:     "case 07",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(1).WithMax(10).WithExclusiveMax(10)),
			expected: model.ErrNumberRangeMaxAndExclusiveMaxMutuallyExclusive,
		},
		{
			name:     "case 08",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(1).WithMax(1)),
			expected: nil,
		},
		{
			name:     "case 09",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(2).WithMax(1)),
			expected: model.ErrNumberRangeMaxLessThanMin,
		},
		{
			name:     "case 10",
			input:    model.NewInteger().WithRange(model.NewRange[int]().WithMin(1).WithExclusiveMax(1)),
			expected: model.ErrNumberRangeMaxLessThanMin,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Validate()
			assert.ErrorIs(tt, actual, tc.expected)
		})
	}
}
