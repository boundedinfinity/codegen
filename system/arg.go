package system

import (
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/loader"
	"boundedinfinity/codegen/renderer"
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

	rd, err := renderer.New(
		renderer.TypeManager(t.typeManager),
		renderer.ProjectManager(t.projectManager),
		renderer.TemplateManager(t.templateManager),
	)

	if err != nil {
		return err
	}

	t.renderer = rd

	ld, err := loader.New(
		// loader.Cacher(c),
		loader.TypeManager(t.typeManager),
		loader.ProjectManager(t.projectManager),
		loader.TemplateManager(t.templateManager),
		loader.Renderer(t.renderer),
	)

	if err != nil {
		return err
	}

	t.loader = ld

	g, err := generator.New(
		generator.TemplateManager(t.templateManager),
		generator.TypeManager(t.typeManager),
		generator.ProjectManager(t.projectManager),
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
