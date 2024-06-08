package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/mapper"
)

///////////////////////////////////////////////////////////////////
// Type
//////////////////////////////////////////////////////////////////

type CodeGenProject struct {
	Name        optioner.Option[string]       `json:"name,omitempty"`
	Description optioner.Option[string]       `json:"description,omitempty"`
	Mappings    mapper.Mapper[string, string] `json:"mappings,omitempty"`
	Operations  []CodeGenOperation            `json:"operations,omitempty"`
	Templates   CodeGenProjectTemplates       `json:"templates,omitempty"`
	Types       []CodeGenType                 `json:"types,omitempty"`
	CodeGenMeta
}

//----------------------------------------------------------------
// Marshal
//----------------------------------------------------------------

func (t *CodeGenProject) UnmarshalJSON(data []byte) error {
	dto := struct {
		Name        optioner.Option[string]       `json:"name,omitempty"`
		Description optioner.Option[string]       `json:"description,omitempty"`
		Mappings    mapper.Mapper[string, string] `json:"mappings,omitempty"`
		Operations  []CodeGenOperation            `json:"operations,omitempty"`
		Templates   CodeGenProjectTemplates       `json:"templates,omitempty"`
		Types       []json.RawMessage             `json:"types,omitempty"`
		CodeGenMeta
	}{}

	if err := json.Unmarshal(data, &dto); err != nil {
		return err
	} else {
		t.Name = dto.Name
		t.Description = dto.Description
		t.Mappings = dto.Mappings
		t.Operations = dto.Operations
		t.Templates = dto.Templates
		t.CodeGenMeta = dto.CodeGenMeta
	}

	for i, raw := range dto.Types {
		if typ, err := UnmarshalCodeGenType(raw); err != nil {
			return errors.Join(fmt.Errorf("type[%v]", i), err)
		} else {
			t.Types = append(t.Types, typ)
		}
	}

	return nil
}

//----------------------------------------------------------------
// Builders
//----------------------------------------------------------------

func NewProject() *CodeGenProject {
	return &CodeGenProject{}
}

func (t *CodeGenProject) WithName(v string) *CodeGenProject {
	t.Name = optioner.OfZero(v)
	return t
}

func (t *CodeGenProject) WithDescription(v string) *CodeGenProject {
	t.Description = optioner.OfZero(v)
	return t
}

func (t *CodeGenProject) WithOperations(v ...*CodeGenOperation) *CodeGenProject {
	for _, operation := range v {
		t.Operations = append(t.Operations, *operation)
	}

	return t
}

func (t *CodeGenProject) WithTypes(v ...CodeGenType) *CodeGenProject {
	t.Types = append(t.Types, v...)
	return t
}
