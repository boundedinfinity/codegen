package entity_test

import (
	"boundedinfinity/codegen/entity"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Entity_Marshal(t *testing.T) {
	tcs := []struct {
		name     string
		input    *entity.Entity_
		expected string
		err      error
	}{
		{
			name:  "case 1",
			input: entity.String().Name("bounded/type/case-01"),
			err:   nil,
			expected: `{
                "entity-type": "string",
                "name": "bounded/type/case-01"
            }`,
		},
		{
			name: "case 2",
			input: entity.String().
				Name("bounded/type/case-02").
				Required(true),
			err: nil,
			expected: `{
		        "entity-type": "string",
		        "name": "bounded/type/case-02",
		        "required": true
		    }`,
		},
		// {
		// 	name:  "case 3",
		// 	input: entity.String(),
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "string"
		//     }`,
		// },
		// {
		// 	name: "case 4",
		// 	input: entity.New().Type(entity.StringType).
		// 		Default(entity.New().Name("something")),
		// 	err: nil,
		// 	expected: `{
		//         "entity-type": "string",
		//         "default": {
		//             "name": "something"
		//         }
		//     }`,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := tc.input.ToJsonIndent()
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err, actual)
			assert.JSONEqf(tt, tc.expected, actual, actual)
		})
	}
}
