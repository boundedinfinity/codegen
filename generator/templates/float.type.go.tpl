package {{ .Lang.Package.Get }}{{ $tname := .Lang.Name.Get }}{{ $tnameLower := lowerFirst $tname }}

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

{{ if .Lang.Imports.Defined }}import (
{{ range .Lang.Imports.Get -}}
    "{{ . }}"
{{ end }}
){{ end }}

type {{ $tname }} float64

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validation
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

{{ if .HasValidation -}}
var {{ $tnameLower }}Validators = []func({{ $tname }}) error{}{{ end }}


{{ if .HasValidation -}}
func (this {{ $tname }}) Validate() error {
    var errs []error

	for _, validator := range {{ $tnameLower }}Validators {
		errs = append(errs, validator(this))
	}
	
    return errors.Join(errs...)
}

{{ if .HasValidation -}}
func init() {
    {{ if .Min.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.FloatMin[{{ $tname }}]("{{ $tname }}", {{ .Min.Get }})){{ end }}
    {{ if .Max.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.FloatMax[{{ $tname }}]("{{ $tname }}", {{ .Max.Get }})){{ end }}
    // {{ if .MultipleOf.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerMultipleOf[{{ $tname }}]("{{ $tname }}", {{ .MultipleOf.Get }})){{ end }}
    {{ if .Positive.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.FloatPositive[{{ $tname }}]("{{ $tname }}")){{ end }}
    {{ if .Negative.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.FloatNegative[{{ $tname }}]("{{ $tname }}")){{ end }}
}
{{ end }}
{{ end }}


// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// SQL
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this {{ $tname }}) Value() (driver.Value, error) {
    return float64(this), nil
}

func (this *{{ $tname }}) Scan(value interface{}) error {
	if value == nil {
		*this = {{ $tname }}(0)
		return nil
	}

	if bv, err := driver.Int32.ConvertValue(value); err == nil {
		if v, ok := bv.({{ .Lang.Type }}); ok {
			*this = {{ $tname }}(v)
			return nil
		}
	}
	
	return errors.New("failed to scan {{ $tname }}")
}

