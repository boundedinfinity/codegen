{{ define "float_type" }}
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

        {{ template "float_validation" dict "Name" $tname "Instance" . "Path" "this" }}
        
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
{{ end }}

{{ define "float_validation" }}
    {{ if eq .Instance.Schema "float" }}
        {{ if .Instance.Min.Defined -}}errs = append(errs, validation.IntegerMin("{{ .Name }}", {{ .Instance.Min.Get }}, {{ .Path }})){{ end }}
        {{ if .Instance.Max.Defined -}}errs = append(errs, validation.FloatMax("{{ .Name }}", {{ .Instance.Max.Get }}, {{ .Path }})){{ end }}
        {{ if .Instance.MultipleOf.Defined -}}errs = append(errs, validation.FloatMultipleOf("{{ .Name }}", {{ .Instance.MultipleOf.Get }}, {{ .Path }})){{ end }}
        {{ if .Instance.Negative.Defined -}}errs = append(errs, validation.FloatNegative("{{ .Name }}", {{ .Path }})){{ end }}
        {{ if .Instance.Negative.Defined -}}errs = append(errs, validation.FloatNegative("{{ .Name }}", {{ .Path }})){{ end }}
    {{ end }}
{{ end }}
