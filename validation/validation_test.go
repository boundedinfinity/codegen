package validation_test

import (
	"boundedinfinity/codegen/validation"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_Validate(t *testing.T) {
	validator := validation.String[string]("stuff").Min(1).Max(5)

	assert.ErrorIs(t, validator.Validate("1"), nil)

	assert.ErrorIs(t, validator.Validate(""), validation.ErrStringMin)
	assert.ErrorIs(t, validator.Validate("123456"), validation.ErrStringMax)

}

func Test_Integer_Validate(t *testing.T) {
	validator := validation.Integer[int]("stuff").Min(1).Max(5)

	assert.ErrorIs(t, validator.Validate(1), nil)

	assert.ErrorIs(t, validator.Validate(0), validation.ErrIntegerMin)
	assert.ErrorIs(t, validator.Validate(6), validation.ErrIntegerMax)
}
