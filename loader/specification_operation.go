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

			t.Output.Models.Namespaces = append(t.Output.Models.Namespaces, ons)
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

func (t *Loader) processInput_Specification_Operation_Namepace2(i int, ins model.BiInput_Specification_Namespace_Operation, tmpls []model.BiInput_Template_Info) (model.BiOutput_Model_Namespace, error) {
	t.modelStack.Push(ins.Name)

	gns := model.BiOutput_Model_Namespace{
		Name:       t.currentNamespace(),
		Models:     make([]model.BiOutput_Model, 0),
		Namespaces: make([]model.BiOutput_Model_Namespace, 0),
	}

	// if ins.Operations != nil {
	// 	for i, styp := range ins.Operations {
	// 		t.reportStack.Push(fmt.Sprintf("model[%v]", i))

	// 		gtyp, err := t.specType2genType(gns, styp)

	// 		if err != nil {
	// 			return gns, err
	// 		}

	// 		if styp.Properties != nil {
	// 			for i, sprop := range styp.Properties {
	// 				t.reportStack.Push(fmt.Sprintf("properties[%v]", i))

	// 				gprop, err := t.specProperty2genProperty(gns, gtyp, sprop)

	// 				if err != nil {
	// 					return gns, err
	// 				}

	// 				gtyp.Properties = append(gtyp.Properties, gprop)

	// 				t.reportStack.Pop()
	// 			}
	// 		}

	// 		if err := t.genTypeImports(&gtyp); err != nil {
	// 			return gns, err
	// 		}

	// 		if tmpls != nil {
	// 			for _, stmpl := range tmpls {
	// 				gtmpl, err := t.processTemplate(gtyp.Namespace, gtyp.Name, stmpl)

	// 				if err != nil {
	// 					return gns, err
	// 				}

	// 				gtyp.Templates = append(gtyp.Templates, gtmpl)
	// 			}
	// 		}

	// 		gns.Models = append(gns.Models, gtyp)

	// 		t.reportStack.Pop()
	// 	}
	// }

	// if ins.Namespaces != nil {
	// 	for i, cns := range ins.Namespaces {
	// 		if gen, err := t.processInput_Specification_Model_Namepace2(i, cns, tmpls); err != nil {
	// 			return gen, err
	// 		} else {
	// 			gns.Namespaces = append(gns.Namespaces, gen)
	// 		}
	// 	}
	// }

	t.modelStack.Pop()
	return gns, nil
}
