package loader

import (
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t *Loader) processInput() error {
	// t.reportStack.Push(fmt.Sprintf("input[%v]", util.SummerySuffix(t.inputPath, model.SUMMERY_SIZE)))
	t.reportStack.Push("loader")
	defer t.reportStack.Pop()

	// t.reportStack.Push(`"%v"`, filepath.Base(t.inputPath))
	// defer t.reportStack.Pop()

	// checkName := func() error {
	// 	t.reportStack.Push("name")
	// 	defer t.reportStack.Pop()

	// 	t.Output.Name = t.input.Name
	// 	return nil
	// }

	// checkVersion := func() error {
	// 	t.reportStack.Push("version")
	// 	defer t.reportStack.Pop()

	// 	return nil
	// }

	// if err := checkName(); err != nil {
	// 	return err
	// }

	// if err := checkVersion(); err != nil {
	// 	return err
	// }

	// if err := t.processInput_Info(); err != nil {
	// 	return err
	// }

	if err := t.processModel1(); err != nil {
		return err
	}

	if err := t.processModel2(); err != nil {
		return err
	}

	if err := t.processModel3(); err != nil {
		return err
	}

	if err := t.processModel4(); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.OutputSpec))
	return nil
}
