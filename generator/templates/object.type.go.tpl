package {{ .Lang.Package.Get }}{{ $tname := .Lang.Name.Get }}{{ $tnameLower := lowerFirst $tname }}

//lint:file-ignore ST1006
// https://staticcheck.dev/docs/checks#ST1006

{{ if .Lang.Imports.Defined }}
import (
{{ range .Lang.Imports.Get}}
    "{{ . }}"
{{ end }}
)
{{ end }}


type {{ $tname }} struct {
{{ if .Properties.Defined -}}
{{ range .Properties.Get }}
    {{ .Lang.Name.Get }} {{ .Lang.Type.Get }}
{{ end }}
{{ end }}
}

{{ if .HasValidation -}}
{{ $validators := printf "%vValidators" $tnameLower }}var {{ $validators }} = []func({{ $tname }}) error{}

func (this {{ $tname }}) Validate() error {
    var errs []error

{{ range .Properties.Get }}{{ $pname := .Lang.Name.Get }}

{{ if eq .Schema "ref" }}errs = append(errs, this.{{ .Lang.Name.Get }}.Validate()){{ end }}

{{ if eq .Schema "string" }}
    {{ if .Min.Defined -}}errs = append(errs, validation.StringMin("{{ $tname }}.{{ $pname }}", {{ .Min.Get }}, this.{{ $pname }})){{ end }}
    {{ if .Max.Defined -}}errs = append(errs, validation.StringMax("{{ $tname }}.{{ $pname }}", {{ .Max.Get }}, this.{{ $pname }})){{ end }}
    {{ if .Regex.Defined -}}errs = append(errs, validation.StringRegex("{{ $tname }}.{{ $pname }}", `{{ .Regex.Get }}`, this.{{ $pname }})){{ end }}
{{ end }}

{{ end }}

    for _, validator := range {{ $validators }} {
        errs = append(errs, validator(this))
    }

    return errors.Join(errs...)
}

{{ end }}
