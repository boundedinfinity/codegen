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

		if input.Info.OutputDir != "" {
			t.OutputSpec.Info.OutputDir = input.Info.OutputDir
		}

		if t.OutputSpec.Info.Namespace == "" && input.Info.Namespace == "" {
			t.OutputSpec.Info.Namespace = input.Info.Namespace
		} else {
			t.Errorf("namespace already defined")
		}

		t.OutputSpec.Info.DumpContext = input.Info.DumpContext

		if len(t.OutputSpec.Info.Primitives) <= 0 {
			t.OutputSpec.Info.Primitives = make(map[string]string)
			for _, p := range util.SchemaTypePrimitives {
				t.OutputSpec.Info.Primitives[string(p)] = ""
				t.dependencies[string(p)] = NewNode(string(p))
			}
		}

		for k, v := range input.Info.Primitives {
			if _, ok := t.dependencies[k]; !ok {
				return t.ErrInvalidPrimitive(k)
			}

			if p, ok := t.OutputSpec.Info.Primitives[k]; ok {
				if p != "" {
					return t.ErrDuplicatePrimitive(k)
				} else {
					t.OutputSpec.Info.Primitives[k] = v
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

		for _, inputTemplate := range input.Specification.Templates {
			t.appendInfoTemplate(inputTemplate)
		}
	}

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

	if err := t.processTemplate1(); err != nil {
		return err
	}

	if err := t.processTemplate2(); err != nil {
		return err
	}

	if err := t.processTemplate3(); err != nil {
		return err
	}

	if err := t.processTemplate4(); err != nil {
		return err
	}

	fmt.Println(util.Jdump(t.OutputSpec))

	t.OutputSpec.ModelMap = t.outputModels
	return nil
}
