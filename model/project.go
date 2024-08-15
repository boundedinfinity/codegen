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

func (this *CodeGenProject) Validate() error {
	if this.Name.Empty() {
		return ErrCodeGenProjectEmptyName
	}

	for k, v := range this.Mappings {
		if k == "" || v == "" {
			return ErrCodeGenProjectInvalidMappingFn(k, v)
		}
	}

	for i, operation := range this.Operations {
		if err := operation.Validate(); err != nil {
			return fmt.Errorf("operation[%v] %w", i, err)
		}
	}

	if err := this.Templates.Validate(); err != nil {
		return err
	}

	for i, typ := range this.Types {
		if err := typ.Validate(); err != nil {
			return fmt.Errorf("type[%v] %w", i, err)
		}
	}

	return nil
}

//----------------------------------------------------------------
// Merge
//----------------------------------------------------------------

func (this *CodeGenProject) Merge(obj CodeGenProject) error {
	this.Name = obj.Name
	this.Description = mergeDescription(this.Description, obj.Description)
	mapper.MergeInto(this.Mappings, obj.Mappings)

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

func (this *CodeGenProject) UnmarshalJSON(data []byte) error {
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
		this.Name = dto.Name
		this.Description = dto.Description
		this.Mappings = dto.Mappings
		this.Operations = dto.Operations
		this.Templates = dto.Templates
		this.Package = dto.Package
	}

	for i, raw := range dto.Types {
		if typ, err := UnmarshalCodeGenType(raw); err != nil {
			return errors.Join(fmt.Errorf("type[%v]", i), err)
		} else {
			this.Types = append(this.Types, typ)
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

func (this *codeGenProjectBuilder) Build() *CodeGenProject {
	return &this.obj
}

func (this *codeGenProjectBuilder) Name(v string) *codeGenProjectBuilder {
	return SetO(this, &this.obj.Name, v)
}

func (this *codeGenProjectBuilder) Description(v string) *codeGenProjectBuilder {
	return SetO(this, &this.obj.Description, v)
}

func (this *codeGenProjectBuilder) Operations(v ...*CodeGenOperation) *codeGenProjectBuilder {
	for _, operation := range v {
		this.obj.Operations = append(this.obj.Operations, *operation)
	}

	return this
}

func (this *codeGenProjectBuilder) Types(v ...CodeGenSchema) *codeGenProjectBuilder {
	this.obj.Types = append(this.obj.Types, v...)
	return this
}
