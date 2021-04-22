package loader

import (
	"boundedinfinity/codegen/model"
	"path"
	"strings"
)

func (t Loader) processOperation6(namespace model.BiOutput_Namespace, input model.BiInput_Operation, output *model.BiOutput_Operation) error {
	output.Namespace = namespace.Namespace
	output.Description = t.splitDescription(input.Description)
	output.Input.Name = input.Input.Name
	output.Input.Description = t.splitDescription(input.Input.Description)
	output.Output.Name = input.Output.Name
	output.Output.Description = t.splitDescription(input.Output.Description)

	inputBuiltin := path.Join(model.NAMESPACE_BUILTIN, input.Input.Type)
	inputAbs := path.Join(t.rootNamespace(), input.Input.Type)
	inputRel := path.Join(namespace.Namespace, input.Input.Type)
	var inputInfo *model.TypeInfo

	if info, ok := t.typeMap[inputBuiltin]; ok {
		inputInfo = info
	} else if info, ok := t.typeMap[inputAbs]; ok {
		inputInfo = info
	} else if info, ok := t.typeMap[inputRel]; ok {
		inputInfo = info
	} else {
		return t.NotFound()
	}

	output.Input.SpecName = path.Join(output.SpecName, output.Input.Name)
	output.Input.Namespace = inputInfo.Namespace
	output.Input.SpecType = inputInfo.SpecType

	if namespace.Namespace == inputInfo.Namespace {
		output.Input.Type = inputInfo.InNamespaceType
	} else {
		output.Input.Type = inputInfo.OutOfNamespaceType
	}

	outputBuiltin := path.Join(model.NAMESPACE_BUILTIN, input.Output.Type)
	outputAbs := path.Join(t.rootNamespace(), input.Output.Type)
	outputRel := path.Join(namespace.Namespace, input.Output.Type)
	var outputInfo *model.TypeInfo

	if info, ok := t.typeMap[outputBuiltin]; ok {
		outputInfo = info
	} else if info, ok := t.typeMap[outputAbs]; ok {
		outputInfo = info
	} else if info, ok := t.typeMap[outputRel]; ok {
		outputInfo = info
	} else {
		return t.NotFound()
	}

	output.Output.SpecName = path.Join(output.SpecName, output.Output.Name)
	output.Output.Namespace = outputInfo.Namespace
	output.Output.SpecType = outputInfo.SpecType

	if namespace.Namespace == outputInfo.Namespace {
		output.Output.Type = outputInfo.InNamespaceType
	} else {
		output.Output.Type = outputInfo.OutOfNamespaceType
	}

	return nil
}

func (t *Loader) processOperation8(namespace model.BiOutput_Namespace, input model.BiInput_Operation, output *model.BiOutput_Operation) error {
	m := make(map[string]bool)

	if !strings.HasPrefix(output.Input.Namespace, model.NAMESPACE_BUILTIN) && output.Input.Namespace != namespace.Namespace {
		if _, ok := m[output.Input.Namespace]; !ok {
			m[output.Input.Namespace] = true
		}
	}

	if !strings.HasPrefix(output.Output.Namespace, model.NAMESPACE_BUILTIN) && output.Output.Namespace != namespace.Namespace {
		if _, ok := m[output.Output.Namespace]; !ok {
			m[output.Output.Namespace] = true
		}
	}

	for k := range m {
		output.Imports = append(output.Imports, k)
	}

	return nil
}
