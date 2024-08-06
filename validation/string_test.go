package validation_test

import (
	"boundedinfinity/codegen/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String(t *testing.T) {
	type CustomString string

	tcs := []struct {
		name    string
		input   CustomString
		fn      func(CustomString) error
		err     error
		message string
	}{
		{
			name:  "case 1",
			input: "something",
			fn:    validation.StringMinFn[CustomString]("case 1", 4),
			err:   nil,
		},
		{
			name:    "case 2",
			input:   "so",
			fn:      validation.StringMinFn[CustomString]("case 2", 4),
			err:     validation.ErrStringLessThanMin,
			message: "case 2 value so less than min value of 4",
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			err := tc.fn(tc.input)
			assert.ErrorIs(tt, err, tc.err)

			if err != nil {
				assert.Equal(tt, tc.message, err.Error())
			}
		})
	}
}
