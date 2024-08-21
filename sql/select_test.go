package sql_test

import (
	"boundedinfinity/codegen/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Select_Generate(t *testing.T) {
	person_table := sql.Table().
		SetIsNotExists(true).SetName("person").
		Column(&sql.ColumnSchema{Name: "id", Type: sql.ColumnTypes.TEXT, PrimaryKey: true}).
		Column(&sql.ColumnSchema{Name: "first_name", Type: sql.ColumnTypes.TEXT}).
		Column(&sql.ColumnSchema{Name: "last_name", Type: sql.ColumnTypes.TEXT})

	_ = sql.Database().
		SetForeignKeys(true).
		SetFormatted(true).
		Table(person_table)

	tcs := []struct {
		name     string
		input    *sql.SelectSchema
		expected string
		err      error
	}{
		{
			name:     "case 1",
			input:    sql.Select().All(person_table),
			expected: `SELECT id, first_name, last_name FROM person;`,
		},
		{
			name: "case 2",
			input: sql.Select().
				All(person_table).
				Where(sql.Where().
					Equal(person_table.GetColumn("id")).
					And().
					Equal(person_table.GetColumn("first_name")),
				),
			expected: `SELECT id, first_name, last_name FROM person WHERE id = ? AND first_name = ?;`,
		},
		{
			name: "case 3",
			input: sql.Select().
				All(person_table).
				Where(sql.Where().
					Equal(person_table.GetColumn("id")).
					And().Equal(person_table.GetColumn("first_name")).
					Or().NotEqual(person_table.GetColumn("last_name")),
				),
			expected: `SELECT id, first_name, last_name FROM person WHERE id = ? AND first_name = ?;`,
		},
	}

	for _, tc := range tcs {
		t.Run(tc.name, func(tt *testing.T) {
			actual, err := tc.input.Generate()

			assert.ErrorIs(tt, err, tc.err)
			assert.Equal(tt, tc.expected, actual)
		})
	}
}
