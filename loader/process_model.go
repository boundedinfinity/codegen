package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
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
	for n, m := range t.inputModels {
		mNode, ok := t.dependencies[n]

		if !ok {
			return t.ErrInvalidType(n)
		}

		handleDep := func(n, v string) error {
			dNode, ok := t.dependencies[v]

			if !ok {
				return t.ErrInvalidType(n)
			}

			mNode.Add(dNode)
			return nil
		}

		switch m.Type {
		case model.SchemaType_String:
			if err := handleDep(m.Type.String(), m.Type.String()); err != nil {
				return err
			}
		case model.SchemaType_Int:
			if err := handleDep(m.Type.String(), m.Type.String()); err != nil {
				return err
			}
		case model.SchemaType_Array:
			switch m.Items.Type {
			case model.SchemaType_String:
				if err := handleDep(m.Items.Type.String(), m.Items.Type.String()); err != nil {
					return err
				}
			}
		case model.SchemaType_Enum:
			// no dependencies
		case model.SchemaType_Ref:
			if err := handleDep(m.Ref, m.Type.String()); err != nil {
				return err
			}
		case model.SchemaType_Complex:
			for _, property := range m.Properties {
				if util.IsSchemaPrimitive(property.Type) {
					dNode, ok := t.dependencies[property.Type.String()]

					if !ok {
						return t.ErrInvalidType(property.Type.String())
					}

					mNode.Add(dNode)
				} else {
					switch property.Type {
					case model.SchemaType_Ref:
						dNode, ok := t.dependencies[property.Ref]

						if !ok {
							return t.ErrInvalidType(property.Type.String())
						}

						mNode.Add(dNode)
					default:
						return t.ErrInvalidType(fmt.Sprintf("%v/%v", m.Name, property.Name))
					}
				}
			}
		default:
			return t.ErrInvalidType(m.Type.String())
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
		if util.IsSchemaPrimitiveS(node.Name) {
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
	jsonOut := orderedmap.New()
	n := outputModel.Name
	n = path.Base(n)

	switch outputModel.Type {
	case model.SchemaType_String:
		jsonOut.Set(n, outputModel.Example)
	case model.SchemaType_Int:
		jsonOut.Set(n, outputModel.Example)
	case model.SchemaType_Enum:
		jsonOut.Set(n, outputModel.Example)
	case model.SchemaType_Array:
		jsonOut.Set(n, outputModel.Example)
	case model.SchemaType_Complex:
		for _, property := range outputModel.Properties {
			if pJsonOut, err := t.buildJson(&property); err != nil {
				return jsonOut, err
			} else {
				for _, k := range pJsonOut.Keys() {
					v, _ := pJsonOut.Get(k)
					jsonOut.Set(k, v)
				}
			}
		}
	case model.SchemaType_Ref:
		if dep, ok := t.outputModels[outputModel.Ref]; ok {
			jsonOut.Set(n, dep.Example)
		} else {
			return orderedmap.New(), t.ErrInvalidType(outputModel.Ref)
		}

	default:
		return orderedmap.New(), t.ErrInvalidType(outputModel.Name)
	}

	return jsonOut, nil
}

func (t *Loader) processModel5() error {
	for name, inputModel := range t.inputModels {
		if inputModel.Type != model.SchemaType_Complex {
			continue
		}

		outputModel, ok := t.outputModels[name]

		if !ok {
			return t.ErrInvalidModel(name)
		}

		modelNamespace := path.Dir(name)

		for _, property := range inputModel.Properties {
			if property.Type != model.SchemaType_Ref {
				continue
			}

			propertyNamespace := path.Dir(property.Ref)

			if modelNamespace != propertyNamespace {
				outputModel.Imports = append(outputModel.Imports, property.Ref)
			}
		}
	}

	return nil
}
