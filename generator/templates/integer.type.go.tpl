package {{ typePackage . }}{{ $length := len .Ranges }}

import (
    "github.com/boundedinfinity/go-commoner/errorer"
)

{{ if .HasValidation -}}
var (
    {{ if .MultipleOf.Defined -}}
    Err{{ typeName . }}MultipleOf = errorer.New("must be a multiple of")
    {{- end }}
    {{ if gt $length 0 -}}
    Err{{ typeName . }}Range = errorer.New("not in range")
    {{- end }}
)
{{- end }}

type {{ typeName . }} {{ langType . }}

func (i {{ typeName . }}) Validate() error {
    {{ if .MultipleOf.Defined -}}
    if i % {{ .MultipleOf.Get }} != 0 {
        return Err{{ typeName . }}MultipleOf.WithValue(i)
    }
    {{- end }}

    {{ if gt $length 0 -}}
    var inRange bool
    

    {{ range .Ranges -}}
    if i {{ if .Min.Defined }}>= {{ .Min.Get }}{{ end }}{{ if .ExclusiveMin.Defined }}> {{ .ExclusiveMin.Get }}{{ end }} && i {{ if .Max.Defined }}<= {{ .Max.Get }}{{ end }}{{ if .ExclusiveMax.Defined }}< {{ .ExclusiveMax.Get }}{{ end }} {
        inRange = true
    }
    {{ end }}

    if !inRange  {
        return Err{{ typeName . }}Range
    }
    {{- end}}

    return nil
}
