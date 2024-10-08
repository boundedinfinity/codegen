package generator_test

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/model"
	"context"
	"database/sql"
	"fmt"
	"os"
	"testing"

	o "github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/stretchr/testify/assert"

	_ "github.com/tursodatabase/go-libsql"
)

var cleanupDb bool = false

func storeNfixTypes(nfixTypes map[string]NfixType) error {

	return nil
}

func storePrefixes(nfixTypes map[string]Prefix) error {

	return nil
}

func storeSuffixes(nfixTypes map[string]Suffix) error {

	return nil
}

func storePeople(nfixTypes map[string]Suffix) error {

	return nil
}

func Test_Db(t *testing.T) {
	nfixTypes := createNfixTypes()
	assert.Nil(t, storeNfixTypes(nfixTypes))

	file := "./test-output/test-0.db"
	db, err := sql.Open("libsql", fmt.Sprintf("file:%s", file))
	assert.Nil(t, err)

	defer func() {
		db.Close()
		if cleanupDb {
			os.Remove(file)
		}
	}()

	ctx := context.Background()

	_, err = db.ExecContext(ctx, `CREATE TABLE IF NOT EXISTS test_0 ( column_1 TEXT );`)
	assert.Nil(t, err)

	_, err = db.ExecContext(ctx, `INSERT INTO test_0 ( column_1 ) VALUES ('something');`)
	assert.Nil(t, err)

	rows, err := db.QueryContext(ctx, `SELECT column_1 FROM test_0;`)
	assert.Nil(t, err)

	var column_1 string
	for rows.Next() {
		assert.Nil(t, rows.Scan(&column_1))
		fmt.Printf("column_1: %s\n", column_1)
	}
}

