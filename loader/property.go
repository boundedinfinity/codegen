package loader

import (
	"boundedinfinity/codegen/model"
	"path"
)

func (t *Loader) propertyProcessor1(namespace model.BiOutput_Namespace, model1 model.BiOutput_Model, input model.BiInput_Property, output *model.BiOutput_Property) error {
	output.Name = input.Name

	if t.isBuiltInType(input.Type) {
		if v, ok := t.builtInTypeMap[input.Type]; ok {
			output.Type = v
			output.Namespace = model.NAMESPACE_BUILTIN
		} else {
			return t.NotFound()
		}
	} else if t.isCustomType(t.absoluteNamespace(input.Type)) {
		if v, ok := t.customTypeMap[t.absoluteNamespace(input.Type)]; ok {
			output.Type = v
			output.Namespace = path.Dir(t.absoluteNamespace(input.Type))
		} else {
			return t.NotFound()
		}
	}

	return nil
}

// func (t *Loader) processProperty(i int, input model.BiInput_Property, output *model.BiOutput_Property) error {
// 	t.reportStack.Push("properties[%v (%v)]", i, input.Name)
// 	defer t.reportStack.Pop()

// 	output.Name = input.Name
// 	output.Description = t.splitDescription(input.Description)
// 	output.Validations = make([]model.BiOutput_Validation, 0)
// 	output.JsonPath = strings.Join(t.namespaceStack.S(), ".")
// 	jname := util.CamelCase(input.Name)
// 	output.JsonStruture = make(map[string]interface{})
// 	found := false

// 	if v, ok := t.builtInTypeMap[input.Type]; ok {
// 		output.Namespace = model.NAMESPACE_BUILTIN
// 		output.Type = v
// 		found = true
// 		switch input.Type {
// 		case "integer":
// 			if v, err := json2Int64(input.Example); err != nil {
// 				return err
// 			} else {
// 				output.JsonStruture[jname] = v
// 			}

// 		case "string":
// 			if v, err := json2Str(input.Example); err != nil {
// 				return err
// 			} else {
// 				output.JsonStruture[jname] = v
// 			}
// 		}
// 	}

// 	if !found {
// 		if v, ok := t.customTypeMap[input.Type]; ok {
// 			output.Namespace = t.absoluteNamespace(input.Type)
// 			output.Type = v
// 			found = true
// 		}
// 	}

// 	if !found {
// 		relType := t.relativeNamespace(input.Type)

// 		if v, ok := t.customTypeMap[relType]; ok {
// 			output.Namespace = t.absoluteNamespace(input.Type)
// 			output.Type = v
// 			found = true
// 		}
// 	}

// 	if !found {
// 		t.report("not found: %v", input.Type)
// 	}

// 	if input.Validations != nil {
// 		for _, validation := range input.Validations {
// 			output.Validations = append(output.Validations, model.BiOutput_Validation{
// 				Minimum:  validation.Minimum,
// 				Maximum:  validation.Maximum,
// 				Required: validation.Required,
// 			})
// 		}
// 	}

// 	return nil
// }
