{{ define "array_validation" }}
    {{ if eq .Instance.Schema "array" }}
        {{ $items := printf "%s..." .Path }}
        {{ template "ref_validation" dict "Name" .Name "Path" $items "Instance" .Instance.Items.Get }}
        {{ template "string_validation" dict "Name" .Name "Path" $items "Instance" .Instance.Items.Get }}
        {{ template "integer_validation" dict "Name" .Name "Path" $items "Instance" .Instance.Items.Get }}
        {{ template "float_validation" dict "Name" .Name "Path" $items "Instance" .Instance.Items.Get }}
        {{ template "array_validation" dict "Name" .Name "Path" $items "Instance" .Instance.Items.Get }}
        
    {{ end }}
{{ end }}
