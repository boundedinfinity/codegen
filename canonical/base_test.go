package canonical_test

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_String_Marshal(t *testing.T) {
	expected := ""
	actual := ""
	assert.JSONEq(t, expected, string(actual))
}
