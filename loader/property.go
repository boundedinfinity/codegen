package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) processProperty1(ctx *WalkContext) error {
	input := ctx.Property.Input
	output := ctx.Property.Output

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	typ := input.Type

	if strings.HasSuffix(typ, model.COLLECTION_SUFFIX) {
		output.Collection = true
		typ = strings.ReplaceAll(typ, model.COLLECTION_SUFFIX, "")
	}

	output.Type = path.Join(t.getNamespace(), typ)
	mt, ok := t.modelMap[output.Type]

	if !ok {
		return t.Errorf("dependency not found : %v", output.Type)
	}

	output.Namespace = mt.Namespace

	if _, ok := t.modelMap[ctx.Model.Output.SpecPath]; !ok {
		return t.Errorf("model not found : %v", ctx.Model.Output.SpecPath)
	}

	if mdep, ok := t.dependencies[ctx.Model.Output.SpecPath]; !ok {
		return t.Errorf("model not found : %v", ctx.Model.Output.SpecPath)
	} else {
		mdep.Add(output.Type)
	}

	return nil
}
func (t *Loader) processProperty2(ctx *WalkContext) error {
	st := ctx.Property.Output.Type
	pm, ok := t.modelMap[ctx.Property.Output.Type]

	if !ok {
		return t.ErrInvalidType(st)
	}

	output := ctx.Property.Output

	if strings.HasSuffix(pm.Namespace, model.NAMESPACE_BUILTIN) {
		output.Type = pm.Type
	} else if ctx.Model.Output.Namespace != pm.Namespace {
		output.Type = fmt.Sprintf("%v.%v", path.Base(pm.Namespace), pm.Name)
		ctx.Model.Output.Imports = append(ctx.Model.Output.Imports, pm.Namespace)
	} else {
		output.Type = pm.Name
	}

	if output.Collection {
		output.Type = fmt.Sprintf("%v%v", model.COLLECTION_SUFFIX, output.Type)
	}

	ctx.Model.Output.Imports = util.StrSliceDedup(ctx.Model.Output.Imports)

	return nil
}
