package {{ typePackage . }}{{ $tname := typeName . }}

{{ if .HasValidation -}}
import (
    {{ if .HasValidation -}}"boundedinfinity/codegen/validation"{{ end }}
    {{ if .HasValidation -}}"errors"{{ end }}
)
{{ end }}

{{ $tnameLower := lowerFirst $tname -}}{{ if .HasValidation -}}var {{ $tnameLower }}Validators = []func({{ $tname }}) error{}{{ end }}

{{ if .HasValidation -}}
func (s {{ $tname }}) Validate() error {
    var errs []error

	for _, validator := range {{ $tnameLower }}Validators {
		errs = append(errs, validator(s))
	}
	
    return errors.Join(errs...)
}
{{ end }}

{{ if .HasValidation -}}
func init() {
    {{ if .Min.Defined -}}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.StringMin[{{ $tname }}]("{{ $tname }}", {{ .Min.Get }})){{ end }}
    {{ if .Max.Defined -}}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.StringMax[{{ $tname }}]("{{ $tname }}", {{ .Max.Get }})){{ end }}
    {{ if .Regex.Defined -}}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.StringRegex[{{ $tname }}]("{{ $tname }}", `{{ .Regex.Get }}`)){{ end }}
}
{{ end }}
