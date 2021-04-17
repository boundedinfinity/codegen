package loader

func (t *Loader) processInput() error {
	// t.reportStack.Push(fmt.Sprintf("input[%v]", util.SummerySuffix(t.inputPath, model.SUMMERY_SIZE)))
	t.reportStack.Push("loader")

	{
		t.reportStack.Push("name")
		t.Output.Name = t.input.Name
		// t.modelStack.Push(t.Output.Name)
		t.reportStack.Pop()
	}

	{
		t.reportStack.Push("version")
		// TODO
		t.reportStack.Pop()
	}

	if err := t.processInput_Info(); err != nil {
		return err
	}

	if err := t.processInput_Specification(); err != nil {
		return err
	}

	// if t.spec.Operations.Namespaces != nil {
	// 	for _, ns := range t.spec.Operations.Namespaces {
	// 		if err := t.processNamespace1(ns); err != nil {
	// 			return err
	// 		}
	// 	}

	// 	for _, sNamespace := range t.spec.Operations.Namespaces {
	// 		gNamespace, err := t.processNamespace2(sNamespace, t.spec.Operations.Templates)

	// 		if err != nil {
	// 			return err
	// 		}

	// 		t.Gen.Operations.Namespaces = append(t.Gen.Operations.Namespaces, gNamespace)
	// 	}
	// }

	t.reportStack.Pop()
	return nil
}
