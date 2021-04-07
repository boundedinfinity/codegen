package loader

import "github.com/boundedinfinity/optional"

type OpenApiVersionSelector struct {
	Openapi optional.StringOptional `json:"openapi" yaml:"openapi"`
}
