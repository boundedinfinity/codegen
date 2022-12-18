package codegen_project

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

type CodeGenProjectTypeFile struct {
	Path o.Option[string] `json:"path,omitempty" yaml:"path,omitempty"`
	Root o.Option[string] `json:"root,omitempty" yaml:"root,omitempty"`
}
