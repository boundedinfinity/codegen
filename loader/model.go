package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t *Loader) processModel1() error {
	for n := range t.inputModels {
		_, ok := t.dependencies[n]

		if ok {
			return t.ErrDuplicateType(n)
		}

		t.dependencies[n] = NewNode(n)
	}

	return nil
}

func (t *Loader) processModel2() error {
	for n, m := range t.inputModels {
		mNode, ok := t.dependencies[n]

		if !ok {
			return t.ErrInvalidType(n)
		}

		if util.IsSchemaPrimitive(m) || util.IsSchemaRef(m) {
			dn := util.TrimSchemaArray(m)
			dNode, ok := t.dependencies[dn]

			if !ok {
				return t.ErrInvalidType(dn)
			}

			mNode.Add(dNode)
		} else if util.IsSchemaRecord(m) {
			for _, property := range m.Properties {
				if util.IsSchemaPrimitive(property) || util.IsSchemaRef(property) {
					dn := util.TrimSchemaArray(property)
					dNode, ok := t.dependencies[dn]

					if !ok {
						return t.ErrInvalidType(dn)
					}

					mNode.Add(dNode)
				}
			}
		} else if util.IsSchemaEnum(m) {
			// Nothing to do here
		} else {
			return t.ErrInvalidType(n)
		}
	}

	return nil
}

func (t *Loader) processModel3() error {
	t.reportf(t.reportStack.S(), "resolving model dependencies")

	var brokenGraph Graph

	for _, g := range t.dependencies {
		brokenGraph = append(brokenGraph, g)
	}

	solvedGraph, err := resolveGraph(brokenGraph)

	if err != nil {
		return err
	}

	t.solvedDependencies = solvedGraph

	return nil
}

func (t *Loader) processModel4() error {
	for _, node := range t.solvedDependencies {
		if model.IsSchemaTypeEnum(node.Name) {
			continue
		}

		t.reportf(t.reportStack.S(), "processing %v", node.Name)
		inputModel, ok := t.inputModels[node.Name]

		if !ok {
			return t.ErrInvalidType(node.Name)
		}

		outputModel := model.NewOutputModel()
		t.modelMap[node.Name] = outputModel
		outputModel.Name = node.Name
		outputModel.Description = t.splitDescription(inputModel.Description)

		if util.IsSchemaPrimitive(inputModel) {
			if err := t.processPrimitive(inputModel, outputModel); err != nil {
				return err
			}
		} else {
			fmt.Printf("missed %v\n", inputModel.Name)
		}
	}

	return nil
}

func (t *Loader) processPrimitive(inputModel model.InputModel, outputModel *model.OutputModel) error {
	isArray := util.IsSchemaArray(inputModel)
	typStr := util.TrimSchemaArray(inputModel)
	typ, err := model.SchemaTypeEnumParse(typStr)

	if err != nil {
		return err
	}

	switch typ {
	case model.SchemaType_String:
		if isArray {
			var v []string
			if err := json2Interface(inputModel.Example, &v, `"%v"`); err != nil {
				return err
			} else {
				outputModel.Example = v
			}
		} else {
			var v string
			if err := json2Interface(inputModel.Example, &v, `"%v"`); err != nil {
				return err
			} else {
				outputModel.Example = v
			}
		}
	case model.SchemaType_Int:
		if isArray {
			fmt.Print("processing int array")
		} else {
			fmt.Print("processing int")
		}
	case model.SchemaType_Double:
		if isArray {
			fmt.Print("processing double array")
		} else {
			fmt.Print("processing double")
		}
	default:
		fmt.Printf("missed primitive %v", typ)
		// t.ErrInvalidPrimitive(typ.String())
	}

	return nil
}
