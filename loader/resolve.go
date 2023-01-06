package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Loader) Resolve() error {
	err := ct.Walker().
		Type(func(project *ct.CodeGenProject, typ ct.CodeGenType) error {
			switch c := typ.(type) {
			case *ct.CodeGenTypeRef:
				resolved := t.typeManager.Resolve(typ)

				if resolved.Defined() {
					c.Resolved = resolved.Get()
				} else {
					ct.ErrCodeGenRefNotFoundv(typ)
				}
			}

			return nil
		}).
		Operation(func(project *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
			switch c := operation.Input.Get().(type) {
			case *ct.CodeGenTypeRef:
				resolved := t.typeManager.Resolve(operation.Input.Get())

				if resolved.Defined() {
					c.Resolved = resolved.Get()
				}
			}

			switch c := operation.Output.Get().(type) {
			case *ct.CodeGenTypeRef:
				resolved := t.typeManager.Resolve(operation.Output.Get())

				if resolved.Defined() {
					c.Resolved = resolved.Get()
				}
			}

			return nil
		}).
		Walk(t.projectManager.Merged)

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
