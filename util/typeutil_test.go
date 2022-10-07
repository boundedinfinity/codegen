package util_test

import (
	"boundedinfinity/codegen/util"
	"testing"

	"github.com/boundedinfinity/jsonschema/model"
	"github.com/stretchr/testify/assert"
)

var (
	input1 = model.JsonSchemaString[string]{
		JsonSchemaCommon: model.JsonSchemaCommon{
			// Id: optioner.Some("https://www.boundedinfinity.com/schema/banking/account-type"),
		},
	}
)

func Test_GoTypeBase(t *testing.T) {
	input := input1
	expected := "AccountType"
	actual := util.GoTypeBase(input.JsonSchemaCommon)

	assert.Equal(t, expected, actual)
}

func Test_GoTypeDir(t *testing.T) {
	input := input1
	expected := "https://www.boundedinfinity.com/schema/banking"
	actual := util.GoTypeDir(input.JsonSchemaCommon)

	assert.Equal(t, expected, actual)
}
