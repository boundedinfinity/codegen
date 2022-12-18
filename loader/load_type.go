package loader

import (
	cp "boundedinfinity/codegen/codegen_project"
	"boundedinfinity/codegen/codegen_type"
	lc "boundedinfinity/codegen/loader_context"
	"boundedinfinity/codegen/util"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/environmenter"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-jsonschema/model"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *Loader) LoadTypePaths(paths ...string) error {
	paths = slicer.Map(paths, environmenter.Sub)
	paths = slicer.Map(paths, filepath.Clean)

	for _, path := range paths {
		ok, err := pather.IsFile(path)

		if err != nil {
			return err
		}

		m, err := marshaler.ReadFromPath(path)

		if err != nil {
			return err
		}

		if ok {
			lci := lc.LoaderFileInfo{
				Root:     pather.Dir(path),
				Source:   path,
				IsFile:   true,
				MimeType: m[path].MimeType,
			}

			if err := t.LoadTypePath(lci, m[path].Data); err != nil {
				return err
			}

			continue
		}

		for source, content := range m {
			lci := lc.LoaderFileInfo{
				Root:     path,
				Source:   source,
				IsFile:   true,
				MimeType: content.MimeType,
			}

			if err := t.LoadTypePath(lci, content.Data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) LoadTypePath(lci lc.LoaderFileInfo, data []byte) error {
	var bs []byte
	var err error

	switch lci.MimeType {
	case mime_type.ApplicationXYaml:
		bs, err = yaml.YAMLToJSON(data)

		if err != nil {
			return err
		}
	case mime_type.ApplicationJson:
		bs = data
	default:
		return cp.ErrMimeTypeUnsupportedv(lci.MimeType)
	}

	switch {
	case util.IsJsonSchemaFile(lci.Source):
		schema, err := model.UnmarshalSchema(bs)

		if err != nil {
			return err
		}

		if err = t.jsonSchemas.Register(lci.Root, lci.Source, schema); err != nil {
			return err
		}

		if err = t.ConvertJsonSchema(lci, schema); err != nil {
			return err
		}

	case util.IsCodeGenSchemaFile(lci.Source):
		var schema cp.CodeGenProjectProject

		if err := json.Unmarshal(bs, &schema); err != nil {
			return err
		}

		if schema.Info.DestDir.Defined() {
			schema.Info.DestDir = o.OfZ(util.EnsureAbs(pather.Dir(lci.Source), schema.Info.DestDir))
		}

		lc := lc.ProjectLoaderContext{
			FileInfo: lci,
			Project:  schema,
		}

		t.projectManager.Register(&lc)
	case util.IsCodeGenSchemaTypeFile(lci.Source):
		if schema, err := codegen_type.UnmarshalJson(bs); err != nil {
			return err
		} else {
			lc := lc.TypeLoaderContext{
				FileInfo: lci,
				Schema:   schema,
			}

			t.typeManager.Register(lc)
		}
	default:
		fmt.Printf("didn't process %v", lci.Source)
	}

	return nil
}
