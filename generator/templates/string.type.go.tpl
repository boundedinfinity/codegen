package {{ typePackage . }}

import (
    "github.com/boundedinfinity/go-commoner/errorer"
    "regexp"
)

{{ if .HasValidation -}}
var (
    {{ if .Min.Defined -}}
    Err{{ typeName . }}Min = errorer.New("below min")
    {{- end }}
    {{ if .Max.Defined -}}
    Err{{ typeName . }}Max = errorer.New("above max")
    {{- end }}
    {{ if .Regex.Defined -}}
    Err{{ typeName . }}Regex = errorer.New(`does not match regex "{{ .Regex.Get -}}"`)
    {{- end }}
)
{{- end }}

type {{ typeName . }} {{ langType . }}

{{ if .Regex.Defined -}}
var (
    _{{ typeName . }}Regex = regexp.MustCompile(`{{ .Regex.Get }}`)
)
{{- end }}

func (s {{ typeName . }}) Validate() error {
    {{ if .Min.Defined -}}
    if len(s) < {{ .Min.Get }} {
        return Err{{ typeName . }}Min.WithValue(len(s))
    }
    {{- end }}

    {{ if .Max.Defined -}}
    if len(s) > {{ .Max.Get }} {
        return Err{{ typeName . }}Max.WithValue(len(s))
    }
    {{- end }}

    {{ if .Regex.Defined -}}
    if !_{{ typeName . }}Regex.MatchString(string(s)) {
        return Err{{ typeName . }}Regex.WithValue(s)
    }
    {{- end }}

    return nil
}
