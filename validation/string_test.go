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
		fn      func(...CustomString) error
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
			err:     validation.ErrStringMin,
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

func Test_String_Error(t *testing.T) {
	err := validation.StringNotEmtpy("something", "a", "b", "", "d", "e")
	assert.ErrorIs(t, err, validation.ErrStringEmpty)

	details, ok := err.(*validation.ErrStringEmtpyDetails)
	assert.True(t, ok)
	assert.Equal(t, 2, details.Index)
	assert.Equal(t, 5, details.Length)
	assert.Equal(t, "something[2] string is empty", details.Error())

}
