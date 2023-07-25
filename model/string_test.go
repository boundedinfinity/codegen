package model_test

import (
	"boundedinfinity/codegen/model"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_Marshal(t *testing.T) {
	input := model.BuildString().
		Min(1).
		Common(model.BuildCommon().
			Desc("Test Description").
			Done()).
		Done()

	actual, err := json.Marshal(input)
	expected := ``

	output := string(actual)

	assert.Nil(t, err)
	assert.Equal(t, expected, output)
}
