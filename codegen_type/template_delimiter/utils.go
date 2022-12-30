package template_delimiter

var (
	m = map[TemplateDelimiter][]string{
		Square: {"[[", "]]"},
		Curly:  {"{{", "}}"},
		Parens: {"((", "))"},
		Angle:  {"<<", ">>"},
	}
)

func Get(d TemplateDelimiter) (string, string) {
	c := m[d]
	return c[0], c[1]
}
