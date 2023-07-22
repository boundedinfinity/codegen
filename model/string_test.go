package model_test

import (
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_Marshal(t *testing.T) {
	input := model.BuildString().Min(1).Done()
	input.Common = model.BuildCommon(&input.Common).Done()
	actual := ``
	expected := ``

	assert.Equal(t, expected, actual)
}
