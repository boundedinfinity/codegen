// Package sql typing information
package sql

import (
	"fmt"
	"strings"

	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

type stringBuiler struct {
	sb strings.Builder
}

func (this *stringBuiler) Writef(format string, a ...any) {
	this.sb.WriteString(fmt.Sprintf(format, a...))
}

func (this stringBuiler) String() string {
	return this.sb.String()
}

func quote(s string) string {
	for _, keyword := range keywords {
		if s == keyword {
			return fmt.Sprintf(`"%s"`, s)
		}
	}

	return s
}

func quotes(s []string) []string {
	return slicer.Map(func(_ int, s string) string { return quote(s) }, s...)
}

// https://sqlite.org/lang_keywords.html
var (
	keywords = []string{
		"ABORT", "ACTION", "ADD", "AFTER", "ALL", "ALTER", "ALWAYS", "ANALYZE", "AND", "AS",
		"ASC", "ATTACH", "AUTOINCREMENT", "BEFORE", "BEGIN", "BETWEEN", "BY", "CASCADE",
		"CASE", "CAST", "CHECK", "COLLATE", "COLUMN", "COMMIT", "CONFLICT", "CONSTRAINT",
		"CREATE", "CROSS", "CURRENT", "CURRENT_DATE", "CURRENT_TIME", "CURRENT_TIMESTAMP",
		"DATABASE", "DEFAULT", "DEFERRABLE", "DEFERRED", "DELETE", "DESC", "DETACH", "DISTINCT",
		"DO", "DROP", "EACH", "ELSE", "END", "ESCAPE", "EXCEPT", "EXCLUDE", "EXCLUSIVE", "EXISTS",
		"EXPLAIN", "FAIL", "FILTER", "FIRST", "FOLLOWING", "FOR", "FOREIGN", "FROM", "FULL",
		"GENERATED", "GLOB", "GROUP", "GROUPS", "HAVING", "IF", "IGNORE", "IMMEDIATE", "IN",
		"INDEX", "INDEXED", "INITIALLY", "INNER", "INSERT", "INSTEAD", "INTERSECT", "INTO", "IS",
		"ISNULL", "JOIN", "KEY", "LAST", "LEFT", "LIKE", "LIMIT", "MATCH", "MATERIALIZED",
		"NATURAL", "NO", "NOT", "NOTHING", "NOTNULL", "NULL", "NULLS", "OF", "OFFSET", "ON", "OR",
		"ORDER", "OTHERS", "OUTER", "OVER", "PARTITION", "PLAN", "PRAGMA", "PRECEDING", "PRIMARY",
		"QUERY", "RAISE", "RANGE", "RECURSIVE", "REFERENCES", "REGEXP", "REINDEX", "RELEASE",
		"RENAME", "REPLACE", "RESTRICT", "RETURNING", "RIGHT", "ROLLBACK", "ROW", "ROWS",
		"SAVEPOINT", "SELECT", "SET", "TABLE", "TEMP", "TEMPORARY", "THEN", "TIES", "TO",
		"TRANSACTION", "TRIGGER", "UNBOUNDED", "UNION", "UNIQUE", "UPDATE", "USING", "VACUUM",
		"VALUES", "VIEW", "VIRTUAL", "WHEN", "WHERE", "WINDOW", "WITH", "WITHOUT",
	}
)

func init() {
	var lowercase []string

	for _, keyword := range keywords {
		lowercase = append(lowercase, stringer.Lowercase(keyword))
	}

	keywords = append(keywords, lowercase...)
}
