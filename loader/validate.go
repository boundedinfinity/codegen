package loader

import (
	"boundedinfinity/codegen/canonical"
	"fmt"
)

func (t *Loader) Validate() error {
	for _, schema := range t.canonicals.All() {
		if err := t.validateReference(schema); err != nil {
			return err
		}
	}

	for _, operation := range t.mergedCodeGen.Operations {
		if operation.Input.Defined() {
			if !t.canonicals.Has(operation.Input.Get()) {
				return fmt.Errorf("reference not found %v", operation.Input.Get())
			}
		}

		if operation.Output.Defined() {
			if !t.canonicals.Has(operation.Output.Get()) {
				return fmt.Errorf("reference not found %v", operation.Output.Get())
			}
		}
	}

	return nil
}

func (t *Loader) validateReference(schema canonical.Canonical) error {
	switch c := schema.(type) {
	case canonical.CanonicalArray:
		return t.validateReference(c.Items)
	case canonical.CanonicalObject:
		for _, prop := range c.Properties {
			if err := t.validateReference(prop); err != nil {
				return err
			}
		}
	case canonical.CanonicalRef:
		if c.Ref.Empty() {
			return fmt.Errorf("reference is empty")
		}

		if !t.canonicals.Has(c.Ref.Get()) {
			return fmt.Errorf("reference not found %v", c.Ref.Get())
		}
	}

	return nil
}
