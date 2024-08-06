{{ define "string_type" }}
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

        {{ template "string_validation" dict "Name" $tname "Instance" . "Path" "this" }}

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
{{ end }}

{{ define "string_validation" }}
    {{ if eq .Instance.Schema "string" }}
        {{ if .Instance.Min.Defined -}}errs = append(errs, validation.StringMin("{{ .Name }}", {{ .Instance.Min.Get }}, {{ .Path }})){{ end }}
        {{ if .Instance.Max.Defined -}}errs = append(errs, validation.StringMax("{{ .Name }}", {{ .Instance.Max.Get }}, {{ .Path }})){{ end }}
        {{ if .Instance.Regex.Defined -}}errs = append(errs, validation.StringRegex("{{ .Name }}", `{{ .Instance.Regex.Get }}`, {{ .Path }})){{ end }}
    {{ end }}
{{ end }}
