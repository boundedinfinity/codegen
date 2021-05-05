package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) processOperationModel(op *model.OutputOperation, input model.InputModel, output *model.OutputModel) error {
	output.Name = input.Name
	typ := path.Join(t.rootName(), input.Type)

	if strings.HasSuffix(typ, model.COLLECTION_SUFFIX) {
		output.Collection = true
		typ = strings.ReplaceAll(typ, model.COLLECTION_SUFFIX, "")
	}

	output.Type = typ

	if mt, ok := t.modelMap[output.Type]; !ok {
		return t.ErrCustomTypeNotFound(output.Type)
	} else {
		output.Namespace = mt.Namespace
		output.JsonStructure = mt.JsonStructure

		if output.Namespace != mt.Namespace {
			op.Imports = append(op.Imports, mt.Namespace)
			output.Type = fmt.Sprintf("%v.%v", path.Base(mt.Namespace), mt.Name)
		} else {
			output.Type = mt.Name
		}

		if output.Collection {
			output.Type = fmt.Sprintf("%v%v", model.COLLECTION_SUFFIX, output.Type)
		}
	}

	return nil
}

func (t *Loader) processOperation1(ctx *WalkContext) error {
	input := ctx.Operation.Input
	output := ctx.Operation.Output

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)

	if err := t.processOperationModel(output, input.Input, &output.Input); err != nil {
		return err
	}

	if err := t.processOperationModel(output, input.Output, &output.Output); err != nil {
		return err
	}

	output.Imports = util.StrSliceDedup(output.Imports)

	return nil
}

func (t *Loader) processOperation2(ctx *WalkContext) error {
	return nil
}
