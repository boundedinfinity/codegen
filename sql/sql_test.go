package sql_test

import (
	"boundedinfinity/codegen/sql"
	"testing"

	"github.com/boundedinfinity/go-commoner/idiomatic/extentioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/stretchr/testify/assert"
)

func Test_Unmarshal_Object(t *testing.T) {
	database := sql.Database().SetForeignKeys(true).SetFormatted(true)

	person_table := sql.Table().
		SetName("person").
		AddColumn(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true})

	name_sub_table := func(name string) *sql.TableSchema {
		table := sql.Table().
			SetName(name).
			SetWithoutRowId(true).
			AddColumn(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
			AddColumn(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
			AddColumn(sql.Column().SetName("index").SetType(sql.ColumnTypes.INTEGER))

		return table
	}

	name_first_table := name_sub_table("name_first")
	name_middle_table := name_sub_table("name_middle")
	name_last_table := name_sub_table("name_last")

	name_table := sql.Table().
		SetName("name").
		SetWithoutRowId(true).
		AddColumn(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true))

	name_type := sql.Table().
		SetName("name_type").
		SetWithoutRowId(true).
		AddColumn(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
		AddColumn(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
		AddColumn(sql.Column().SetName("description").SetType(sql.ColumnTypes.TEXT))

	label_table := sql.Table().
		SetName("label").
		AddColumn(sql.Column().SetName("id").SetType(sql.ColumnTypes.TEXT).SetPrimaryKey(true)).
		AddColumn(sql.Column().SetName("name").SetType(sql.ColumnTypes.TEXT).SetIndexed(true)).
		AddColumn(sql.Column().SetName("description").SetType(sql.ColumnTypes.TEXT))

	database.
		AddTable(name_table).
		AddTable(name_first_table).
		AddTable(name_middle_table).
		AddTable(name_last_table).
		AddTable(name_type).
		AddTable(person_table).
		AddTable(label_table).
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
			_, err := pather.Dirs.EnsureErr("./test-output")
			assert.ErrorIs(tt, err, nil)
			filename := extentioner.Join(tc.name, ".sql")
			path := pather.Paths.Join("./test-output", filename)
			assert.ErrorIs(tt, tc.input.WriteSqlFile(path), nil)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
