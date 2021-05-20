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
			if pJsonOut, err := t.buildJson(property); err != nil {
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
			if property.Type != model.SchemaType_Ref {
				continue
			}

			propertyNamespace := path.Dir(property.Ref)

			if modelNamespace != propertyNamespace {
				outputModel.Imports = append(outputModel.Imports, property.Ref)
				property.Imported = true
			}
		}
	}

	return nil
}
