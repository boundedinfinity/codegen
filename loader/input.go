package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t *Loader) processType(input model.InputModel, output *model.OutputModel) error {
	t.reportStack.Push("type")
	defer t.reportStack.Pop()

	// spath := t.appendNamespace()

	return nil
}

func (t *Loader) processInput() error {
	// input := t.inputSpec.Type
	// var output *model.OutputModel
	// output = model.NewOutputModel()

	// if err := t.processType(input, output); err != nil {
	// 	return err
	// }

	// t.reportStack.Push(fmt.Sprintf("input[%v]", util.SummerySuffix(t.inputPath, model.SUMMERY_SIZE)))
	t.reportStack.Push("loader")
	defer t.reportStack.Pop()

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

	if err := t.walkSpec(t.processNamespace1, WALKTYPE_NAMESPACE); err != nil {
		return err
	}

	if err := t.walkSpec(t.processModel1, WALKTYPE_MODEL); err != nil {
		return err
	}

	if err := t.walkSpec(t.processProperty1, WALKTYPE_PROPERTY); err != nil {
		return err
	}

	if err := t.processExamples(); err != nil {
		return err
	}

	if err := t.walkSpec(t.processProperty2, WALKTYPE_PROPERTY); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.OutputSpec))
	return nil
}

func (t *Loader) processExamples() error {
	t.reportf(t.reportStack.S(), "resolving model dependencies")

	var brokenGraph Graph

	for _, g := range t.dependencies {
		brokenGraph = append(brokenGraph, g)
	}

	solvedGraph, err := resolveGraph(brokenGraph)

	if err != nil {
		return err
	}

	for _, node := range solvedGraph {
		t.reportf(t.reportStack.S(), "processing node %v", node.name)

		if err := t.walkSpec(t.processModel2(node.name), WALKTYPE_MODEL); err != nil {
			return err
		}
	}

	return nil
}
