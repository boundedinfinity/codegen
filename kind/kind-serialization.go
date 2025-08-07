package kind

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
)

type KindSerializationConfig struct {
	Name string         `json:"name,omitempty,omitzero"`
	Case caser.CaseType `json:"case,omitempty,omitzero"`
}

func (this KindSerializationConfig) String() string {
	return fmt.Sprintf("name: %s, case: %s", this.Name, this.Case)
}
