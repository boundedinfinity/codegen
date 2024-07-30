package processor_test

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/processor"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Process_Files(t *testing.T) {
	tcs := []struct {
		name     string
		inputs   []string
		expected []*model.CodeGenProject
		err      error
	}{
		{
			name:   "process files",
			inputs: []string{"./internal/test/data/label.yaml"},
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			proc := processor.New()

			err := proc.ProcessFiles(tc.inputs...)
			assert.ErrorIs(tt, err, tc.err)

			fullName, fullNameOk := proc.TypeMap["codegen/schema/person/full-name"]
			assert.True(t, fullNameOk)
			assert.Equal(t, "codegen/schema/person/full-name", fullName.Common().Id.Get())
			assert.Equal(t, "FullName", fullName.Common().Name.Get())
		})
	}
}
