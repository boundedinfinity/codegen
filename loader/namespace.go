package loader

import "boundedinfinity/codegen/model"

func (t *Loader) processNamespace1(si int, ns model.BiInput_Namespace) error {
	if si < 0 {
		t.reportStack.Push("specification")
	} else {
		t.reportStack.Push("namespace[%v]", si)
	}

	t.modelStack.Push(ns.Name)

	if ns.Models != nil {
		for i, m := range ns.Models {
			if err := t.processModel1(i, m); err != nil {
				return err
			}
		}
	}

	if ns.Operations != nil {
		for i, o := range ns.Operations {
			if err := t.processOperation1(i, o); err != nil {
				return err
			}
		}
	}

	if ns.Templates != nil {
		for i, tmpl := range ns.Templates {
			if err := t.processTemplate1(i, tmpl); err != nil {
				return err
			}
		}
	}

	if ns.Namespaces != nil {
		for i, child := range ns.Namespaces {
			if err := t.processNamespace1(i, child); err != nil {
				return err
			}
		}
	}

	t.modelStack.Pop()
	t.reportStack.Pop()
	return nil
}

func (t *Loader) processNamespace2(si int, ns model.BiInput_Namespace) error {
	if si < 0 {
		t.reportStack.Push("specification")
	} else {
		t.reportStack.Push("namespace[%v]", si)
	}

	t.modelStack.Push(ns.Name)

	if ns.Models != nil {
		for i, input := range ns.Models {
			output, err := t.processModel2(i, input)

			if err != nil {
				return err
			}

			t.Output.Models = append(t.Output.Models, output)
		}
	}

	if ns.Operations != nil {
		for i, input := range ns.Operations {
			output, err := t.processOperation2(i, input)

			if err != nil {
				return err
			}

			t.Output.Operations = append(t.Output.Operations, output)
		}
	}

	if ns.Name != "" {
		qns := t.currentNamespace()
		tmpls, err := t.getTemplates(qns, model.TemplateType_NAMESPACE)

		if err != nil {
			return err
		}

		nstmpls := make([]model.BiOutput_Template, 0)

		for _, itmpl := range tmpls {
			otmpl, err := t.processTemplate2(qns, "", itmpl)

			if err != nil {
				return err
			}

			nstmpls = append(nstmpls, otmpl)
		}

		t.Output.Namespaces = append(t.Output.Namespaces, model.BiOutput_Namespace{
			Namespace: qns,
			Templates: nstmpls,
		})
	}

	if ns.Namespaces != nil {
		for i, child := range ns.Namespaces {
			if err := t.processNamespace2(i, child); err != nil {
				return err
			}
		}
	}

	t.modelStack.Pop()
	t.reportStack.Pop()
	return nil
}
