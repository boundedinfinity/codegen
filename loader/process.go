package loader

import "boundedinfinity/codegen/codegen_type"

func (t *Loader) ExtractPaths(projects []codegen_type.CodeGenProject) []string {
	var paths []string

	codegen_type.Walker().Type(func(project codegen_type.CodeGenProject, typ codegen_type.CodeGenType) error {
		switch c := typ.(type) {
		case *codegen_type.CodeGenTypePath:
			if c.SourcePath.Defined() {
				paths = append(paths, c.SourcePath.Get())
			}
		}

		return nil
	}).Walk(projects...)

	return paths
}

func (t *Loader) ProcessTemplates() error {
	var templatePaths []string

	for _, lc := range t.projectManager.Projects {
		for _, file := range lc.Templates.Files {
			if file.SourcePath.Defined() {
				templatePaths = append(templatePaths, file.SourcePath.Get())
			}
		}
	}

	if err := t.LoadTemplatePaths(templatePaths...); err != nil {
		return err
	}

	return nil
}
