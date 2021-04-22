package loader

import (
	"boundedinfinity/codegen/util"
	"fmt"
)

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

	if err := t.walk(-1, t.input.Specification, t.namespaceProcssor1, t.modelProcessor1, t.propertyProcessor1, nil); err != nil {
		return err
	}

	if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessor2, nil); err != nil {
		return err
	}

	if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessor3, nil); err != nil {
		return err
	}

	var brokenGraph Graph

	for _, g := range t.depNodes {
		brokenGraph = append(brokenGraph, g)
	}

	solvedGraph, err := resolveGraph(brokenGraph)

	if err != nil {
		return err
	}

	for _, node := range solvedGraph {
		t.report("processing node %v", node.name)

		if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessor4(node.name), nil); err != nil {
			return err
		}
	}

	if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessorJson, nil); err != nil {
		return err
	}

	if err := t.walk(-1, t.input.Specification, nil, nil, nil, t.processOperation6); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.modelMap))

	return nil
}
