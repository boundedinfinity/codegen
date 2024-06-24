package {{ typePackage . }}

{{- $length := len .Ranges -}}

import (
    "github.com/boundedinfinity/go-commoner/errorer"
    "regexp"
)

type {{ typeName . }} {{ .BaseType }}

{{ if .HasValidation -}}
var (
    {{ if .MultipleOf.Defined -}}
    Err{{ typeName . }}MultipleOf = errorer.New("must be a multiple of")
    {{- end }}
    {{ if gt $length 0 -}}
    Err{{ typeName . }}Max = errorer.New("greater than max")
    Err{{ typeName . }}Min = errorer.New("less than min")
    {{- end }}
)
{{- end }}
