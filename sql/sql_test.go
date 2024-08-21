package sql_test

import (
	"boundedinfinity/codegen/sql"
	"context"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/stretchr/testify/assert"
)

func Test_Generate(t *testing.T) {
	database := sql.Database().SetForeignKeys(true).SetFormatted(true)

	person_table := sql.Table().
		SetName("person").
		Column(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true})

	name_sub_table := func(name string) *sql.TableSchema {
		table := sql.Table().
			SetName(name).
			SetWithoutRowId(true).
			Column(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
			Column(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
			Column(sql.Column().SetName("index").SetType(sql.ColumnTypes.INTEGER))

		return table
	}

	name_first_table := name_sub_table("name_first")
	name_middle_table := name_sub_table("name_middle")
	name_last_table := name_sub_table("name_last")

	name_table := sql.Table().
		SetName("name").
		SetWithoutRowId(true).
		Column(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true))

	name_type := sql.Table().
		SetName("name_type").
		SetWithoutRowId(true).
		Column(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
		Column(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
		Column(sql.Column().SetName("description").SetType(sql.ColumnTypes.TEXT))

	label_table := sql.Table().
		SetName("label").
		Column(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
		Column(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
		Column(sql.Column().SetName("description").SetType(sql.ColumnTypes.TEXT))

	database.
		Table(name_table).
		Table(name_first_table).
		Table(name_middle_table).
		Table(name_last_table).
		Table(name_type).
		Table(person_table).
		Table(label_table).
		OneToMany(name_table, name_first_table).
		OneToMany(name_table, name_middle_table).
		OneToMany(name_table, name_last_table).
		ReferencedTo(name_table, name_type).
		OneToMany(person_table, name_table).
		ManyToMany(person_table, label_table)

	tcs := []struct {
		name     string
		input    *sql.DatabaseSchema
		expected string
	}{
		{
			name:     "case 1",
			input:    database,
			expected: ``,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Generate()
			path := testOutputPath(tc.name, "sql")
			assert.ErrorIs(tt, tc.input.WriteSqlFile(path), nil)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}

func testOutputPath(name, ext string) string {
	_, err := pather.Dirs.EnsureErr("./test-output")
	if err != nil {
		panic(err)
	}
	filename := extentioner.Join(name, ext)
	path := pather.Paths.Join("./test-output", filename)
	return path
}

func Test_Database(t *testing.T) {
	database := sql.Database().SetForeignKeys(true).SetFormatted(true)

	person_table := sql.Table().
		SetIsNotExists(true).SetName("person").
		Column(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true})

	database.
		Table(person_table)

	tcs := []struct {
		name     string
		input    *sql.DatabaseSchema
		expected string
	}{
		{
			name:     "case 1",
			input:    database,
			expected: ``,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			path := testOutputPath(tc.name, "db")

			manager, err := sql.NewManager(path, context.Background())
			assert.ErrorIs(tt, err, nil)
			defer manager.Cleanup()

			err = manager.EnsureTables(*tc.input)
			assert.ErrorIs(tt, err, nil)
		})
	}
}
