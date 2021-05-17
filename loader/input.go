package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t *Loader) processInput() error {
	// t.reportStack.Push(fmt.Sprintf("input[%v]", util.SummerySuffix(t.inputPath, model.SUMMERY_SIZE)))
	t.reportStack.Push("loader")
	defer t.reportStack.Pop()

	for _, inputPath := range t.inputPaths {
		var input model.InputFile

		if err := util.UnmarshalFromFile(inputPath, &input); err != nil {
			return err
		}

		for _, p := range util.SchemaTypePrimitives {
			t.primitiveMap[string(p)] = ""
			t.dependencies[string(p)] = NewNode(string(p))
		}

		for k, v := range input.Info.Primitives {
			if _, ok := t.dependencies[k]; !ok {
				return t.ErrInvalidPrimitive(k)
			}

			if p, ok := t.primitiveMap[k]; ok {
				if p != "" {
					return t.ErrDuplicatePrimitive(k)
				} else {
					t.primitiveMap[k] = v
				}
			} else {
				return t.ErrInvalidPrimitive(k)
			}
		}

		for _, m := range input.Specification.Models {
			if _, ok := t.inputModels[m.Name]; ok {
				return t.ErrDuplicateType(m.Name)
			} else {
				t.inputModels[m.Name] = m
			}
		}

		for _, o := range input.Specification.Operations {
			if _, ok := t.inputOperations[o.Name]; ok {
				return t.ErrorDuplicateOperation(o.Name)
			} else {
				t.inputOperations[o.Name] = o
			}
		}
	}

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

	if err := t.processModel5(); err != nil {
		return err
	}

	if err := t.processOperation1(); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.outputModels))
	// fmt.Println(util.Jdump(t.OutputSpec))
	return nil
}
