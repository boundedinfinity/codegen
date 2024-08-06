{{ define "ref_validation" }}
    {{ if eq .Instance.Schema "ref" }}
        errs = append(errs, {{ .Path }}.Validate())
    {{ end }}
{{ end }}
