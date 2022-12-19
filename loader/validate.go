package loader

import (
	cp "boundedinfinity/codegen/codegen_type"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) Validate() error {
	for _, lc := range t.typeManager.All() {
		if err := t.validateReference(lc.Schema); err != nil {
			return err
		}
	}

	for _, operation := range t.projectManager.Merged.Operations {
		if err := t.validateReferenceId(operation.Input); err != nil {
			return err
		}

		if err := t.validateReferenceId(operation.Output); err != nil {
			return err
		}
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

func (t *Loader) validateReference(schema cp.CodeGenType) error {
	switch c := schema.(type) {
	case *cp.CodeGenTypeArray:
		return t.validateReference(c.Items)
	case *cp.CodeGenTypeObject:
		for _, prop := range c.Properties {
			if err := t.validateReference(prop); err != nil {
				return err
			}
		}
	case *cp.CodeGenTypeRef:
		if c.Ref.Empty() {
			return fmt.Errorf("reference is empty")
		}

		if !t.typeManager.Has(c.Ref.Get()) {
			return fmt.Errorf("reference not found %v", c.Ref.Get())
		}
	}

	return nil
}
