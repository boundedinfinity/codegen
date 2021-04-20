package loader

func (t *Loader) processInput() error {
	// t.reportStack.Push(fmt.Sprintf("input[%v]", util.SummerySuffix(t.inputPath, model.SUMMERY_SIZE)))
	t.reportStack.Push("loader")
	defer t.reportStack.Pop()

	checkName := func() error {
		t.reportStack.Push("name")
		defer t.reportStack.Pop()

		t.Output.Name = t.input.Name
		return nil
	}

	checkVersion := func() error {
		t.reportStack.Push("version")
		defer t.reportStack.Pop()

		return nil
	}

	if err := checkName(); err != nil {
		return err
	}

	if err := checkVersion(); err != nil {
		return err
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

	// if err := t.processNamespace3(-1, t.input.Specification); err != nil {
	// 	return err
	// }

	return nil
}
