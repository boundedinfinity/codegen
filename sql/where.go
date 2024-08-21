package sql

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Where() *WhereClauseSchema {
	return &WhereClauseSchema{
		clauses: []whereFragment{
			&joinFragment{operator: "WHERE"},
		},
	}
}

type WhereClauseSchema struct {
	clauses []whereFragment
}

func (this *WhereClauseSchema) Generate() string {
	var sb stringBuiler
	var clauses []string

	for _, clause := range this.clauses {
		clauses = append(clauses, clause.Generate())
	}

	sb.Writef(stringer.Join(" ", clauses...))

	return sb.String()
}

func (this *WhereClauseSchema) And() *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&joinFragment{operator: "AND"}))
}

func (this *WhereClauseSchema) Or() *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&joinFragment{operator: "OR"}))
}

func (this *WhereClauseSchema) Not() *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&joinFragment{operator: "NOT"}))
}

func (this *WhereClauseSchema) Equal(column *ColumnSchema) *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&comparisonFragment{
		operator: "=",
		column:   column,
	}))
}

func (this *WhereClauseSchema) NotEqual(column *ColumnSchema) *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&comparisonFragment{
		operator: "!=",
		column:   column,
	}))
}

func (this *WhereClauseSchema) Like(column *ColumnSchema, pattern, escape string) *WhereClauseSchema {
	return appendAndReturn(this, &this.clauses, whereFragment(&likeFragment{
		pattern: pattern,
		escape:  escape,
		column:  column,
	}))
}

type whereFragment interface {
	Generate() string
}

var _ whereFragment = &comparisonFragment{}

type comparisonFragment struct {
	operator string
	column   *ColumnSchema
}

func (this *comparisonFragment) Generate() string {
	return fmt.Sprintf("%s %s ?", this.column.Name, this.operator)
}

var _ whereFragment = &joinFragment{}

type joinFragment struct {
	operator string
}

func (this *joinFragment) Generate() string {
	return this.operator
}

var _ whereFragment = &likeFragment{}

// https://www.sqlitetutorial.net/sqlite-like/

type likeFragment struct {
	pattern string
	escape  string
	column  *ColumnSchema
}

func (this *likeFragment) Generate() string {
	if this.escape == "" {
		return fmt.Sprintf("%s LIKE '%s'", this.column.Name, this.pattern)
	}

	return fmt.Sprintf("%s LIKE '%s' ESCAPE '%s'", this.column.Name, this.pattern, this.escape)
}

// type whereKeyword string

// var whereKeywords = whereLogicalOperators{
// 	where:                "WHERE",
// 	All:                  "ALL",
// 	And:                  "AND",
// 	Any:                  "ANY",
// 	Between:              "BETWEEN",
// 	Exists:               "EXISTS",
// 	In:                   "IN",
// 	Like:                 "LIKE",
// 	Not:                  "NOT",
// 	Or:                   "OR",
// 	EqualTo:              "=",
// 	NotEqualTo:           "!=",
// 	LessThan:             "<",
// 	GreaterThan:          ">",
// 	LessThanOrEqualTo:    "<=",
// 	GreaterThanOrEqualTo: ">=",
// }

// type whereLogicalOperators struct {
// 	where                whereKeyword
// 	All                  whereKeyword
// 	And                  whereKeyword
// 	Any                  whereKeyword
// 	Between              whereKeyword
// 	Exists               whereKeyword
// 	In                   whereKeyword
// 	Like                 whereKeyword
// 	Not                  whereKeyword
// 	Or                   whereKeyword
// 	EqualTo              whereKeyword
// 	NotEqualTo           whereKeyword
// 	LessThan             whereKeyword
// 	GreaterThan          whereKeyword
// 	LessThanOrEqualTo    whereKeyword
// 	GreaterThanOrEqualTo whereKeyword
// }
