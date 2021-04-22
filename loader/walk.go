package loader

import (
	"boundedinfinity/codegen/model"
	"path"
)

type NamespaceProcessor func(model.BiInput_Namespace, *model.BiOutput_Namespace) error
type ModelProcessor func(model.BiOutput_Namespace, model.BiInput_Model, *model.BiOutput_Model) error
type PropertyProcessor func(model.BiOutput_Namespace, *model.BiOutput_Model, model.BiInput_Property, *model.BiOutput_Property) error
type OperationProcessor func(model.BiOutput_Namespace, model.BiInput_Operation, *model.BiOutput_Operation) error

func (t *Loader) walk(i int, inputNamespace model.BiInput_Namespace,
	namespaceProcessor NamespaceProcessor,
	modelProcessor ModelProcessor,
	propertyProcessor PropertyProcessor,
	osp OperationProcessor,
) error {
	if i < 0 {
		t.reportStack.Push("specification")
	} else {
		t.reportStack.Push("namespace[%v (%v)]", i, inputNamespace.Name)
	}

	defer t.reportStack.Pop()

	t.namespaceStack.Push(inputNamespace.Name)
	defer t.namespaceStack.Pop()

	var outputNamesapce *model.BiOutput_Namespace

	if v, ok := t.namespaceMap[t.currentNamespace()]; ok {
		outputNamesapce = v
	} else {
		outputNamesapce = model.New_BiOutput_Namespace()
		t.namespaceMap[t.currentNamespace()] = outputNamesapce
	}

	if namespaceProcessor != nil {
		if err := namespaceProcessor(inputNamespace, outputNamesapce); err != nil {
			return err
		}
	}

	if inputNamespace.Models != nil {
		for modelIndex, inputModel := range inputNamespace.Models {
			var outputModel *model.BiOutput_Model

			modelName := path.Join(t.currentNamespace(), inputModel.Name)

			if m, ok := t.modelMap[modelName]; ok {
				outputModel = m
			} else {
				outputModel = model.New_BiOutput_Model()
				outputModel.Name = inputModel.Name
				outputModel.SpecName = modelName
				t.modelMap[modelName] = outputModel
			}

			modelWrap := func() error {
				t.reportStack.Push("model[%v (%v)]", modelIndex, inputModel.Name)
				defer t.reportStack.Pop()

				if modelProcessor != nil {
					if err := modelProcessor(*outputNamesapce, inputModel, outputModel); err != nil {
						return err
					}
				}

				return nil
			}

			if err := modelWrap(); err != nil {
				return err
			}

			if inputModel.Properties != nil && propertyProcessor != nil {
				for properyIndex, inputPropery := range inputModel.Properties {
					propertyName := path.Join(modelName, inputPropery.Name)

					var outputProperty *model.BiOutput_Property

					if v, ok := t.propertyMap[propertyName]; ok {
						outputProperty = v
					} else {
						outputProperty = model.New_BiOutput_Property()
						outputProperty.SpecName = propertyName
						t.propertyMap[propertyName] = outputProperty
						outputModel.Properties = append(outputModel.Properties, outputProperty)
					}

					propertyWrap := func() error {
						t.reportStack.Push("property[%v (%v)]", properyIndex, inputPropery.Name)
						defer t.reportStack.Pop()

						if err := propertyProcessor(*outputNamesapce, outputModel, inputPropery, outputProperty); err != nil {
							return err
						}

						return nil
					}

					if err := propertyWrap(); err != nil {
						return err
					}
				}
			}
		}
	}

	if inputNamespace.Operations != nil && osp != nil {
		for operationIndex, inputOperation := range inputNamespace.Operations {
			var outputOperation *model.BiOutput_Operation

			operationName := path.Join(t.currentNamespace(), inputOperation.Name)

			if o, ok := t.operationMap[operationName]; ok {
				outputOperation = o
			} else {
				outputOperation = model.New_BiOutput_Operation()
				outputOperation.Name = inputOperation.Name
				outputOperation.SpecName = operationName
				t.operationMap[operationName] = outputOperation
			}

			operationWrap := func() error {
				t.reportStack.Push("operation[%v (%v)]", operationIndex, inputOperation.Name)
				defer t.reportStack.Pop()

				if err := osp(*outputNamesapce, inputOperation, outputOperation); err != nil {
					return err
				}

				return nil
			}

			if err := operationWrap(); err != nil {
				return err
			}
		}
	}

	if inputNamespace.Namespaces != nil {
		for ci, cns := range inputNamespace.Namespaces {
			if err := t.walk(ci, cns, namespaceProcessor, modelProcessor, propertyProcessor, osp); err != nil {
				return err
			}
		}
	}

	return nil
}
