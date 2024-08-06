package {{ .Lang.Package.Get }}{{ $tname := .Lang.Name.Get }}{{ $tnameLower := lowerFirst $tname }}

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

{{ if .Lang.Imports.Defined }}import (
{{ range .Lang.Imports.Get -}}
    "{{ . }}"
{{ end }}
){{ end }}

type {{ $tname }} int

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validation
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

{{ if .HasValidation -}}{{ $validators := printf "%vValidators" $tnameLower }}
var {{ $validators }} = []func({{ $tname }}) error{}


func (this {{ $tname }}) Validate() error {
    var errs []error

    {{ if .Min.Defined -}}errs = append(errs, validation.IntegerMin("{{ $tname }}", {{ .Min.Get }}, this)){{ end }}
    {{ if .Max.Defined -}}errs = append(errs, validation.IntegerMax("{{ $tname }}", {{ .Max.Get }}, this)){{ end }}
    {{ if .MultipleOf.Defined -}}errs = append(errs, validation.IntegerMultipleOf("{{ $tname }}", {{ .MultipleOf.Get }}, this)){{ end }}
    {{ if .Negative.Defined -}}errs = append(errs, validation.IntegerNegative("{{ $tname }}", this)){{ end }}
    {{ if .Negative.Defined -}}errs = append(errs, validation.IntegerNegative("{{ $tname }}", this)){{ end }}
    
	for _, validator := range {{ $tnameLower }}Validators {
		errs = append(errs, validator(this))
	}
	
    return errors.Join(errs...)
}
{{ end }}


// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// SQL
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this {{ $tname }}) Value() (driver.Value, error) {
    return int(this), nil
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

