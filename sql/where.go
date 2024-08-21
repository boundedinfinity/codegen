package sql

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func Where() *WhereClauseSchema {
	return &WhereClauseSchema{
		fragments: []whereFragment{
			&joinFragment{operator: "WHERE"},
		},
	}
}

type WhereClauseSchema struct {
	fragments []whereFragment
}

func (this *WhereClauseSchema) Generate() string {
	var sb stringBuiler
	var fragments []string

	for _, fragment := range this.fragments {
		fragments = append(fragments, fragment.Generate())
	}

	sb.Writef(stringer.Join(" ", fragments...))

	return sb.String()
}

func (this *WhereClauseSchema) And() *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&joinFragment{operator: "AND"}))
}

func (this *WhereClauseSchema) Or() *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&joinFragment{operator: "OR"}))
}

func (this *WhereClauseSchema) Not() *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&joinFragment{operator: "NOT"}))
}

func (this *WhereClauseSchema) Equal(column *ColumnSchema) *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&comparisonFragment{
		operator: "=",
		column:   column,
	}))
}

func (this *WhereClauseSchema) NotEqual(column *ColumnSchema) *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&comparisonFragment{
		operator: "!=",
		column:   column,
	}))
}

func (this *WhereClauseSchema) Like(column *ColumnSchema, pattern, escape string) *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&likeFragment{
		pattern: pattern,
		escape:  escape,
		column:  column,
	}))
}

func (this *WhereClauseSchema) InList(column *ColumnSchema, items ...any) *WhereClauseSchema {
	return appendAndReturn(this, &this.fragments, whereFragment(&inListFragment{
		items:  items,
		column: column,
	}))
}

type whereFragment interface {
	Generate() string
}

var _ whereFragment = &inListFragment{}

type inListFragment struct {
	items  []any
	column *ColumnSchema
}

func (this *inListFragment) Generate() string {
	return fmt.Sprintf("%s IN (%s)",
		getColumnName(true, this.column),
		stringer.Join(", ", stringer.AsStrings(this.items...)...),
	)
}

var _ whereFragment = &comparisonFragment{}

type comparisonFragment struct {
	operator string
	column   *ColumnSchema
}

func (this *comparisonFragment) Generate() string {
	return fmt.Sprintf("%s %s ?", getColumnName(true, this.column), this.operator)
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
	name := getColumnName(true, this.column)

	if this.escape == "" {
		return fmt.Sprintf("%s LIKE '%s'", name, this.pattern)
	}

	return fmt.Sprintf("%s LIKE '%s' ESCAPE '%s'", name, this.pattern, this.escape)
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
