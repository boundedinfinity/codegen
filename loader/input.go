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

	if err := t.processNamespace1(-1, t.input.Specification); err != nil {
		return err
	}

	if err := t.processNamespace2(-1, t.input.Specification); err != nil {
		return err
	}

	t.reportStack.Pop()
	return nil
}
