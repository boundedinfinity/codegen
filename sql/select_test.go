package sql_test

import (
	"boundedinfinity/codegen/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Select_Generate(t *testing.T) {
	person_table := sql.Table().
		SetIsNotExists(true).SetName("person").
		AddColumn(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true}).
		AddColumn(&sql.ColumnSchema{Name: "first_name", Type: sql.ColumnTypes.TEXT}).
		AddColumn(&sql.ColumnSchema{Name: "last_name", Type: sql.ColumnTypes.TEXT})

	_ = sql.Database().
		SetForeignKeys(true).
		SetFormatted(true).
		AddTable(person_table)

	tcs := []struct {
		name     string
		input    *sql.SelectSchema
		expected string
	}{
		{
			name: "case 1",
			input: sql.Select().
				Table(person_table).
				Where(sql.Where().Table(person_table.GetColumn("id"))).Formatted(true).
				Where(sql.Where().Table(person_table.GetColumn("first_name"))).Formatted(true),
			expected: ``,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual := tc.input.Generate()

			assert.Equal(tt, actual, "")
		})
	}
}
