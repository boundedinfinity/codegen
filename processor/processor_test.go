package processor_test

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/processor"
	"testing"

	o "github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/davecgh/go-spew/spew"
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

			assertString(tt, proc, &model.CodeGenString{
				CodeGenCommon: model.CodeGenCommon{
					Id:         o.Some("codegen/schema/util/id"),
					Name:       o.Some("Id"),
					Package:    o.Some("util"),
					ImportPath: o.Some("boundedinfinity-test/codegen/schema/util"),
				},
			})

			assertObject(tt, proc,
				&model.CodeGenObject{
					CodeGenCommon: model.CodeGenCommon{
						Id:         o.Some("codegen/schema/util/label"),
						Name:       o.Some("Label"),
						Package:    o.Some("util"),
						ImportPath: o.Some("boundedinfinity-test/codegen/schema/util"),
						YamlName:   o.Some("label"),
						JsonName:   o.Some("label"),
						SqlName:    o.Some("label"),
					},
					Properties: o.Some([]model.CodeGenType{
						// &model.CodeGenRef{
						// 	Ref: o.Some("codegen/schema/util/id"),
						// 	Resolved: &model.CodeGenString{
						// 		CodeGenCommon: model.CodeGenCommon{
						// 			Id:         o.Some("codegen/schema/util/id"),
						// 			Name:       o.Some("Id"),
						// 			Package:    o.Some("util"),
						// 			ImportPath: o.Some("boundedinfinity-test/codegen/schema/util"),
						// 		},
						// 	},
						// },
						&model.CodeGenString{
							CodeGenCommon: model.CodeGenCommon{
								Name: o.Some("name"),
							},
							Min: o.Some(1),
							Max: o.Some(10),
						},
					}),
				},
			)
		})
	}
}

func assertString(tt *testing.T, proc *processor.Processor, expected *model.CodeGenString) {
	assert.True(tt, expected.Id.Defined())
	typ, ok1 := proc.TypeMap[expected.Id.Get()]
	assert.True(tt, ok1)
	assert.Equal(tt, expected.GetType(), typ.GetType())
	actual, ok := typ.(*model.CodeGenString)
	assert.True(tt, ok)
	assert.Equal(tt, expected, actual)
}

func assertObject(tt *testing.T, proc *processor.Processor, expected *model.CodeGenObject) {
	assert.True(tt, expected.Id.Defined())
	typ, ok1 := proc.TypeMap[expected.Id.Get()]
	assert.True(tt, ok1)
	assert.Equal(tt, expected.GetType(), typ.GetType())
	actual, ok := typ.(*model.CodeGenObject)
	assert.True(tt, ok)
	assert.Equal(tt, expected, actual, spew.Sdump(actual))
}
