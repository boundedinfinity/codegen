package {{ typePackage . }}{{ $tname := typeName . }}

{{ if .HasValidation -}}
import (
    {{ if .HasValidation -}}"boundedinfinity/codegen/validation"{{ end }}
    {{ if .HasValidation -}}"errors"{{ end }}
)
{{ end }}

{{ $tnameLower := lowerFirst $tname -}}{{ if .HasValidation -}}var {{ $tnameLower }}Validators = []func({{ $tname }}) error{}{{ end }}

{{ if .HasValidation -}}
func (t {{ $tname }}) Validate() error {
    var errs []error

	for _, validator := range {{ $tnameLower }}Validators {
		errs = append(errs, validator(t))
	}
	
    return errors.Join(errs...)
}
{{ end }}

{{ if .HasValidation -}}
func init() {
    {{ if .Min.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerMin[{{ $tname }}]("{{ $tname }}", {{ .Min.Get }})){{ end }}
    {{ if .Max.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerMax[{{ $tname }}]("{{ $tname }}", {{ .Max.Get }})){{ end }}
    {{ if .MultipleOf.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerMultipleOf[{{ $tname }}]("{{ $tname }}", {{ .MultipleOf.Get }})){{ end }}
    {{ if .Positive.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerPositive[{{ $tname }}]("{{ $tname }}")){{ end }}
    {{ if .Negative.Defined }}{{ $tnameLower }}Validators = append({{ $tnameLower }}Validators, validation.IntegerNegative[{{ $tname }}]("{{ $tname }}")){{ end }}
}
{{ end }}
