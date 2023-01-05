package loader

import (
	"boundedinfinity/codegen/codegen_type"
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"
	"encoding/json"
	"fmt"
	"path/filepath"

	"github.com/boundedinfinity/go-commoner/environmenter"
	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/pather"
	"github.com/boundedinfinity/go-commoner/slicer"
	"github.com/boundedinfinity/go-marshaler"
	"github.com/boundedinfinity/go-mimetyper/mime_type"
	"github.com/ghodss/yaml"
)

func (t *Loader) LoadProjectPaths(paths ...string) ([]codegen_type.CodeGenProject, error) {
	var projects []codegen_type.CodeGenProject

	paths = slicer.Map(paths, environmenter.Sub)
	paths = slicer.Map(paths, filepath.Clean)

	for _, path := range paths {
		ok, err := pather.IsFile(path)

		if err != nil {
			return projects, err
		}

		m, err := marshaler.ReadFromPath(path)

		if err != nil {
			return projects, err
		}

		if ok {
			sourceMeta := ct.SourceMeta{
				RootPath:       o.Some(pather.Dir(path)),
				SourcePath:     o.Some(path),
				SourceMimeType: m[path].MimeType,
			}

			if project, err := t.LoadTypePath(sourceMeta, m[path].Data); err != nil {
				return projects, err
			} else {
				projects = append(projects, project)
			}

			continue
		}

		for source, content := range m {
			sourceMeta := ct.SourceMeta{
				RootPath:       o.Some(path),
				SourcePath:     o.Some(source),
				SourceMimeType: content.MimeType,
			}

			if project, err := t.LoadTypePath(sourceMeta, content.Data); err != nil {
				return projects, err
			} else {
				projects = append(projects, project)
			}
		}
	}

	return projects, nil
}

func (t *Loader) LoadTypePath(sourceMeta ct.SourceMeta, data []byte) (codegen_type.CodeGenProject, error) {
	var project codegen_type.CodeGenProject
	var bs []byte
	var err error

	switch sourceMeta.SourceMimeType {
	case mime_type.ApplicationXYaml:
		bs, err = yaml.YAMLToJSON(data)

		if err != nil {
			return project, err
		}
	case mime_type.ApplicationJson:
		bs = data
	default:
		return project, ct.ErrMimeTypeUnsupportedv(sourceMeta.SourceMimeType)
	}

	switch {
	case util.IsCodeGenFile(sourceMeta.SourcePath.Get()):
		if err := json.Unmarshal(bs, &project); err != nil {
			return project, err
		}

		project.SourceMeta = sourceMeta
		var operations []codegen_type.CodeGenProjectOperation

		for _, operation := range project.Operations {
			operation.SourceMeta = sourceMeta
			operations = append(operations, operation)
		}

		project.Operations = operations

	case util.IsJsonSchemaFile(sourceMeta.SourcePath.Get()):
		// js, err := model.UnmarshalSchema(bs)

		// if err != nil {
		// 	return err
		// }

		// if err = t.jsonSchemas.Register(lci.RootPath.Get(), lci.SourcePath.Get(), js); err != nil {
		// 	return err
		// }

		// lc := ct.CodeGenType{
		// 	FileInfo: lci,
		// }

		// if err = t.ConvertJsonSchema(&lc, js); err != nil {
		// 	return err
		// }

		// if err := t.typeManager.Register(lc); err != nil {
		// 	return err
		// }
	default:
		fmt.Printf("didn't process %v", sourceMeta.SourcePath)
	}

	return project, nil
}
