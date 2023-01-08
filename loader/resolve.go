package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) Resolve(projects ...*ct.CodeGenProject) error {
	err := ct.WalkType(func(_ *ct.CodeGenProject, typ ct.CodeGenType) error {
		return t.typeManager.ResolveRef(typ)
	}, projects...)

	if err != nil {
		return err
	}

	err = ct.WalkOperation(func(_ *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
		if operation.Input.Defined() {
			if err := t.typeManager.ResolveRef(operation.Input.Get()); err != nil {
				return err
			}
		}

		if operation.Output.Defined() {
			if err := t.typeManager.ResolveRef(operation.Output.Get()); err != nil {
				return err
			}
		}

		return nil
	}, projects...)

	if err != nil {
		return err
	}

	return nil
}

func (t *Loader) validateReferenceId(id optioner.Option[string]) error {
	if id.Empty() {
		// TODO
	}

	if !t.typeManager.Has(id.Get()) {
		return fmt.Errorf("reference not found %v", id.Get())
	}

	return nil
}

func (t *Loader) validateReference(typ ct.CodeGenType) error {
	switch c := typ.(type) {
	case *ct.CodeGenTypeArray:
		return t.validateReference(c.Items)
	case *ct.CodeGenTypeObject:
		for _, prop := range c.Properties {
			if err := t.validateReference(prop); err != nil {
				return err
			}
		}
	case *ct.CodeGenTypeRef:
		if c.Ref.Empty() {
			return fmt.Errorf("reference is empty")
		}

		if !t.typeManager.Has(c.Ref.Get()) {
			return fmt.Errorf("reference not found %v", c.Ref.Get())
		}
	}

	return nil
}
