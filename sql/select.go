package sql

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Select() *SelectSchema {
	return &SelectSchema{}
}

type SelectSchema struct {
	table     *TableSchema
	columns   []*ColumnSchema
	clauses   [][]*WhereClauseSchema
	formatted bool
}

func (this *SelectSchema) Generate() string {
	var sb stringBuiler

	create := func(whereColumns []*WhereClauseSchema) {
		columnNames := columnNames(this.table.Columns)
		columnNames = quotes(columnNames)

		sb.Writef("SELECT %s FROM %s", stringer.Join(", ", columnNames...), this.table.Name)

		if len(whereColumns) > 0 {
			sb.Writef(" WHERE ")

			for _, whereClauses := range this.clauses {
				whereClauses := slicer.Map(
					func(_ int, clause *WhereClauseSchema) string { return clause.Generate() },
					whereClauses...)
				sb.Writef(stringer.Join(" ", whereClauses...))
			}
		}

		sb.Writef(";")

		if this.formatted {
			sb.Writef("\n")
		}
	}

	create([]*WhereClauseSchema{})

	for _, wheres := range this.clauses {
		create(wheres)
	}

	return sb.String()
}

func (this *SelectSchema) Table(table *TableSchema) *SelectSchema {
	return setAndReturn(this, &this.table, table)
}

func (this *SelectSchema) Formatted(enabed bool) *SelectSchema {
	return setAndReturn(this, &this.formatted, enabed)
}

func (this *SelectSchema) Column(columns ...*ColumnSchema) *SelectSchema {
	return appendAndReturn(this, &this.columns, columns...)
}

func (this *SelectSchema) Where(clauses ...*WhereClauseSchema) *SelectSchema {
	return appendAndReturn(this, &this.clauses, clauses)
}
