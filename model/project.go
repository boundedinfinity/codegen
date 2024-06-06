package model

import (
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
	Operations  []CodeGenProject              `json:"operations,omitempty"`
	Templates   CodeGenProjectTemplates       `json:"templates,omitempty"`
	Types       []CodeGenType                 `json:"types,omitempty"`
	Source      optioner.Option[string]       `json:"source,omitempty"`
	LocalSource optioner.Option[string]       `json:"local-source,omitempty"`
}

//////////////////////////////////////////////////////////////////
// Builders
//////////////////////////////////////////////////////////////////

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

func (t *CodeGenProject) WithOperations(v ...CodeGenProject) *CodeGenProject {
	t.Operations = append(t.Operations, v...)
	return t
}

func (t *CodeGenProject) WithTypes(v ...CodeGenType) *CodeGenProject {
	t.Types = append(t.Types, v...)
	return t
}
