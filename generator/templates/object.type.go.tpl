{{ define "object_type" }}
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

    // ////////////////////////////////////////////////////////////////////////////////////////////////////////
    // Validation
    // ////////////////////////////////////////////////////////////////////////////////////////////////////////
    
    {{ if .HasValidation -}}
        {{ $validators := printf "%sValidators" $tnameLower }}
        var {{ $validators }} = []func({{ $tname }}) error{}
        func (this {{ $tname }}) Validate() error {
            var errs []error

            {{ range .Properties.Get }}
                {{ $pname := printf "%s.%s" $tname .Lang.Name.Get }}
                {{ $path := printf "this.%s" .Lang.Name.Get }}

                {{ template "ref_validation" dict "Name" $pname "Path" $path "Instance" . }}
                {{ template "string_validation" dict "Name" $pname "Path" $path "Instance" . }}
                {{ template "integer_validation" dict "Name" $pname "Path" $path "Instance" . }}
            {{ end }}

            for _, validator := range {{ $validators }} {
                errs = append(errs, validator(this))
            }

            return errors.Join(errs...)
        }
    {{ end }}

    // ////////////////////////////////////////////////////////////////////////////////////////////////////////
    // SQL
    // ////////////////////////////////////////////////////////////////////////////////////////////////////////
    

{{ end }}
