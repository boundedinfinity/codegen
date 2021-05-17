package loader

import (
	"boundedinfinity/codegen/model"
)

type ExampleExtractor func(example interface{}) (interface{}, error)

func (t *Loader) json2Interface(inputModel model.InputModel, outputModel *model.OutputModel, v interface{}) error {
	// if inputModel.Model.Example == nil {
	// 	return nil
	// }

	// if v == nil {
	// 	return nil
	// }

	// s := fmt.Sprintf("%v", inputModel.Model.Example)
	// bs := []byte(s)

	// switch inputModel.Source {
	// case model.InputSource_Yaml:
	// 	if err := yaml.Unmarshal(bs, v); err != nil {
	// 		return err
	// 	}
	// case model.InputSource_Json:
	// 	if err := json.Unmarshal(bs, v); err != nil {
	// 		return err
	// 	}
	// default:
	// 	return t.ErrInvalidSource(inputModel.Source)
	// }

	// outputModel.Example = v

	return nil
}
