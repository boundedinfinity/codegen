package loader

import (
	cp "boundedinfinity/codegen/codegen_type"
	"fmt"
)

func (t *Loader) Validate() error {
	for _, lc := range t.typeManager.All() {
		if err := t.validateReference(lc.Schema); err != nil {
			return err
		}
	}

	for _, operation := range t.projectManager.Merged.Operations {
		if operation.Input.Defined() {
			if !t.typeManager.Has(operation.Input.Get()) {
				return fmt.Errorf("reference not found %v", operation.Input.Get())
			}
		}

		if operation.Output.Defined() {
			if !t.typeManager.Has(operation.Output.Get()) {
				return fmt.Errorf("reference not found %v", operation.Output.Get())
			}
		}
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
