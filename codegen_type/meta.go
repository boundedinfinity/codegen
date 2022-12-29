package codegen_type

import o "github.com/boundedinfinity/go-commoner/optioner"

type CodeGenTypeMeta struct {
	//Since included at this specific version.  Fist instance should be the same as the Version.
	Since o.Option[string] `json:"since,omitempty"`

	//Version is the current version of the thing
	Version o.Option[string] `json:"version,omitempty"`
}

type CodeGenTypeDeprecated struct {
	//Since Deprecated since a specific version
	Since o.Option[string] `json:"since,omitempty"`

	//Removal Planned to be removed at a specific version
	Removal o.Option[string] `json:"removal,omitempty"`

	//Name of an alternative Planned to be removed at a specific version
	Alternative o.Option[string] `json:"alternative,omitempty"`
}
