package loader

import (
	"boundedinfinity/codegen/model"
	"fmt"
	"path"
)

func (t *Loader) processInput_Specification_Operation() error {
	t.reportStack.Push("operation")

	if t.input.Specification.Operation.Namespaces != nil {
		for i, ns := range t.input.Specification.Operation.Namespaces {
			if err := t.processInput_Specification_Operation_Namepace1(i, ns); err != nil {
				return nil
			}
		}
	}

	for i, ins := range t.input.Specification.Operation.Namespaces {
		if t.input.Template.Operations.Operation != nil {
			ons, err := t.processInput_Specification_Operation_Namepace2(i, ins, t.input.Template.Operations.Operation)

			if err != nil {
				return nil
			}

			t.Output.Operations.Namespaces = append(t.Output.Operations.Namespaces, ons)
		}
	}

	t.reportStack.Pop()
	return nil
}

func (t *Loader) processInput_Specification_Operation_Namepace1(i int, ins model.BiInput_Specification_Namespace_Operation) error {
	t.reportStack.Push(fmt.Sprintf("namespaces[%v]", i))
	t.modelStack.Push(ins.Name)

	if ins.Operations != nil {
		for i, o := range ins.Operations {
			t.reportStack.Push(fmt.Sprintf("operation[%v]", i))

			var k string
			var v string

			k = t.currentOperationName(o)
			v = t.currentNamespace()
			v = path.Base(v)
			v = fmt.Sprintf("%v.%v", v, path.Base(k))
			t.addMappedType(k, TypeInfo{
				InNamespace:  path.Base(k),
				OutNamespace: v,
				Namespace:    t.currentNamespace(),
			})

			t.reportStack.Pop()
		}
	}

	if ins.Namespaces != nil {
		for j, cns := range ins.Namespaces {
			if err := t.processInput_Specification_Operation_Namepace1(j, cns); err != nil {
				return err
			}
		}
	}

	t.modelStack.Pop()
	t.reportStack.Pop()
	return nil
}

func (t *Loader) processInput_Specification_Operation_Namepace2(i int, ins model.BiInput_Specification_Namespace_Operation, tmpls []model.BiInput_Template_Info) (model.BiOutput_Operation_Namespace, error) {
	t.modelStack.Push(ins.Name)

	gns := model.BiOutput_Operation_Namespace{
		Name:       t.currentNamespace(),
		Operations: make([]model.BiOutput_Operation, 0),
		Namespaces: make([]model.BiOutput_Operation_Namespace, 0),
	}

	if ins.Operations != nil {
		for i, sop := range ins.Operations {
			t.reportStack.Push(fmt.Sprintf("operation[%v]", i))

			gop, err := t.specOperation2genOperation(gns, sop)

			if err != nil {
				return gns, err
			}

			if sop.Inputs != nil {
				for i, sprop := range sop.Inputs {
					t.reportStack.Push(fmt.Sprintf("inputs[%v]", i))

					gprop, err := t.specIO2genProperty(gns, gop, sprop)

					if err != nil {
						return gns, err
					}

					gop.Inputs = append(gop.Inputs, gprop)

					t.reportStack.Pop()
				}
			}

			if sop.Outputs != nil {
				for i, sprop := range sop.Outputs {
					t.reportStack.Push(fmt.Sprintf("outputs[%v]", i))

					gprop, err := t.specIO2genProperty(gns, gop, sprop)

					if err != nil {
						return gns, err
					}

					gop.Outputs = append(gop.Outputs, gprop)

					t.reportStack.Pop()
				}
			}

			if err := t.genOperationImports(&gop); err != nil {
				return gns, err
			}

			if tmpls != nil {
				for _, stmpl := range tmpls {
					gtmpl, err := t.processTemplate(gop.Namespace, gop.Name, stmpl)

					if err != nil {
						return gns, err
					}

					gop.Templates = append(gop.Templates, gtmpl)
				}
			}

			gns.Operations = append(gns.Operations, gop)
			t.reportStack.Pop()
		}
	}

	if ins.Namespaces != nil {
		for i, cns := range ins.Namespaces {
			if gen, err := t.processInput_Specification_Operation_Namepace2(i, cns, tmpls); err != nil {
				return gen, err
			} else {
				gns.Namespaces = append(gns.Namespaces, gen)
			}
		}
	}

	t.modelStack.Pop()
	return gns, nil
}
