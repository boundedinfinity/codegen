package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"strings"
)

func (t *Loader) propertyProcessor1(namespace model.BiOutput_Namespace, model1 *model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)

	return nil
}

func (t *Loader) propertyProcessor2(namespace model.BiOutput_Namespace, model1 *model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
	typ := input.Type
	var isCollection bool

	if strings.HasSuffix(typ, model.COLLECTION_SUFFIX) {
		isCollection = true
		typ = strings.TrimSuffix(typ, model.COLLECTION_SUFFIX)
	}

	builtIn := path.Join(model.NAMESPACE_BUILTIN, typ)
	abs := path.Join(t.rootNamespace(), typ)
	rel := path.Join(namespace.Namespace, typ)

	if info, ok := t.typeMap[builtIn]; ok {
		output.SpecType = builtIn
		output.Namespace = info.Namespace
		output.Type = info.InNamespaceType
	} else if info, ok := t.typeMap[abs]; ok {
		output.SpecType = abs
		output.Namespace = info.Namespace

		if namespace.Namespace == abs {
			output.Type = info.InNamespaceType
		} else {
			output.Type = info.OutOfNamespaceType
		}
	} else if info, ok := t.typeMap[rel]; ok {
		output.SpecType = rel
		output.Namespace = info.Namespace

		if namespace.Namespace == rel {
			output.Type = info.InNamespaceType
		} else {
			output.Type = info.OutOfNamespaceType
		}
	} else {
		return t.NotFound()
	}

	if isCollection {
		output.Type = strings.Join([]string{model.COLLECTION_SUFFIX, output.Type}, "")
	}

	return nil
}

func (t *Loader) propertyProcessor3(namespace model.BiOutput_Namespace, model1 *model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
	if strings.HasPrefix(output.Namespace, model.NAMESPACE_BUILTIN) {
		return nil
	}

	if _, ok := t.modelMap[output.SpecType]; !ok {
		return nil
	}

	extNode, ok := t.depNodes[output.SpecType]

	if !ok {
		extNode = NewNode(output.SpecType)
		t.depNodes[extNode.name] = extNode
	}

	thisNode, ok := t.depNodes[model1.SpecName]

	if !ok {
		thisNode = NewNode(model1.SpecName)
		t.depNodes[thisNode.name] = thisNode
	}

	thisNode.Add(extNode.name)

	return nil
}

func (t *Loader) propertyProcessor4(specName string) PropertyProcessor {
	return func(namespace model.BiOutput_Namespace, model1 *model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
		if specName != model1.SpecName {
			return nil
		}

		return t.propertyProcessorJson(namespace, model1, input, output)
	}
}

func (t *Loader) propertyProcessorJson(namespace model.BiOutput_Namespace, model1 *model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
	if output.JsonPath != "" {
		return nil
	}

	output.JsonPath = util.CamelCase(output.Name)

	switch output.Namespace {
	case "builtin/biginteger":
		if v, err := json2Int64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/smallinteger":
		if v, err := json2Int64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/integer":
		if v, err := json2Int64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}

	case "builtin/float":
		if v, err := json2Float64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/bigfloat":
		if v, err := json2Float64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/smallfloat":
		if v, err := json2Float64(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/boolean":
		if v, err := json2Boolean(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	case "builtin/string":
		if v, err := json2Str(input.Example); err != nil {
			return err
		} else {
			output.JsonStruture[output.JsonPath] = v
			output.Example = fmt.Sprintf("%v", v)
		}
	default:
		if v, ok := t.modelMap[output.SpecType]; ok {
			output.JsonStruture[output.JsonPath] = v.JsonStruture
		} else {
			return t.NotFound()
		}
	}

	for k, v := range output.JsonStruture {
		model1.JsonStruture[k] = v
	}

	return nil
}
