module boundedinfinity/codegen

go 1.21.5

require (
	github.com/boundedinfinity/go-commoner v1.0.36
	github.com/spf13/cobra v1.8.0
	github.com/stretchr/testify v1.9.0
)

require (
	github.com/boundedinfinity/go-mimetyper v1.0.18 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/invopop/yaml v0.3.1 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

// go mod edit -replace=github.com/boundedinfinity/go-commoner=../go-commoner

replace github.com/boundedinfinity/go-commoner => ../go-commoner
