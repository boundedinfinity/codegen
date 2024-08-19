package validation_test

import (
	"boundedinfinity/codegen/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Integer(t *testing.T) {
	type CustomInteger int

	tcs := []struct {
		name    string
		input   CustomInteger
		fn      func(CustomInteger) error
		err     error
		message string
	}{
		{
			name:  "case 1",
			input: 5,
			fn:    validation.IntegerMinFn[CustomInteger]("case 1", 4),
			err:   nil,
		},
		{
			name:    "case 2",
			input:   3,
			fn:      validation.IntegerMinFn[CustomInteger]("case 2", 4),
			err:     validation.ErrIntegerMin,
			message: "case 2 value 3 is less than min value of 4",
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
