package system

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/template_manager"
	"os"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/optioner"
)

func (t *System) init() error {
	if t.workDir.Empty() {
		homeDir := os.Getenv("HOME")
		var workDir string

		if homeDir != "" {
			workDir = filepath.Join(homeDir, ".config", DEFAULT_WORKDIR_NAME)
		} else {
			workDir = filepath.Join(os.TempDir(), DEFAULT_WORKDIR_NAME)
		}

		t.workDir = optioner.Some(workDir)
	}

	if t.cacheDir.Empty() {
		cacheDir := filepath.Join(t.workDir.Get(), "cache")
		t.cacheDir = optioner.Some(cacheDir)
	}

	// c, err := cacher.New(cacher.CacheDir(t.cacheDir.Get()))

	// if err != nil {
	// 	return err
	// }

	// t.cacher = c

	ld, err := loader.New(
		// loader.Cacher(c),
		loader.Canonicals(t.typeManager),
		loader.ProjectManager(t.projectManager),
	)

	if err != nil {
		return err
	}

	t.loader = ld

	tm, err := template_manager.New(
		// template_manager.Cacher(c),
		template_manager.TypeManaager(t.typeManager),
		template_manager.ProjectManager(t.projectManager),
	)

	if err != nil {
		return err
	}

	t.tm = tm

	g, err := generator.New(
		generator.TemplateManager(t.tm),
		generator.TypeManager(t.typeManager),
		generator.ProjectManager(t.projectManager),
		generator.Loader(t.loader),
	)

	if err != nil {
		return err
	}

	t.generator = g

	return nil
}

var (
	DEFAULT_WORKDIR_NAME = "codegen"
)

type Arg func(*System)
