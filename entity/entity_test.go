package entity_test

import (
	"boundedinfinity/codegen/entity"
	"strings"
	"testing"
	"time"
	"unicode"

	"github.com/stretchr/testify/assert"
)

var (
	case1 = entity.String().QName("bounded/type/case-01")
	case2 = entity.String().
		QName("bounded/type/case-02").
		Required(true)
	case3  = entity.String().Min(2)
	case4  = entity.String().StartsWith("something")
	case5  = entity.Integer().Positive()
	case6  = entity.Integer().Range(entity.IntegerRange{Min: 100, Max: 200})
	case7  = entity.Boolean().Required(true)
	case8  = entity.DateTime().Range(entity.DateTimeRanges.NextWeekFrom(time.Time{}.Add(1 * time.Second)))
	case9  = entity.Array().Items(entity.String().QName("something-1"))
	case10 = entity.Array().Items(entity.Ref().Entity(entity.String().QName("something-2")))
	case11 = entity.Object().Properties(
		entity.Ref().Entity(entity.String().QName("something-2")),
		entity.Array().Items(entity.Ref().Entity(entity.String().QName("something-3"))),
	)
	case12 = entity.Operation().Input(
		entity.Ref().Entity(entity.String().QName("something-1")),
		entity.Ref().Entity(entity.String().QName("something-2")),
	).Outputs(
		entity.Ref().Entity(entity.String().QName("something-3")),
		entity.Ref().Entity(entity.String().QName("something-4")),
	).Comments("comments here")
	case13 = entity.Union().Entity(
		entity.Ref().Entity(entity.String().QName("something-1")),
		entity.Ref().Entity(entity.String().QName("something-2")),
	).Comments("comments here").AdditionalValidation(true)
	case14 = entity.Data().Entity(
		entity.Ref().Entity(entity.String().QName("something-1")),
	).Comments("comments here").
		Item(map[string]any{
			"q-name": "something-1",
		})
)

func firstNonSpace(s string) int {
	for i, c := range s {
		if !unicode.IsSpace(c) {
			return i
		}
	}

	return -1
}

func fixIndent(s string) string {
	var longest int
	text := strings.TrimSpace(s)

	for _, line := range strings.Split(text, "\n") {
		i := firstNonSpace(line)

		if i > longest {
			longest = i
		}
	}

	var builder strings.Builder

	for _, line := range strings.Split(text, "\n") {
		i := firstNonSpace(line)

		if i >= longest {
			line = line[longest:]
		}

		builder.WriteString(line + "\n")
	}

	return builder.String()
}

func Test_Entity_Marshal_Json(t *testing.T) {
	tcs := []struct {
		name     string
		input    entity.Marshalable
		expected string
		err      error
	}{
		{
			name:  "case 1",
			input: case1,
			err:   nil,
			expected: `{
		        "entity-type": "string",
		        "q-name": "bounded/type/case-01"
		    }`,
		},
		{
			name:  "case 2",
			input: case2,
			err:   nil,
			expected: `{
		        "entity-type": "string",
		        "q-name": "bounded/type/case-02",
		        "required": true
		    }`,
		},
		{
			name:  "case 3",
			input: case3,
			err:   nil,
			expected: `{
		        "entity-type": "string",
		        "min": 2
		    }`,
		},
		{
			name:  "case 4",
			input: case4,
			err:   nil,
			expected: `{
		        "entity-type": "string",
		        "starts-with": "something"
		    }`,
		},
		{
			name:  "case 5",
			input: case5,
			err:   nil,
			expected: `{
		        "entity-type": "integer",
		        "positive": true
		    }`,
		},
		{
			name:  "case 6",
			input: case6,
			err:   nil,
			expected: `{
		        "entity-type": "integer",
		        "ranges": [
		            {
		                "max": 200,
		                "min": 100
		            }
		        ]
		    }`,
		},
		{
			name:  "case 7",
			input: case7,
			err:   nil,
			expected: `{
		        "entity-type": "boolean",
		    	"required": true
		    }`,
		},
		{
			name:  "case 8",
			input: case8,
			err:   nil,
			expected: `{
		        "entity-type": "date-time",
		        "ranges": [
		            {
		                "max": -62134991999,
		                "min": -62135596799
		            }
		        ]
		    }`,
		},
		{
			name:  "case 9",
			input: case9,
			err:   nil,
			expected: `{
                "entity-type": "array",
                "items": {
                    "entity-type": "string",
                    "q-name": "something-1"
                }
            }`,
		},
		{
			name:  "case 10",
			input: case10,
			err:   nil,
			expected: `{
                "entity-type": "array",
                "items": {
                    "entity-type": "ref",
                    "ref": "something-2"
                }
            }`,
		},
		{
			name:  "case 11",
			input: case11,
			err:   nil,
			expected: `{
                "entity-type": "object",
                "props": [
                    {
                        "entity-type": "ref",
                        "ref": "something-2"
                    },
                    {
                        "entity-type": "array",
                        "items": {
                            "entity-type": "ref",
                            "ref": "something-3"
                        }
                    }
                ]
            }`,
		},
		{
			name:  "case 12",
			input: case12,
			err:   nil,
			expected: `{
                "comments": "comments here",
                "inputs": [
                    {
                        "entity-type": "ref",
                        "ref": "something-1"
                    },
                    {
                        "entity-type": "ref",
                        "ref": "something-2"
                    }
                ],
                "outputs": [
                    {
                        "entity-type": "ref",
                        "ref": "something-3"
                    },
                    {
                        "entity-type": "ref",
                        "ref": "something-4"
                    }
                ]
            }`,
		},
		{
			name:  "case 13",
			input: case13,
			err:   nil,
			expected: `{
                "comments": "comments here",
                "additional-validation": true,
                "entities": [
                    {
                        "entity-type": "ref",
                        "ref": "something-1"
                    },
                    {
                        "entity-type": "ref",
                        "ref": "something-2"
                    }
                ],
                "entity-type": "union"
            }`,
		},
		{
			name:  "case 14",
			input: case14,
			err:   nil,
			expected: `{
                "comments": "comments here",
                "entity": {
                    "entity-type": "ref",
                    "ref": "something-1"
                },
                "items": [
                    {
                        "q-name": "something-1"
                    }
                ]
            }`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := tc.input.ToJsonIndent()
			actual := string(bs)

			assert.ErrorIs(tt, err, tc.err, actual)
			assert.JSONEqf(tt, tc.expected, actual, actual)
		})
	}
}

