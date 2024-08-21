package sql

import (
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Select() *SelectSchema {
	return &SelectSchema{}
}

type SelectSchema struct {
	columns   []*ColumnSchema
	clauses   []*WhereClauseSchema
	formatted bool
}

func (this *SelectSchema) Generate() (string, error) {
	var sb stringBuiler
	var tables []*TableSchema

	if len(this.columns) <= 0 {
		return sb.String(), ErrSelectNoColumns
	}

	tables = slicer.Map(func(_ int, column *ColumnSchema) *TableSchema {
		return column.Table
	}, this.columns...)

	tables = slicer.UniqFn(func(_ int, table *TableSchema) string {
		return table.Name
	}, tables...)

	sb.Writef("SELECT ")

	if len(tables) > 1 {
		sb.Writef(stringer.Join(", ", getQualifiedColumnNames(this.columns)...))
	} else {
		sb.Writef(stringer.Join(", ", getColumnNames(this.columns)...))
	}

	sb.Writef(" FROM %s", stringer.Join(", ", getTableNames(tables)...))

	if len(this.clauses) > 0 {
		clauses := slicer.Map(func(_ int, clause *WhereClauseSchema) string {
			return clause.Generate()
		}, this.clauses...)

		sb.Writef(" %s", stringer.Join(" ", clauses...))
	}

	sb.Writef(";")

	if this.formatted {
		sb.Writef("\n")
	}

	return sb.String(), nil
}

func (this *SelectSchema) All(tables ...*TableSchema) *SelectSchema {
	var columns []*ColumnSchema

	for _, table := range tables {
		columns = append(columns, table.Columns...)
	}

	return appendAndReturn(this, &this.columns, columns...)
}

func (this *SelectSchema) Column(columns ...*ColumnSchema) *SelectSchema {
	return appendAndReturn(this, &this.columns, columns...)
}

func (this *SelectSchema) Where(clauses ...*WhereClauseSchema) *SelectSchema {
	return appendAndReturn(this, &this.clauses, clauses...)
}

func (this *SelectSchema) Formatted(enabed bool) *SelectSchema {
	return setAndReturn(this, &this.formatted, enabed)
}
