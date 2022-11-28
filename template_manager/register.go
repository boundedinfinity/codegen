package template_manager

import "boundedinfinity/codegen/model"

func (t *TemplateManager) Register(files ...model.CodeGenSchemaTemplateFile) error {
	for _, file := range files {
		if err := t.register(file); err != nil {
			return err
		}
	}

	return nil
}

func (t *TemplateManager) register(file model.CodeGenSchemaTemplateFile) error {

	return nil
}
