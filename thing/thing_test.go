package thing_test

import (
	"boundedinfinity/codegen/thing"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_XX(t *testing.T) {
	x := thing.BuildInt().Max(42).Name("int-name").Build()
	assert.Equal(t, x.Name, "int-name")
	assert.Equal(t, x.Max, 42)
}
