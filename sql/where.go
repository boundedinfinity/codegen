package sql

import (
	"github.com/boundedinfinity/go-commoner/functional/optioner"
)

func Where() *WhereClauseSchema {
	return &WhereClauseSchema{}
}

type WhereClauseSchema struct {
	column     *ColumnSchema
	comparison optioner.Option[WhereComparisonOperator]
	logical    optioner.Option[WhereLogicalOperator]
}

func (this *WhereClauseSchema) Generate() string {
	var sb stringBuiler

	comparison := this.comparison.OrElse(WhereComparisonOperators.EqualTo)

	if this.logical.Defined() {
		sb.Writef("%s ", this.logical.Get())
	}

	sb.Writef("%s %s ?", quote(this.column.Name), comparison)

	return sb.String()
}

func (this *WhereClauseSchema) Table(column *ColumnSchema) *WhereClauseSchema {
	return setAndReturn(this, &this.column, column)
}

func (this *WhereClauseSchema) Compare(operator WhereComparisonOperator) *WhereClauseSchema {
	return setAndReturn(this, &this.comparison, optioner.Some(operator))
}

func (this *WhereClauseSchema) Logical(operator WhereLogicalOperator) *WhereClauseSchema {
	return setAndReturn(this, &this.logical, optioner.Some(operator))
}

type WhereComparisonOperator string

var WhereComparisonOperators = whereComparisonOperators{
	EqualTo:              "=",
	NotEqualTo:           "!=",
	LessThan:             "<",
	GreaterThan:          ">",
	LessThanOrEqualTo:    "<=",
	GreaterThanOrEqualTo: ">=",
}

type whereComparisonOperators struct {
	EqualTo              WhereComparisonOperator
	NotEqualTo           WhereComparisonOperator
	LessThan             WhereComparisonOperator
	GreaterThan          WhereComparisonOperator
	LessThanOrEqualTo    WhereComparisonOperator
	GreaterThanOrEqualTo WhereComparisonOperator
}

type WhereLogicalOperator string

var WhereLogicalOperators = whereLogicalOperators{
	All:     "ALL",
	And:     "AND",
	Any:     "ANY",
	Between: "BETWEEN",
	Exists:  "EXISTS",
	In:      "IN",
	Like:    "LIKE",
	Not:     "NOT",
	Or:      "OR",
}

type whereLogicalOperators struct {
	All     WhereComparisonOperator
	And     WhereComparisonOperator
	Any     WhereComparisonOperator
	Between WhereComparisonOperator
	Exists  WhereComparisonOperator
	In      WhereComparisonOperator
	Like    WhereComparisonOperator
	Not     WhereComparisonOperator
	Or      WhereComparisonOperator
}
