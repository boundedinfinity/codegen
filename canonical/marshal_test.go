package canonical_test

import (
	c "boundedinfinity/codegen/canonical"
	"strings"
	"testing"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/stretchr/testify/assert"
)

var (
	tests = []struct {
		name     string
		input    string
		expected c.Canonical
	}{
		{
			"ref1",
			`
ref: https://www.boundedinfinity.com/schema/something
description: ref1
`,
			&c.CanonicalRef{
				CanonicalBase: c.CanonicalBase{
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
			&c.CanonicalString{
				CanonicalBase: c.CanonicalBase{
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
			&c.CanonicalArray{
				CanonicalBase: c.CanonicalBase{
					Description: o.Some("array1"),
				},
				Min: o.Some(3),
				Max: o.Some(4),
				Items: &c.CanonicalString{
					CanonicalBase: c.CanonicalBase{
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
			&c.CanonicalObject{
				CanonicalBase: c.CanonicalBase{
					Description: o.Some("object1"),
				},
				Properties: []c.Canonical{
					&c.CanonicalString{
						CanonicalBase: c.CanonicalBase{
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

func Test_Canonical_Unmarshal(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			clean := strings.Trim(tt.input, " \t")
			actual, err := c.UnmarshalCanonicalSchemaYaml([]byte(clean))
			assert.Nil(t, err)
			assertSchema(t, tt.expected, actual)
		})
	}
}

func assertSchema(t *testing.T, e, a c.Canonical) {
	assert.IsType(t, e, a)
	switch ce := e.(type) {
	case *c.CanonicalArray:
		assertSchemaBase(t, ce.CanonicalBase, a.(*c.CanonicalArray).CanonicalBase)
		assertSchema(t, ce.Items, a.(*c.CanonicalArray).Items)
	case *c.CanonicalDate:
	case *c.CanonicalDateTime:
	case *c.CanonicalDuration:
	case *c.CanonicalEnum:
	case *c.CanonicalFloat:
	case *c.CanonicalInteger:
	case *c.CanonicalObject:
		assertSchemaBase(t, ce.CanonicalBase, a.(*c.CanonicalObject).CanonicalBase)
		assert.Equal(t, len(ce.Properties), len(a.(*c.CanonicalObject).Properties))

		for i, property := range ce.Properties {
			assertSchema(t, property, a.(*c.CanonicalObject).Properties[i])
		}
	case *c.CanonicalString:
		assertSchemaBase(t, ce.CanonicalBase, a.(*c.CanonicalString).CanonicalBase)
		assert.Equal(t, ce.Max, a.(*c.CanonicalString).Max)
		assert.Equal(t, ce.Min, a.(*c.CanonicalString).Min)
	case *c.CanonicalUuid:
	case *c.CanonicalRef:
		assertSchemaBase(t, ce.CanonicalBase, a.(*c.CanonicalRef).CanonicalBase)
		assert.Equal(t, ce.Ref, a.(*c.CanonicalRef).Ref)
	default:
		assert.Fail(t, "invalid type")
	}
}

func assertSchemaBase(t *testing.T, e, a c.CanonicalBase) {
	assert.Equal(t, e.Description, a.Description)
	assert.Equal(t, e.Id, a.Id)
}
