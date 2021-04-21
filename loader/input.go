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
		t.report(node.name)

		if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessor4(node.name), nil); err != nil {
			return err
		}
	}

	if err := t.walk(-1, t.input.Specification, nil, nil, t.propertyProcessorJson, nil); err != nil {
		return err
	}

	if err := t.walk(-1, t.input.Specification, nil, t.modelProcessor5, nil, nil); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.modelMap))

	return nil
}

// RubiconStatus := NewNode("github.comcast.com/BusinessServices/mercury-client/model/RubiconStatus")
// Rubicon_Adtran_Msp_Neighbor := NewNode("github.comcast.com/BusinessServices/mercury-client/adtran/v2/Rubicon_Adtran_Msp_Neighbor")
// Rubicon_Adtran_Msp := NewNode("github.comcast.com/BusinessServices/mercury-client/adtran/v2/Rubicon_Adtran_Msp")
// Rubicon_Adtran_Msp.Add(RubiconStatus.name)
// Rubicon_Adtran_Msp.Add(Rubicon_Adtran_Msp_Neighbor.name)
// Rubicon_Broadsoft_Cdr_Filename_Store_Filename := NewNode("github.comcast.com/BusinessServices/mercury-client/bcd/v1/store/cdr/filename/Rubicon_Broadsoft_Cdr_Filename_Store_Filename")
// Rubicon_Broadsoft_Cdr_Filename_Store_Filename.Add(RubiconStatus.name)

// var workingGraph Graph
// // workingGraph = append(workingGraph, nodeA, nodeB, nodeC, nodeD, nodeE, nodeF, nodeG, nodeH, nodeI, nodeJ, nodeK)
// workingGraph = append(workingGraph, RubiconStatus, Rubicon_Adtran_Msp_Neighbor, Rubicon_Adtran_Msp, Rubicon_Broadsoft_Cdr_Filename_Store_Filename)
// fmt.Printf(">>> A working dependency graph\n")
// t.displayGraph(workingGraph)

// resolved, err := resolveGraph(workingGraph)
// if err != nil {
// 	fmt.Printf("Failed to resolve dependency graph: %s\n", err)
// } else {
// 	fmt.Println("The dependency graph resolved successfully")
// }
// for _, node := range resolved {
// 	fmt.Println(node.name)
// }
// fmt.Printf(">>> A working dependency graph\n")
// t.displayGraph(resolved)
