module boundedinfinity/codegen

go 1.18

require (
	github.com/boundedinfinity/go-commoner v1.0.17
	github.com/boundedinfinity/go-jsonschema v1.0.2
	github.com/boundedinfinity/go-marshaler v1.0.5
	github.com/boundedinfinity/go-mimetyper v1.0.13
	github.com/boundedinfinity/go-urischemer v1.0.1
	github.com/ghodss/yaml v1.0.0
	github.com/spf13/cobra v1.5.0
	github.com/stretchr/testify v1.8.0
	gopkg.in/yaml.v2 v2.4.0
)

require (
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

require (
	github.com/boundedinfinity/mimetyper v1.0.10 // indirect
	github.com/inconshreveable/mousetrap v1.0.1 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/boundedinfinity/go-jsonschema => ../go-jsonschema

replace github.com/boundedinfinity/go-commoner => ../go-commoner

replace github.com/boundedinfinity/go-mimetyper => ../go-mimetyper

replace github.com/boundedinfinity/go-marshaler => ../go-marshaler

replace github.com/boundedinfinity/go-urischemer => ../go-urischemer
