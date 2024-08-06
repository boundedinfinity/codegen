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
	Types       []CodeGenSchema               `json:"types,omitempty"`
	Package     optioner.Option[string]       `json:"package,omitempty"`
	OutputRoot  optioner.Option[string]       `json:"output-root,omitempty"`
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

func BuildProject() *codeGenProjectBuilder {
	return &codeGenProjectBuilder{}
}

type codeGenProjectBuilder struct {
	obj CodeGenProject
}

func (t *codeGenProjectBuilder) Build() *CodeGenProject {
	return &t.obj
}

func (t *codeGenProjectBuilder) Name(v string) *codeGenProjectBuilder {
	return SetO(t, &t.obj.Name, v)
}

func (t *codeGenProjectBuilder) Description(v string) *codeGenProjectBuilder {
	return SetO(t, &t.obj.Description, v)
}

func (t *codeGenProjectBuilder) Operations(v ...*CodeGenOperation) *codeGenProjectBuilder {
	for _, operation := range v {
		t.obj.Operations = append(t.obj.Operations, *operation)
	}

	return t
}

func (t *codeGenProjectBuilder) Types(v ...CodeGenSchema) *codeGenProjectBuilder {
	t.obj.Types = append(t.obj.Types, v...)
	return t
}
