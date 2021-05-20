package loader

import (
	"boundedinfinity/codegen/model"
	"boundedinfinity/codegen/util"
	"fmt"
	"path"
	"path/filepath"
	"strings"
)

func (t *Loader) processTemplate1() error {
	for namespace, inputTemplates := range t.inputTemplates {
		newTemplates := make([]model.InputTemplate, 0)

		for _, inputTemplate := range inputTemplates {
			newTemplate := model.InputTemplate{
				Header:    inputTemplate.Header,
				Type:      inputTemplate.Type,
				Namespace: inputTemplate.Namespace,
			}

			if filepath.IsAbs(inputTemplate.Path) {
				if ok, err := util.PathExists(inputTemplate.Path); err != nil {
					return err
				} else if !ok {
					return t.ErrTemplatePathNotFound(inputTemplate.Path)
				} else {
					newTemplate.Path = inputTemplate.Path
				}
			} else {
				var found bool

				for _, inputPath := range t.inputPaths {
					abs, err := filepath.Abs(inputPath)

					if err != nil {
						return err
					}

					abs = filepath.Dir(inputPath)
					abs = filepath.Join(abs, inputTemplate.Path)

					if ok, err := util.PathExists(abs); err != nil {
						return err
					} else if !ok {
						return t.ErrTemplatePathNotFound(inputTemplate.Path)
					} else {
						newTemplate.Path = abs
						found = true
						break
					}
				}

				if !found && t.OutputSpec.Info.InputDir != "" {
					if t.OutputSpec.Info.InputDir != "" {
						relPath := filepath.Join(t.OutputSpec.Info.InputDir, inputTemplate.Path)
						abs, err := filepath.Abs(relPath)

						if err != nil {
							return err
						}

						if ok, err := util.PathExists(abs); err != nil {
							return err
						} else if !ok {
							return t.ErrTemplatePathNotFound(inputTemplate.Path)
						} else {
							newTemplate.Path = abs
						}
					} else {
						return t.ErrTemplatePathNotFound(inputTemplate.Path)
					}
				}
			}

			newTemplates = append(newTemplates, newTemplate)
		}

		t.inputTemplates[namespace] = newTemplates
	}

	return nil
}

func (t *Loader) processTemplate2() error {
	for _, inputTemplates := range t.inputTemplates {
		for _, inputTemplate := range inputTemplates {
			for {
				t.appendInfoTemplate(inputTemplate)
				inputTemplate.Namespace = path.Dir(inputTemplate.Namespace)

				if inputTemplate.Namespace == "." {
					break
				}
			}
		}
	}

	return nil
}

func (t *Loader) processTemplate3() error {
	for _, outputModel := range t.outputModels {
		namespace := outputModel.Name

		for {
			namespace = path.Dir(namespace)

			if inputTemplates, ok := t.inputTemplates[namespace]; ok {
				for _, inputTemplate := range inputTemplates {
					if inputTemplate.Type != model.TemplateType_Model {
						continue
					}

					outputTemplate := model.NewOutputTemplate()
					outputTemplate.Type = inputTemplate.Type
					outputTemplate.Input = inputTemplate.Path
					outputTemplate.Header = t.splitDescription(inputTemplate.Header)
					outputModel.Templates = append(outputModel.Templates, outputTemplate)
				}
			}

			if namespace == "." {
				break
			}
		}
	}

	for _, outputOperation := range t.outputOperations {
		namespace := outputOperation.Name

		for {
			namespace = path.Dir(namespace)

			if inputTemplates, ok := t.inputTemplates[namespace]; ok {
				for _, inputTemplate := range inputTemplates {
					if inputTemplate.Type != model.TemplateType_Operation {
						continue
					}

					outputTemplate := model.NewOutputTemplate()
					outputTemplate.Type = inputTemplate.Type
					outputTemplate.Input = inputTemplate.Path
					outputTemplate.Header = t.splitDescription(inputTemplate.Header)
					outputOperation.Templates = append(outputOperation.Templates, outputTemplate)
				}
			}

			if namespace == "." {
				break
			}
		}
	}

	return nil
}

func (t *Loader) processTemplate4() error {
	for _, outputModel := range t.outputModels {
		for _, template := range outputModel.Templates {
			if err := t.processOutputTemplate(path.Base(outputModel.Name), path.Dir(outputModel.Name), template); err != nil {
				return err
			}
		}
	}

	for _, outputOperation := range t.outputOperations {
		for _, template := range outputOperation.Templates {
			if err := t.processOutputTemplate(path.Base(outputOperation.Name), path.Dir(outputOperation.Name), template); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) processOutputTemplate(name, namespace string, template *model.OutputTemplate) error {
	var templateTypeExt string
	var templateExt string
	var fullPath string

	templateTypeExt = template.Input
	templateTypeExt = filepath.Ext(templateTypeExt)
	templateTypeExt = strings.TrimPrefix(templateTypeExt, ".")

	templateExt = template.Input
	templateExt = strings.TrimSuffix(templateExt, filepath.Ext(template.Input))
	templateExt = filepath.Ext(templateExt)
	templateExt = strings.TrimPrefix(templateExt, ".")

	fullPath = name

	if t.OutputSpec.Info.FilenameMarker != model.DEFAULT_FILENAME_DISABLE {
		if t.OutputSpec.Info.FilenameMarker != "" {
			fullPath = fmt.Sprintf("%v.%v", fullPath, t.OutputSpec.Info.FilenameMarker)
		} else {
			fullPath = fmt.Sprintf("%v.%v", fullPath, model.DEFAULT_FILENAME_MARKER)
		}
	}

	fullPath = fmt.Sprintf("%v.%v", fullPath, templateExt)
	fullPath = path.Join(t.OutputSpec.Info.OutputDir, namespace, fullPath)

	template.Output = fullPath
	return nil
}
