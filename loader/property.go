package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"strings"
)

func (t *Loader) processProperty(i int, input model.BiInput_Property, output *model.BiOutput_Property) error {
	t.reportStack.Push("properties[%v]", i)
	defer t.reportStack.Pop()

	output.Name = input.Name
	output.Description = t.splitDescription(input.Description)
	output.Validations = make([]model.BiOutput_Validation, 0)
	output.JsonPath = strings.Join(t.modelStack.S(), ".")
	jname := util.CamelCase(input.Name)
	output.JsonStruture = make(map[string]interface{})

	if v, ok := t.builtInTypeMap[input.Type]; ok {
		output.Namespace = model.NAMESPACE_BUILTIN
		output.Type = v
		switch input.Type {
		case "integer":
			if v, err := json2Int64(input.Example); err != nil {
				return err
			} else {
				output.JsonStruture[jname] = v
			}
			return nil
		case "string":
			if v, err := json2Str(input.Example); err != nil {
				return err
			} else {
				output.JsonStruture[jname] = v
			}
			return nil
		default:
			// if typeInfo, ok := t.getMappedType(input.Type); ok {
			//     m[name] = typeInfo.
			// } else {
			//     return model.NotFoundErr
			// }
		}
	}

	if input.Validations != nil {
		for _, validation := range input.Validations {
			output.Validations = append(output.Validations, model.BiOutput_Validation{
				Minimum:  validation.Minimum,
				Maximum:  validation.Maximum,
				Required: validation.Required,
			})
		}
	}

	return nil
}
