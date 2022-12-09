package generator

import (
	"boundedinfinity/codegen/canonical"
	"boundedinfinity/codegen/template_manager"
	"boundedinfinity/codegen/util"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/boundedinfinity/go-commoner/extentioner"
	"github.com/boundedinfinity/go-commoner/pather"
)

func (t *Generator) GenerateModel(schema canonical.Canonical) error {
	renders, err := t.tm.RenderModel(schema)

	if err != nil {
		return err
	}

	for _, render := range renders {
		fmt.Println(render.Schema.SchemaId())

		if err := t.writeModel(render); err != nil {
			return err
		}
	}

	return nil
}

func (t *Generator) writeModel(output template_manager.ModelOutput) error {
	info := t.codeGenSchema.Info
	ns := util.SchemaNamepace(info, output.Schema)
	file := output.Path
	file = filepath.Base(file)
	file = extentioner.Strip(file)
	file = filepath.Base(ns) + "." + file
	path := ns
	path = strings.ReplaceAll(ns, info.Namespace.Get(), "")
	path = filepath.Dir(path)
	path = filepath.Join(t.destDir, info.RootDir.Get(), path, file)

	if err := pather.DirEnsure(filepath.Dir(path)); err != nil {
		return err
	}

	if err := os.WriteFile(path, output.Output, t.fileMode); err != nil {
		return err
	}

	return nil
}

func (t *Generator) dumpSchema(schema canonical.Canonical) string {
	title := strings.Repeat("=", 5)
	title = fmt.Sprintf("%v %v %v", title, schema.SchemaId(), title)
	sep := strings.Repeat("=", len(title))
	return fmt.Sprintf("%v\n%v", sep, title)
}
