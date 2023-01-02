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
			lci := ct.SourceMeta{
				RootPath:       o.Some(pather.Dir(path)),
				SourcePath:     o.Some(path),
				SourceMimeType: m[path].MimeType,
			}

			if err := t.LoadTypePath(lci, m[path].Data); err != nil {
				return err
			}

			continue
		}

		for source, content := range m {
			lci := ct.SourceMeta{
				RootPath:       o.Some(path),
				SourcePath:     o.Some(source),
				SourceMimeType: content.MimeType,
			}

			if err := t.LoadTypePath(lci, content.Data); err != nil {
				return err
			}
		}
	}

	return nil
}

func (t *Loader) LoadTypePath(lci ct.SourceMeta, data []byte) error {
	var bs []byte
	var err error

	switch lci.SourceMimeType {
	case mime_type.ApplicationXYaml:
		bs, err = yaml.YAMLToJSON(data)

		if err != nil {
			return err
		}
	case mime_type.ApplicationJson:
		bs = data
	default:
		return ct.ErrMimeTypeUnsupportedv(lci.SourceMimeType)
	}

	switch {
	case util.IsCodeGenSchemaFile(lci.SourcePath.Get()):
		var schema ct.CodeGenProject

		if err := json.Unmarshal(bs, &schema); err != nil {
			return err
		}

		if schema.Info.DestDir.Defined() {
			schema.Info.DestDir = o.OfZ(util.EnsureAbs(pather.Dir(lci.SourcePath.Get()), schema.Info.DestDir))
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
				ProjectContext: &ctx,
				Name:           operation.Name,
				Description:    operation.Description,
				Input:          operation.Input,
				Output:         operation.Output,
			}

			if err := t.projectManager.RegisterOperation(&opCtx); err != nil {
				return err
			}
		}

	case util.IsCodeGenSchemaTypeFile(lci.SourcePath.Get()):
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
	case util.IsJsonSchemaFile(lci.SourcePath.Get()):
		js, err := model.UnmarshalSchema(bs)

		if err != nil {
			return err
		}

		if err = t.jsonSchemas.Register(lci.RootPath.Get(), lci.SourcePath.Get(), js); err != nil {
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
		fmt.Printf("didn't process %v", lci.SourcePath)
	}

	return nil
}
