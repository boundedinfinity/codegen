package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
)

func (t Loader) walkSpec(fn WalkFunc, ws ...WalkType) error {
	ctx := &WalkContext{
		Namespace: &WalkContextNamespace{
			Input: t.inputSpec.Specification,
		},
	}

	if err := t.walkNamespace(ctx, fn, ws...); err != nil {
		return err
	}

	return nil
}

func (t Loader) walkNamespace(ctx *WalkContext, fn WalkFunc, ws ...WalkType) error {
	nsSpecPath := t.appendNamespace(ctx.Namespace.Input.Name)
	var outputNamespace *model.OutputNamespace

	if v, ok := t.namespaceMap[nsSpecPath]; ok {
		outputNamespace = v
	} else {
		outputNamespace = model.NewOutputNamespace()
		outputNamespace.SpecPath = nsSpecPath
		outputNamespace.Namespace = nsSpecPath
		t.namespaceMap[nsSpecPath] = outputNamespace
		t.OutputSpec.Namespaces = append(t.OutputSpec.Namespaces, outputNamespace)
	}

	ctx.Namespace.Output = outputNamespace

	if fn != nil && ctx.Namespace.Input.Models != nil {
		for _, inputModel := range ctx.Namespace.Input.Models {
			modelWrapper := func() error {
				modelSpecPath := t.appendNamespace(inputModel.Name)
				defer t.namespaceStack.Pop()
				var outputModel *model.OutputModel

				if v, ok := t.modelMap[modelSpecPath]; ok {
					outputModel = v
				} else {
					outputModel = model.NewOutputModel()
					outputModel.SpecPath = modelSpecPath
					outputModel.Namespace = outputNamespace.Namespace
					t.modelMap[modelSpecPath] = outputModel
					t.OutputSpec.Models = append(t.OutputSpec.Models, outputModel)
				}

				modelCtx := &WalkContext{
					Namespace: ctx.Namespace,
					Model: &WalkContextModel{
						Input:  inputModel,
						Output: outputModel,
					},
				}

				if ContainsWalkType(WALKTYPE_MODEL, ws...) {
					if err := fn(modelCtx); err != nil {
						return err
					}
				}

				if ContainsWalkType(WALKTYPE_PROPERTY, ws...) && inputModel.Properties != nil {
					for _, inputProperty := range inputModel.Properties {
						propertyWrapper := func() error {
							propertySpecPath := t.appendNamespace(inputProperty.Name)
							defer t.namespaceStack.Pop()
							var outputProperty *model.OutputModel

							if v, ok := t.modelMap[propertySpecPath]; ok {
								outputProperty = v
							} else {
								outputProperty = model.NewOutputModel()
								outputProperty.SpecPath = propertySpecPath
								t.modelMap[propertySpecPath] = outputProperty
								outputModel.Properties = append(outputModel.Properties, outputProperty)
							}

							propertyCtx := &WalkContext{
								Namespace: modelCtx.Namespace,
								Model:     modelCtx.Model,
								Property: &WalkContextProperty{
									Input:  inputProperty,
									Output: outputProperty,
								},
							}

							if err := fn(propertyCtx); err != nil {
								return err
							}

							return nil
						}

						if err := propertyWrapper(); err != nil {
							return err
						}
					}

				}

				return nil
			}

			if err := modelWrapper(); err != nil {
				return err
			}
		}
	}

	if ContainsWalkType(WALKTYPE_NAMESPACE, ws...) && fn != nil {
		if err := fn(ctx); err != nil {
			return err
		}
	}

	if ctx.Namespace.Input.Namespaces != nil {
		for _, input := range ctx.Namespace.Input.Namespaces {
			nctx := &WalkContext{
				Namespace: &WalkContextNamespace{
					Input: input,
				},
			}

			if err := t.walkNamespace(nctx, fn, ws...); err != nil {
				return err
			}
		}
	}

	return nil
}

type WalkType string

const (
	WALKTYPE_NAMESPACE WalkType = "namespace"
	WALKTYPE_MODEL     WalkType = "model"
	WALKTYPE_PROPERTY  WalkType = "property"
	WALKTYPE_OPERATION WalkType = "operation"
)

func ContainsWalkType(v WalkType, vs ...WalkType) bool {
	for _, x := range vs {
		if v == x {
			return true
		}
	}
	return false
}

type WalkFunc func(ctx *WalkContext) error

type WalkContext struct {
	Namespace *WalkContextNamespace
	Model     *WalkContextModel
	Property  *WalkContextProperty
}

type WalkContextNamespace struct {
	Input  model.InputNamespace
	Output *model.OutputNamespace
}

type WalkContextModel struct {
	Input  model.InputModel
	Output *model.OutputModel
}

type WalkContextProperty struct {
	Input  model.InputModel
	Output *model.OutputModel
}

func (t *Loader) dumpCtx(ctx *WalkContext) error {
	fmt.Println(util.Jdump(ctx))
	return nil
}
