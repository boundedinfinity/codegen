package loader

import (
	"boundedinfinity/codegen/util"
	"path"
	"strings"
)

func (t *Loader) processModel1(ctx *WalkContext) error {
	input := ctx.Model.Input
	output := ctx.Model.Output

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)

	if strings.Contains(input.Type, "/") {
		output.Type = path.Join(t.rootName(), input.Type)
	} else {
		output.Type = input.Type
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
	switch ctx.Model.Input.Type {
	case "string":
		if v, err := json2Str(ctx.Model.Input.Example); err != nil {
			return err
		} else {
			ctx.Model.Output.JsonStructure[ctx.Model.Input.Name] = v
		}
	case "int":
		if v, err := json2Int64(ctx.Model.Input.Example); err != nil {
			return err
		} else {
			ctx.Model.Output.JsonStructure[ctx.Model.Input.Name] = v
		}
	default:
		return t.ErrInvalidType(ctx.Model.Input.Type)
	}

	return nil
}

func (t *Loader) processModelComplexType(ctx *WalkContext) error {
	for _, property := range ctx.Model.Output.Properties {
		pv, ok := t.modelMap[property.Type]

		if !ok {
			return t.ErrCustomTypeNotFound(property.Type)
		}

		if len(pv.JsonStructure) > 1 {
			ctx.Model.Output.JsonStructure[util.CamelCase(property.Name)] = pv.JsonStructure
		} else {
			for _, ev := range pv.JsonStructure {
				ctx.Model.Output.JsonStructure[util.CamelCase(property.Name)] = ev
			}
		}
	}

	return nil
}
