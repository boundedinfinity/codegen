package system

import (
	"boundedinfinity/codegen/cacher"
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

	ld, err := loader.New(
		loader.Cacher(c),
		loader.Jsonschemas(t.jsonSchemas),
		loader.Canonicals(t.canonicals),
		loader.MergedCodeGen(t.mergedCodeGen),
	)

	if err != nil {
		return err
	}

	t.loader = ld

	tm, err := template_manager.New(
		template_manager.Cacher(c),
		template_manager.CanonicalCombined(t.canonicals),
		template_manager.CodeGenSchema(t.mergedCodeGen),
	)

	if err != nil {
		return err
	}

	t.tm = tm

	g, err := generator.New(
		generator.DestDir(t.outputDir.Get()),
		generator.TemplateManager(t.tm),
		generator.Canonicals(t.canonicals),
		generator.CodeGenSchema(t.mergedCodeGen),
	)

	if err != nil {
		return err
	}

	t.generator = g

	return nil
}
