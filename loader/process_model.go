package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path"

	"github.com/iancoleman/orderedmap"
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
	for _, m := range t.inputModels {
		if err := t.processModelDep(nil, &m); err != nil {
			return err
		}
	}

	return nil
}

func (t *Loader) processModelDep(p, m *model.InputModel) error {
	addDep := func(depender, dependee string) error {
		dependerNode, ok := t.dependencies[depender]

		if !ok {
			return t.ErrCustomTypeNotFound(depender)
		}

		dependeeNode, ok := t.dependencies[dependee]

		if !ok {
			return t.ErrCustomTypeNotFound(dependee)
		}

		dependerNode.Add(dependeeNode)
		return nil
	}

	fn := func(a, b *model.InputModel) string {
		if a != nil && a.Name != "" {
			return a.Name
		}

		if b != nil && b.Name != "" {
			return b.Name
		}

		return "turd"
	}

	switch m.Type {
	case model.SchemaType_String, model.SchemaType_Boolean,
		model.SchemaType_Int, model.SchemaType_Long,
		model.SchemaType_Float, model.SchemaType_Double:
		dependee := fn(p, m)
		if err := addDep(dependee, m.Type.String()); err != nil {
			return err
		}
	case model.SchemaType_Array:
		if p != nil {
			if err := t.processModelDep(p, m.Items); err != nil {
				return err
			}
		} else {
			if err := t.processModelDep(m, m.Items); err != nil {
				return err
			}
		}
	case model.SchemaType_Enum:
		if err := addDep(m.Name, model.SchemaType_String.String()); err != nil {
			return err
		}
	case model.SchemaType_Ref:
		if err := addDep(p.Name, m.Ref); err != nil {
			return err
		}
	case model.SchemaType_Complex:
		for _, property := range m.Properties {
			if err := t.processModelDep(m, &property); err != nil {
				return err
			}
		}
	default:
		return t.ErrInvalidType(m.Type.String())
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
		if util.IsSchemaSimpleTypeS(node.Name) {
			continue
		}

		t.reportf(t.reportStack.S(), "processing %v", node.Name)
		inputModel, ok := t.inputModels[node.Name]

		if !ok {
			return t.ErrInvalidType(node.Name)
		}

		outputModel := model.NewOutputModelWithInput(&inputModel)
		t.outputModels[node.Name] = outputModel
		t.OutputSpec.Models = append(t.OutputSpec.Models, outputModel)

		if jsonOut, err := t.buildJson(outputModel); err != nil {
			return err
		} else {
			outputModel.Json = jsonOut
		}
	}

	return nil
}

func (t *Loader) buildJson(outputModel *model.OutputModel) (*orderedmap.OrderedMap, error) {
	jout := orderedmap.New()
	n := outputModel.FullName
	n = path.Base(n)

	switch outputModel.Type {
	case model.SchemaType_String, model.SchemaType_Byte, model.SchemaType_Boolean,
		model.SchemaType_Int, model.SchemaType_Long,
		model.SchemaType_Float, model.SchemaType_Double,
		model.SchemaType_Enum, model.SchemaType_Array:
		jout.Set(n, outputModel.Example)
	case model.SchemaType_Complex:
		for _, property := range outputModel.Properties {
			if pJsonOut, err := t.buildJson(property); err != nil {
				return jout, err
			} else {
				for _, k := range pJsonOut.Keys() {
					v, _ := pJsonOut.Get(k)
					jout.Set(k, v)
				}
			}
		}
	case model.SchemaType_Ref:
		if dep, ok := t.outputModels[outputModel.Ref]; ok {
			jout.Set(n, dep.Example)
		} else {
			return orderedmap.New(), t.ErrInvalidType(outputModel.Ref)
		}

	default:
		return orderedmap.New(), t.ErrInvalidType(outputModel.FullName)
	}

	return jout, nil
}

func (t *Loader) processModel5() error {
	for name, omodel := range t.outputModels {
		if resolved, err := t.processRefs(omodel); err != nil {
			return err
		} else {
			t.outputModels[name] = resolved
		}
	}

	return nil
}

func (t *Loader) processRefs(input *model.OutputModel) (*model.OutputModel, error) {
	var output *model.OutputModel

	switch input.Type {
	case model.SchemaType_Array:
		switch input.Items.Type {
		case model.SchemaType_Ref:
			if ref, ok := t.outputModels[input.Items.Ref]; ok {
				output = model.NewOutputModelWithOutput(input)
				output.Items = model.NewOutputModelWithOutput(ref)
			} else {
				return output, t.ErrInvalidType(input.Items.Type.String())
			}
		case model.SchemaType_String, model.SchemaType_Byte, model.SchemaType_Boolean,
			model.SchemaType_Int, model.SchemaType_Long,
			model.SchemaType_Float, model.SchemaType_Double,
			model.SchemaType_Enum:
			output = model.NewOutputModelWithOutput(input)
		default:
			return output, t.ErrInvalidType(input.Items.Type.String())
		}
	case model.SchemaType_Ref:
		if ref, ok := t.outputModels[input.Ref]; ok {
			output = model.NewOutputModelWithOutput(ref)
			output.Name = input.Name

			description := append([]string{}, input.Description...)
			if len(description) > 0 {
				description = append(description, "")
			}

			description = append(description, ref.Description...)
			output.Description = description
		} else {
			return output, t.ErrInvalidModel(input.Ref)
		}
	case model.SchemaType_Complex:
		properties := make([]*model.OutputModel, 0)

		for _, property := range input.Properties {
			if resolved, err := t.processRefs(property); err != nil {
				return output, err
			} else {
				properties = append(properties, resolved)
			}
		}

		output = input
		input.Properties = properties
	default:
		output = input
	}

	return output, nil
}

func (t *Loader) processModel6() error {
	for name, outputModel := range t.outputModels {
		if outputModel.Type != model.SchemaType_Complex {
			continue
		}

		outputModel, ok := t.outputModels[name]

		if !ok {
			return t.ErrInvalidModel(name)
		}

		modelNamespace := path.Dir(name)

		for _, property := range outputModel.Properties {
			switch property.Type {
			case model.SchemaType_Complex:
			default:
				continue
			}

			propertyNamespace := path.Dir(property.FullName)

			if modelNamespace != propertyNamespace {
				outputModel.Imports = append(outputModel.Imports, property.FullName)
				property.Imported = true
			}
		}
	}

	return nil
}

func (t *Loader) processModel7() error {
	for name, omodel := range t.outputModels {
		if util.IsSchemaSimpleType(omodel.Type) {
			delete(t.outputModels, name)
		}
	}

	return nil
}
