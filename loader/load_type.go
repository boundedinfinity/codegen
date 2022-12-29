package loader

import (
	ct "boundedinfinity/codegen/codegen_type"
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
			lci := ct.LoaderFileInfo{
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
			lci := ct.LoaderFileInfo{
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

func (t *Loader) LoadTypePath(lci ct.LoaderFileInfo, data []byte) error {
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
		return ct.ErrMimeTypeUnsupportedv(lci.MimeType)
	}

	switch {
	case util.IsCodeGenSchemaFile(lci.Source):
		var schema ct.CodeGenProjectProject

		if err := json.Unmarshal(bs, &schema); err != nil {
			return err
		}

		if schema.Info.DestDir.Defined() {
			schema.Info.DestDir = o.OfZ(util.EnsureAbs(pather.Dir(lci.Source), schema.Info.DestDir))
		}

		ctx := ct.ProjectContext{
			FileInfo: lci,
			Project:  schema,
		}

		if err := t.projectManager.RegisterProject(&ctx); err != nil {
			return err
		}

		for _, operation := range ctx.Project.Operations {
			opCtx := ct.OperationContext{
				ProjectLoaderContext: &ctx,
				Name:                 operation.Name,
				Description:          operation.Description,
				Input:                operation.Input,
				Output:               operation.Output,
			}

			if err := t.projectManager.RegisterOperation(&opCtx); err != nil {
				return err
			}
		}

	case util.IsCodeGenSchemaTypeFile(lci.Source):
		var schema ct.CodeGenType

		if err := ct.UnmarshalJson(bs, &schema); err != nil {
			return err
		} else {
			lc := ct.CodeGenTypeContext{
				FileInfo: lci,
				Schema:   schema,
			}

			if err := t.typeManager.Register(lc); err != nil {
				return err
			}
		}
	case util.IsJsonSchemaFile(lci.Source):
		js, err := model.UnmarshalSchema(bs)

		if err != nil {
			return err
		}

		if err = t.jsonSchemas.Register(lci.Root, lci.Source, js); err != nil {
			return err
		}

		lc := ct.CodeGenTypeContext{
			FileInfo: lci,
		}

		if err = t.ConvertJsonSchema(&lc, js); err != nil {
			return err
		}

		if err := t.typeManager.Register(lc); err != nil {
			return err
		}
	default:
		fmt.Printf("didn't process %v", lci.Source)
	}

	return nil
}
