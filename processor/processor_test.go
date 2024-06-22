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
			err := processor.New().ProcessFiles(tc.inputs...)

			if tc.err != nil {
				assert.Equal(tt, tc.err, err, tc.name)
			} else {
				assert.Nil(tt, err, tc.name)
			}
		})
	}
}