func Test_Generate(t *testing.T) {
	tcs := []struct {
		name     string
		lang     string
		input    model.CodeGenSchema
		expected map[string]string
		err      error
		newErr   error
	}{
		{
			name: "object-01",
			lang: "go",
			input: &model.CodeGenObject{
				CodeGenCommon: model.CodeGenCommon{
					Id: o.Some("test-output/codegen/schema/label/label"),
					Lang: model.CodeGenLangData{
						Package: o.Some("label"),
						Name:    o.Some("Label"),
						Imports: o.Some([]string{
							"boundedinfinity/codegen/validation",
							"errors",
							"boundedinfinity/codegen/schema/util",
						}),
					},
				},
				Properties: o.Some([]model.CodeGenSchema{
					&model.CodeGenRef{
						CodeGenCommon: model.CodeGenCommon{
							Name: o.Some("name"),
							Lang: model.CodeGenLangData{
								Name: o.Some("Name"),
								Type: o.Some("util.String01"),
							},
						},
						Ref: o.Some("test-output/codegen/schema/util/string-01"),
						Resolved: &model.CodeGenString{
							CodeGenCommon: model.CodeGenCommon{
								Id: o.Some("test-output/codegen/schema/util/string-01"),
								Lang: model.CodeGenLangData{
									Name: o.Some("String01"),
									Type: o.Some("String01"),
								},
							},
							Min:   o.Some(2),
							Max:   o.Some(10),
							Regex: o.Some(`.*`),
						},
					},

					&model.CodeGenString{
						CodeGenCommon: model.CodeGenCommon{
							Lang: model.CodeGenLangData{
								Name: o.Some("AString"),
								Type: o.Some("string"),
							},
						},
						Min:   o.Some(2),
						Max:   o.Some(10),
						Regex: o.Some(`.*`),
					},

					&model.CodeGenInteger{
						Number: model.Number[int]{
							CodeGenCommon: model.CodeGenCommon{
								Id: o.Some("test-output/codegen/schema/util/integer-01"),
								Lang: model.CodeGenLangData{
									Package: o.Some("util"),
									Name:    o.Some("Integer01"),
									Type:    o.Some("int"),
									Imports: o.Some([]string{
										"errors",
										"database/sql/driver",
										"boundedinfinity/codegen/validation",
									}),
								},
							},
							Max:        o.Some(100),
							Min:        o.Some(5),
							MultipleOf: o.Some(5),
							Positive:   o.Some(true),
						},
					},
					&model.CodeGenArray{
						CodeGenCommon: model.CodeGenCommon{
							Name: o.Some("description"),
							Lang: model.CodeGenLangData{
								Name: o.Some("Description"),
								Type: o.Some("[]string"),
							},
						},
						Items: o.Some(model.CodeGenSchema(&model.CodeGenString{
							CodeGenCommon: model.CodeGenCommon{
								Lang: model.CodeGenLangData{
									Type: o.Some("string"),
								},
							},
							Min:   o.Some(2),
							Max:   o.Some(10),
							Regex: o.Some(`.*`),
						})),
					},
				}),
			},
			expected: map[string]string{},
		},
		// {
		// 	name: "object-string-01",
		// 	lang: "go",
		// 	input: &model.CodeGenObject{
		// 		CodeGenCommon: model.CodeGenCommon{
		// 			Id: o.Some("test-output/codegen/schema/label/label"),
		// 			Lang: model.CodeGenLangData{
		// 				Package: o.Some("label"),
		// 				Name:    o.Some("Label"),
		// 				Imports: o.Some([]string{
		// 					"boundedinfinity/codegen/validation",
		// 					"errors",
		// 					"boundedinfinity/codegen/schema/util",
		// 				}),
		// 			},
		// 		},
		// 		Properties: o.Some([]model.CodeGenSchema{
		// 			&model.CodeGenRef{
		// 				CodeGenCommon: model.CodeGenCommon{
		// 					Name: o.Some("name"),
		// 					Lang: model.CodeGenLangData{
		// 						Name: o.Some("Name"),
		// 						Type: o.Some("util.String01"),
		// 					},
		// 				},
		// 				Ref: o.Some("test-output/codegen/schema/util/string-01"),
		// 				Resolved: &model.CodeGenString{
		// 					CodeGenCommon: model.CodeGenCommon{
		// 						Id: o.Some("test-output/codegen/schema/util/string-01"),
		// 						Lang: model.CodeGenLangData{
		// 							Name: o.Some("String01"),
		// 							Type: o.Some("String01"),
		// 						},
		// 					},
		// 					Min:   o.Some(2),
		// 					Max:   o.Some(10),
		// 					Regex: o.Some(`.*`),
		// 				},
		// 			},

		// 			&model.CodeGenString{
		// 				CodeGenCommon: model.CodeGenCommon{
		// 					Lang: model.CodeGenLangData{
		// 						Name: o.Some("AString"),
		// 						Type: o.Some("string"),
		// 					},
		// 				},
		// 				Min:   o.Some(2),
		// 				Max:   o.Some(10),
		// 				Regex: o.Some(`.*`),
		// 			},

		// 			&model.CodeGenArray{
		// 				CodeGenCommon: model.CodeGenCommon{
		// 					Name: o.Some("description"),
		// 					Lang: model.CodeGenLangData{
		// 						Name: o.Some("Description"),
		// 						Type: o.Some("[]string"),
		// 					},
		// 				},
		// 				Items: o.Some(model.CodeGenSchema(&model.CodeGenString{
		// 					CodeGenCommon: model.CodeGenCommon{
		// 						Lang: model.CodeGenLangData{
		// 							Type: o.Some("string"),
		// 						},
		// 					},
		// 					Min:   o.Some(2),
		// 					Max:   o.Some(10),
		// 					Regex: o.Some(`.*`),
		// 				})),
		// 			},
		// 		}),
		// 	},
		// 	expected: map[string]string{},
		// },
		{
			name: "string-01",
			lang: "go",
			input: &model.CodeGenString{
				CodeGenCommon: model.CodeGenCommon{
					Id: o.Some("test-output/codegen/schema/util/string-01"),
					Lang: model.CodeGenLangData{
						Package: o.Some("util"),
						Name:    o.Some("String01"),
						Type:    o.Some("string"),
						Imports: o.Some([]string{
							"errors",
							"database/sql/driver",
							"boundedinfinity/codegen/validation",
						}),
					},
				},
				Max:   o.Some(50),
				Min:   o.Some(1),
				Regex: o.Some(`.*`),
			},
			expected: map[string]string{},
		},
		{
			name: "integer-01",
			lang: "go",
			input: &model.CodeGenInteger{
				Number: model.Number[int]{
					CodeGenCommon: model.CodeGenCommon{
						Id: o.Some("test-output/codegen/schema/util/integer-01"),
						Lang: model.CodeGenLangData{
							Package: o.Some("util"),
							Name:    o.Some("Integer01"),
							Type:    o.Some("int"),
							Imports: o.Some([]string{
								"errors",
								"database/sql/driver",
								"boundedinfinity/codegen/validation",
							}),
						},
					},
					Max:        o.Some(100),
					Min:        o.Some(5),
					MultipleOf: o.Some(5),
					Positive:   o.Some(true),
				},
			},
			expected: map[string]string{},
		},
		// {
		// 	name: "float-01",
		// 	lang: "go",
		// 	input: &model.CodeGenFloat{
		// 		Number: model.Number[float64]{
		// 			CodeGenCommon: model.CodeGenCommon{
		// 				Id: o.Some("test-output/codegen/schema/util/float-01"),
		// 				Lang: model.CodeGenLangData{
		// 					Package: o.Some("util"),
		// 					Name:    o.Some("Float01"),
		// 					Type:    o.Some("float64"),
		// 					Imports: o.Some([]string{
		// 						"errors",
		// 						"database/sql/driver",
		// 						"boundedinfinity/codegen/validation",
		// 					}),
		// 				},
		// 			},
		// 			Max:        o.Some(100.0),
		// 			Min:        o.Some(5.0),
		// 			MultipleOf: o.Some(5.0),
		// 			Positive:   o.Some(true),
		// 		},
		// 	},
		// 	expected: map[string]string{},
		// },
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			gen, err := generator.New(tc.lang)
			assert.ErrorIs(tt, err, tc.newErr)
			// actual, err := gen.GenerateType(tc.input)
			_, err = gen.WriteType(tc.input)
			assert.ErrorIs(tt, err, tc.err)
			// assert.Equal(tt, tc.expected, actual)
		})
	}
}
