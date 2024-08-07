{{ define "array_validation" }}
    {{ if eq .Instance.Schema "array" }}
        for i, item := range {{ .Path }} {
            name := fmt.Sprintf("{{ .Name }}%d", i)
            {{ template "ref_validation" dict "Name" "name" "Path" "item" "Instance" .Instance.Items.Get }}
            {{ template "string_validation" dict "Name" "name" "Path" "item" "Instance" .Instance.Items.Get }}
            {{ template "integer_validation" dict "Name" "name" "Path" "item" "Instance" .Instance.Items.Get }}
            {{ template "float_validation" dict "Name" "name" "Path" "item" "Instance" .Instance.Items.Get }}
            {{ template "array_validation" dict "Name" "name" "Path" "item" "Instance" .Instance.Items.Get }}
        }
    {{ end }}
{{ end }}
