package {{ typePackage . }}

import (
    {{ if .HasValidation -}}"github.com/boundedinfinity/go-commoner/errorer"{{ end }}
    {{ if .Regex.Defined }}"regexp"{{ end }}
    {{ if or .Includes.Defined .Excludes.Defined }}"github.com/boundedinfinity/go-commoner/idiomatic/stringer"{{ end }}
)

type {{ typeName . }} {{ langType . }}

{{ if .Regex.Defined }}var _{{ typeName . }}Regex = regexp.MustCompile(`{{ .Regex.Get }}`){{- end }}

{{ if .HasValidation -}}
var (
    {{ if .Min.Defined }}Err{{ typeName . }}Min = errorer.New("below min"){{ end }}
    {{ if .Max.Defined }}Err{{ typeName . }}Max = errorer.New("above max"){{ end }}
    {{ if .Regex.Defined }}Err{{ typeName . }}Regex = errorer.New(`does not match regex "{{ .Regex.Get -}}"`){{ end }}
    {{ if .Includes.Defined }}Err{{ typeName . }}NotIncluded = errorer.New(`none of the following included {{ sjoin .Includes.Get ", " }}`){{ end }}
    {{ if .Excludes.Defined }}Err{{ typeName . }}NotExcluded = errorer.New(`one or more included {{ sjoin .Excludes.Get ", " }}`){{ end }}
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

    {{ if .Includes.Defined -}}
    if !stringer.ContainsAny(string(s), {{ sjoin .Includes.Get ", " }}) {
        return Err{{ typeName . }}NotIncluded.WithValue(s)
    }
    {{- end }}

    {{ if .Excludes.Defined -}}
    if stringer.ContainsAny(string(s), {{ sjoin .Excludes.Get ", " }}) {
        return Err{{ typeName . }}NotExcluded.WithValue(s)
    }
    {{- end }}

    return nil
}
