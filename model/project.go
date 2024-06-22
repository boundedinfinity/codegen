package model

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/boundedinfinity/go-commoner/errorer"
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
	Package     optioner.Option[string]       `json:"package,omitempty"`
	OutputRoot  optioner.Option[string]       `json:"output-root,omitempty"`
	CodeGenMeta
}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

var (
	ErrCodeGenProjectEmptyName        = errorer.New("empty name")
	ErrCodeGenProjectInvalidMapping   = errorer.New("invalid mapping")
	ErrCodeGenProjectInvalidMappingFn = ErrCodeGenProjectInvalidMapping.FormatFn("key: %v, val: %v, %w")
)

func (t *CodeGenProject) Validate() error {
	if t.Name.Empty() {
		return ErrCodeGenProjectEmptyName
	}

	for k, v := range t.Mappings {
		if k == "" || v == "" {
			return ErrCodeGenProjectInvalidMappingFn(k, v)
		}
	}

	for i, operation := range t.Operations {
		if err := operation.Validate(); err != nil {
			return fmt.Errorf("operation[%v] %w", i, err)
		}
	}

	if err := t.Templates.Validate(); err != nil {
		return err
	}

	for i, typ := range t.Types {
		if err := typ.Validate(); err != nil {
			return fmt.Errorf("type[%v] %w", i, err)
		}
	}

	return nil
}

//----------------------------------------------------------------
// Merge
//----------------------------------------------------------------

func (t *CodeGenProject) Merge(obj CodeGenProject) error {
	t.Name = obj.Name
	t.Description = mergeDescription(t.Description, obj.Description)
	mapper.MergeInto(t.Mappings, obj.Mappings)

	// operationG := slicer.Group(func(_ int, operation CodeGenOperation) string {
	// 	return operation.Name.Get()
	// }, t.Operations...)

	// for _, operation := range obj.Operations {
	// 	if found, ok := operationG[operation.Name.Get()]; ok {
	// 		if err := found.Merge(operation); err != nil {
	// 			return err
	// 		}
	// 	} else {
	// 		t.Operations = append(t.Operations, operation)
	// 	}
	// }

	// t.CodeGenMeta.Merge(obj.CodeGenMeta)
	return nil
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
		Package     optioner.Option[string]       `json:"package,omitempty"`
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
		t.Package = dto.Package
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
