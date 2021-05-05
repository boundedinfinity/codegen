package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"path"
	"strings"
)

func (t *Loader) processModel1(ctx *WalkContext) error {
	input := ctx.Model.Input
	output := ctx.Model.Output

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)

	typ := input.Type

	if strings.HasSuffix(typ, model.COLLECTION_SUFFIX) {
		output.Collection = true
		typ = strings.ReplaceAll(typ, model.COLLECTION_SUFFIX, "")
	}

	if strings.Contains(input.Type, "/") {
		output.Type = path.Join(t.rootName(), typ)
	} else {
		output.Type = typ
	}

	if _, ok := t.dependencies[output.SpecPath]; !ok {
		t.dependencies[output.SpecPath] = NewNode(output.SpecPath)
	}

	return nil
}

func (t *Loader) processModel2(specPath string) WalkFunc {
	return func(ctx *WalkContext) error {
		if specPath != ctx.Model.Output.SpecPath {
			return nil
		}

		if ctx.Model.Input.Properties == nil {
			if err := t.processModelBuiltinType(ctx); err != nil {
				return nil
			}
		} else {
			if err := t.processModelComplexType(ctx); err != nil {
				return nil
			}
		}

		return nil
	}
}

func (t *Loader) processModelBuiltinType(ctx *WalkContext) error {
	input := ctx.Model.Input
	output := ctx.Model.Output

	extract := func(fn ExampleExtractor) error {
		jname := util.CamelCase(ctx.Model.Input.Name)
		if v, err := fn(input.Example); err != nil {
			return err
		} else {
			if output.Collection {
				output.JsonStructure[jname] = []interface{}{v}
			} else {
				output.JsonStructure[jname] = v
			}
		}

		return nil
	}

	switch input.Type {
	case "string":
		return extract(json2Str)
	case "int":
		return extract(json2Int64)
	default:
		return t.ErrInvalidType(input.Type)
	}
}

func (t *Loader) processModelComplexType(ctx *WalkContext) error {
	output := ctx.Model.Output

	for _, property := range ctx.Model.Output.Properties {
		pv, ok := t.modelMap[property.Type]
		jname := util.CamelCase(property.Name)

		if !ok {
			return t.ErrCustomTypeNotFound(property.Type)
		}

		if len(pv.JsonStructure) > 1 {
			if property.Collection {
				output.JsonStructure[jname] = []interface{}{pv.JsonStructure}
			} else {
				output.JsonStructure[jname] = pv.JsonStructure
			}
		} else {
			for _, ev := range pv.JsonStructure {
				if property.Collection {
					output.JsonStructure[jname] = []interface{}{ev}
				} else {
					output.JsonStructure[jname] = ev
				}
			}
		}
	}

	return nil
}
