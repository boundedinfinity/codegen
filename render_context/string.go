package render_context

import o "github.com/boundedinfinity/go-commoner/optioner"

type RenderContextString struct {
	RenderContextBase
	Min                   o.Option[int]
	Max                   o.Option[int]
	Regex                 o.Option[string]
	RegexErrorDescription o.Option[string]
	IntegerMin            o.Option[int]
	IntegerMax            o.Option[int]
	LetterMin             o.Option[int]
	LetterMax             o.Option[int]
	LowerCaseMin          o.Option[int]
	LowerCaseMax          o.Option[int]
	UpperCaseMin          o.Option[int]
	UpperCaseMax          o.Option[int]
	SymbolMin             o.Option[int]
	SymbolMax             o.Option[int]
}

func (t RenderContextString) HasValidation() bool {
	return t.Min.Defined() || t.Max.Defined() || t.Regex.Defined() ||
		t.IntegerMin.Defined() || t.IntegerMax.Defined() ||
		t.LetterMin.Defined() || t.LetterMax.Defined() ||
		t.LowerCaseMin.Defined() || t.LowerCaseMax.Defined() ||
		t.UpperCaseMin.Defined() || t.UpperCaseMax.Defined() ||
		t.SymbolMin.Defined() || t.SymbolMax.Defined()
}

var _ RenderContext = &RenderContextString{}
