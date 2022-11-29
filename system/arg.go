package system

import (
	"boundedinfinity/codegen/cacher"
	"boundedinfinity/codegen/generator"
	"boundedinfinity/codegen/template_manager"
	"os"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/optioner"
)

var (
	DEFAULT_WORKDIR_NAME = "codegen"
)

type Arg func(*System)

func OutputDir(v string) Arg {
	return func(t *System) {
		t.outputDir = optioner.Some(v)
	}
}

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

	if t.outputDir.Empty() {
		outputDir := filepath.Join(t.workDir.Get(), "output")
		t.outputDir = optioner.Some(outputDir)
	}

	if t.cacheDir.Empty() {
		cacheDir := filepath.Join(t.workDir.Get(), "cache")
		t.cacheDir = optioner.Some(cacheDir)
	}

	c, err := cacher.New(cacher.CacheDir(t.cacheDir.Get()))

	if err != nil {
		return err
	}

	t.cacher = c

	tm, err := template_manager.New(template_manager.Cacher(c))

	if err != nil {
		return err
	}

	t.tm = tm

	g, err := generator.New(
		generator.DestDir(t.outputDir.Get()),
		generator.TemplateManager(t.tm),
	)

	if err != nil {
		return err
	}

	t.generator = g

	return nil
}
