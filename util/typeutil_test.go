package util_test

import (
	"boundedinfinity/codegen/util"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/stretchr/testify/assert"
)

var (
	input1 = model.JsonSchemaString{
		Id: o.Some(model.IdT("https://www.boundedinfinity.com/schema/banking/account-type")),
	}
)

func Test_GoTypeBase(t *testing.T) {
	input := input1
	expected := "AccountType"
	actual := util.GoTypeBase(input)

	assert.Equal(t, expected, actual)
}

func Test_GoTypeDir(t *testing.T) {
	input := input1
	expected := "https://www.boundedinfinity.com/schema/banking"
	actual := util.GoTypeDir(input)

	assert.Equal(t, expected, actual)
}