func Test_Entity_Marshal_Yaml(t *testing.T) {
	tcs := []struct {
		name     string
		input    entity.Marshalable
		expected string
		err      error
	}{
		{
			name:  "case 1",
			input: case1,
			err:   nil,
			expected: `
                entity-type: string
                q-name: bounded/type/case-01
            `,
		},
		// {
		// 	name:  "case 2",
		// 	input: case2,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "string",
		//         "q-name": "bounded/type/case-02",
		//         "required": true
		//     }`,
		// },
		// {
		// 	name:  "case 3",
		// 	input: case3,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "string",
		//         "min": 2
		//     }`,
		// },
		// {
		// 	name:  "case 4",
		// 	input: case4,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "string",
		//         "starts-with": "something"
		//     }`,
		// },
		// {
		// 	name:  "case 5",
		// 	input: case5,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "integer",
		//         "positive": true
		//     }`,
		// },
		// {
		// 	name:  "case 6",
		// 	input: case6,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "integer",
		//         "ranges": [
		//             {
		//                 "max": 200,
		//                 "min": 100
		//             }
		//         ]
		//     }`,
		// },
		// {
		// 	name:  "case 7",
		// 	input: case7,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "boolean",
		//     	"required": true
		//     }`,
		// },
		// {
		// 	name:  "case 8",
		// 	input: case8,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "date-time",
		//         "ranges": [
		//             {
		//                 "max": -62134991999,
		//                 "min": -62135596799
		//             }
		//         ]
		//     }`,
		// },
		// {
		// 	name:  "case 9",
		// 	input: case9,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "array",
		//         "items": {
		//             "entity-type": "string",
		//             "q-name": "something-1"
		//         }
		//     }`,
		// },
		// {
		// 	name:  "case 10",
		// 	input: case10,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "array",
		//         "items": {
		//             "entity-type": "ref",
		//             "ref": "something-2"
		//         }
		//     }`,
		// },
		// {
		// 	name:  "case 11",
		// 	input: case11,
		// 	err:   nil,
		// 	expected: `{
		//         "entity-type": "object",
		//         "props": [
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-2"
		//             },
		//             {
		//                 "entity-type": "array",
		//                 "items": {
		//                     "entity-type": "ref",
		//                     "ref": "something-3"
		//                 }
		//             }
		//         ]
		//     }`,
		// },
		// {
		// 	name:  "case 12",
		// 	input: case12,
		// 	err:   nil,
		// 	expected: `{
		//         "comments": "comments here",
		//         "inputs": [
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-1"
		//             },
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-2"
		//             }
		//         ],
		//         "outputs": [
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-3"
		//             },
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-4"
		//             }
		//         ]
		//     }`,
		// },
		// {
		// 	name:  "case 13",
		// 	input: case13,
		// 	err:   nil,
		// 	expected: `{
		//         "comments": "comments here",
		//         "additional-validation": true,
		//         "entities": [
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-1"
		//             },
		//             {
		//                 "entity-type": "ref",
		//                 "ref": "something-2"
		//             }
		//         ],
		//         "entity-type": "union"
		//     }`,
		// },
		// {
		// 	name:  "case 14",
		// 	input: case14,
		// 	err:   nil,
		// 	expected: `{
		//         "comments": "comments here",
		//         "entity": {
		//             "entity-type": "ref",
		//             "ref": "something-1"
		//         },
		//         "items": [
		//             {
		//                 "q-name": "something-1"
		//             }
		//         ]
		//     }`,
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			bs, err := tc.input.ToYaml()
			actual := string(bs)
			expected := fixIndent(tc.expected)

			assert.ErrorIs(tt, err, tc.err, actual)
			assert.YAMLEq(tt, expected, actual, actual)
		})
	}
}
