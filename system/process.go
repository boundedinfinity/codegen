package system

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *System) Process(paths ...string) error {
	projects, err := t.loader.LoadProjectPath(o.None[string](), paths...)

	if err != nil {
		return err
	}

	if err := t.loader.MergeProjects(projects...); err != nil {
		return err
	}

	if err := t.loader.Resolve(); err != nil {
		return err
	}

	if err := t.loader.ProcessTemplates(); err != nil {
		return err
	}

	if err := t.loader.ProcessNamespace(); err != nil {
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
