package loader

func (t *Loader) processInput_Specification() error {
	t.reportStack.Push("specification")

	if err := t.processInput_Specification_Model(); err != nil {
		return err
	}

	if err := t.processInput_Specification_Operation(); err != nil {
		return err
	}

	t.reportStack.Pop()
	return nil
}
