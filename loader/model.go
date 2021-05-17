package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"reflect"
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

		if util.IsSchemaPrimitive(m.Type) {
			dNode, ok := t.dependencies[m.Type.String()]

			if !ok {
				return t.ErrInvalidType(m.Type.String())
			}

			mNode.Add(dNode)
		} else {
			switch m.Type {
			case model.SchemaType_Ref:
				dNode, ok := t.dependencies[m.Ref.Ref]

				if !ok {
					return t.ErrInvalidType(m.Type.String())
				}

				mNode.Add(dNode)
			case model.SchemaType_Complex:
				for _, property := range m.Complex.Properties {
					if util.IsSchemaPrimitive(property.Type) {
						dNode, ok := t.dependencies[property.Type.String()]

						if !ok {
							return t.ErrInvalidType(property.Type.String())
						}

						mNode.Add(dNode)
					} else {
						switch property.Type {
						case model.SchemaType_Ref:
							dNode, ok := t.dependencies[property.Ref.Ref]

							if !ok {
								return t.ErrInvalidType(property.Type.String())
							}

							mNode.Add(dNode)
						default:
							return t.ErrInvalidType(fmt.Sprintf("%v/%v", m.Name, property.Name))
						}
					}
				}
			case model.SchemaType_Enum:
				// Nothing to do here
			default:
				return t.ErrInvalidType(n)
			}
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

		outputModel := model.NewOutputModel()
		t.modelMap[node.Name] = outputModel
		outputModel.Name = node.Name
		outputModel.Description = t.splitDescription(inputModel.Description)

		if example, err := t.buildExample(inputModel); err != nil {
			return err
		} else {
			outputModel.Example = example
		}
	}

	return nil
}

func (t *Loader) buildExample(inputModel model.InputModel) (map[string]interface{}, error) {
	exampleItem := func(name string, v interface{}) map[string]interface{} {
		nname := path.Base(name)
		r := make(map[string]interface{})

		switch reflect.TypeOf(v).Kind() {
		case reflect.Slice:
			var ss []string
			s := reflect.ValueOf(v)
			for i := 0; i < s.Len(); i++ {
				ss = append(ss, fmt.Sprintf("%v", s.Index(i)))
			}
			r[nname] = ss
		default:
			r[nname] = fmt.Sprintf("%v", v)
		}
		return r
	}

	switch inputModel.Type {
	case model.SchemaType_String:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.StringArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.String.Example), nil
		}
	case model.SchemaType_Boolean:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.BoolArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Bool.Example), nil
		}
	case model.SchemaType_Int:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.IntArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Int.Example), nil
		}
	case model.SchemaType_Long:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.LongArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Long.Example), nil
		}
	case model.SchemaType_Float:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.FloatArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Float.Example), nil
		}
	case model.SchemaType_Double:
		if inputModel.Array {
			return exampleItem(inputModel.Name, inputModel.DoubleArray.Example), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Double.Example), nil
		}
	case model.SchemaType_Enum:
		if inputModel.Array {
			return exampleItem(inputModel.Name, []string{inputModel.Enum.Example}), nil
		} else {
			return exampleItem(inputModel.Name, inputModel.Enum.Example), nil
		}
	case model.SchemaType_Complex:
		combined := make(map[string]interface{})

		for _, property := range inputModel.Complex.Properties {
			if example, err := t.buildExample(property); err != nil {
				return map[string]interface{}{}, nil
			} else {
				for k, v := range example {
					combined[k] = v
				}
			}
		}
		return combined, nil
	case model.SchemaType_Ref:
		combined := make(map[string]interface{})
		ref, ok := t.modelMap[inputModel.Ref.Ref]

		if !ok {
			t.ErrInvalidType(inputModel.Ref.Ref)
		}

		combined[inputModel.Name] = ref.Example
		return combined, nil
	default:
		return map[string]interface{}{}, t.ErrInvalidType(inputModel.Type.String())
	}
}
