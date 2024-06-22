package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	tcs := []struct {
		name     string
		lang     string
		inputs   *model.CodeGenProject
		expected map[string]string
		err      error
		newErr   error
	}{
		{
			name:     "process files",
			inputs:   nil,
			expected: map[string]string{},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			gen, err := generator.New(tc.lang)
			assert.ErrorIs(tt, err, tc.newErr)

			actual, err := gen.Generate(tc.inputs)
			assert.ErrorIs(tt, err, tc.err)
			assert.ElementsMatch(tt, tc.expected, actual)
		})
	}
}
