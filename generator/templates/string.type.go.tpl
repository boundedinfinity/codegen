package {{ .Lang.Package.Get }}{{ $tname := .Lang.Name.Get }}{{ $tnameLower := lowerFirst $tname }}

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

{{ if .Lang.Imports.Defined }}import (
{{ range .Lang.Imports.Get -}}
    "{{ . }}"
{{ end }}
){{ end }}

type {{ $tname }} string

// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// Validation
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

{{ if .HasValidation -}}{{ $validators := printf "%vValidators" $tnameLower }}
var {{ $validators }} = []func({{ $tname }}) error{}

func (this {{ $tname }}) Validate() error {
    var errs []error

    {{ if .Min.Defined -}}errs = append(errs, validation.StringMin("{{ $tname }}", {{ .Min.Get }}, this)){{ end }}
    {{ if .Max.Defined -}}errs = append(errs, validation.StringMax("{{ $tname }}", {{ .Max.Get }}, this)){{ end }}
    {{ if .Regex.Defined -}}errs = append(errs, validation.StringRegex("{{ $tname }}", `{{ .Regex.Get }}`, this)){{ end }}
    
	for _, validator := range {{ $validators }} {
		errs = append(errs, validator(this))
	}
	
    return errors.Join(errs...)
}
{{ end }}


// ////////////////////////////////////////////////////////////////////////////////////////////////////////
// SQL
// ////////////////////////////////////////////////////////////////////////////////////////////////////////

func (this {{ $tname }}) Value() (driver.Value, error) {
    return string(this), nil
}

func (this *{{ $tname }}) Scan(value interface{}) error {
	if value == nil {
		*this = {{ $tname }}("")
		return nil
	}

	if bv, err := driver.String.ConvertValue(value); err == nil {
		if v, ok := bv.(string); ok {
			*this = {{ $tname }}(v)
			return nil
		}
	}
	
	return errors.New("failed to scan {{ $tname }}")
}
