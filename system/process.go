package system

import "boundedinfinity/codegen/codegen_type"

func (t *System) Process(paths ...string) error {
	var allProjects []*codegen_type.CodeGenProject

	for len(paths) > 0 {
		projects, err := t.loader.LoadProjectPaths(paths...)

		if err != nil {
			return err
		}

		allProjects = append(allProjects, projects...)

		if paths = t.loader.ExtractProjectPaths(projects); err != nil {
			return err
		}
	}

	if err := t.loader.MergeProjects(allProjects...); err != nil {
		return err
	}

	if err := t.loader.ProcessTemplates(); err != nil {
		return err
	}

	if err := t.loader.Resolve(); err != nil {
		return err
	}

	// if err := t.generator.Process(); err != nil {
	// 	return err
	// }

	// if err := t.generator.Generate(); err != nil {
	// 	return err
	// }

	return nil
}
