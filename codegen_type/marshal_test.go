package codegen_type_test

import (
	c "boundedinfinity/codegen/codegen_type"
	"strings"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

var (
	tests = []struct {
		name     string
		input    string
		expected c.CodeGenType
	}{
		{
			"ref1",
			`
ref: https://www.boundedinfinity.com/schema/something
description: ref1
`,
			&c.CodeGenTypeRef{
				CodeGenTypeBase: c.CodeGenTypeBase{
					Description: o.Some("ref1"),
				},
				Ref: o.Some("https://www.boundedinfinity.com/schema/something"),
			},
		},
		{
			"string1",
			`
type: string
description: string1
min: 1
max: 2
`,
			&c.CodeGenTypeString{
				CodeGenTypeBase: c.CodeGenTypeBase{
					Description: o.Some("string1"),
				},
				Min: o.Some(1),
				Max: o.Some(2),
			},
		},
		{
			"array1",
			`
type: array
description: array1
min: 3
max: 4
items: 
    type: string
    description: string1
    min: 1
    max: 2
`,
			&c.CodeGenTypeArray{
				CodeGenTypeBase: &c.CodeGenTypeBase{
					Description: o.Some("array1"),
				},
				Min: o.Some(3),
				Max: o.Some(4),
				Items: &c.CodeGenTypeString{
					CodeGenTypeBase: c.CodeGenTypeBase{
						Description: o.Some("string1"),
					},
					Min: o.Some(1),
					Max: o.Some(2),
				},
			},
		},
		{
			"object1",
			`
type: object
description: object1
properties: 
    -   type: string
        description: string1
        min: 1
        max: 2
`,
			&c.CodeGenTypeObject{
				CodeGenTypeBase: c.CodeGenTypeBase{
					Description: o.Some("object1"),
				},
				Properties: []c.CodeGenType{
					&c.CodeGenTypeString{
						CodeGenTypeBase: c.CodeGenTypeBase{
							Description: o.Some("string1"),
							Name:        o.Some("aname"),
						},
						Min: o.Some(1),
						Max: o.Some(2),
					},
				},
			},
		},
	}
)

func Test_CodeGen_Unmarshal(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clean := strings.Trim(tt.input, " \t")
			actual, err := c.UnmarshalYaml([]byte(clean))
			assert.Nil(t, err)
			assertSchema(t, tt.expected, actual)
		})
	}
}

func assertSchema(t *testing.T, e, a c.CodeGenType) {
	assert.IsType(t, e, a)
	switch ce := e.(type) {
	case *c.CodeGenTypeArray:
		assertSchemaBase(t, *e.Base(), *a.Base())
		assertSchema(t, ce.Items, a.(*c.CodeGenTypeArray).Items)
	case *c.CodeGenTypeDate:
	case *c.CodeGenTypeDateTime:
	case *c.CodeGenTypeDuration:
	case *c.CodeGenTypeEnum:
	case *c.CodeGenTypeFloat:
	case *c.CodeGenTypeInteger:
	case *c.CodeGenTypeObject:
		assertSchemaBase(t, ce.CodeGenTypeBase, a.(*c.CodeGenTypeObject).CodeGenTypeBase)
		assert.Equal(t, len(ce.Properties), len(a.(*c.CodeGenTypeObject).Properties))

		for i, property := range ce.Properties {
			assertSchema(t, property, a.(*c.CodeGenTypeObject).Properties[i])
		}
	case *c.CodeGenTypeString:
		assertSchemaBase(t, ce.CodeGenTypeBase, a.(*c.CodeGenTypeString).CodeGenTypeBase)
		assert.Equal(t, ce.Max, a.(*c.CodeGenTypeString).Max)
		assert.Equal(t, ce.Min, a.(*c.CodeGenTypeString).Min)
	case *c.CodeGenTypeUuid:
	case *c.CodeGenTypeRef:
		assertSchemaBase(t, ce.CodeGenTypeBase, a.(*c.CodeGenTypeRef).CodeGenTypeBase)
		assert.Equal(t, ce.Ref, a.(*c.CodeGenTypeRef).Ref)
	default:
		assert.Fail(t, "invalid type")
	}
}

func assertSchemaBase(t *testing.T, e, a c.CodeGenTypeBase) {
	assert.Equal(t, e.Description, a.Description)
	assert.Equal(t, e.Id, a.Id)
}
