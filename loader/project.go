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

func typePath(project *ct.CodeGenProject, typ ct.CodeGenType) error {
	*typ.Source() = project.SourceMeta
	return nil
}

func operationPath(project *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
	operation.SourceMeta = project.SourceMeta
	return nil
}

func normalizePath(root o.Option[string], paths ...string) ([]string, error) {
	var ns []string
	paths = slicer.Dedup(paths)

	for _, path := range paths {
		path = environmenter.Sub(path)
		path = filepath.Clean(path)

		if !filepath.IsAbs(path) && root.Defined() {
			if abs, err := filepath.Abs(path); err != nil {
				return ns, err
			} else {
				path = abs
			}
		}

		ns = append(ns, path)
	}

	return ns, nil
}

func (t *Loader) LoadProjectPath(root o.Option[string], paths ...string) ([]*codegen_type.CodeGenProject, error) {
	var projects []*codegen_type.CodeGenProject

	paths, err := normalizePath(root, paths...)

	if err != nil {
		return projects, err
	}

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

			if project, err := t.loadProjectPath(sourceMeta, m[path].Data); err != nil {
				return projects, err
			} else {
				projects = append(projects, &project)
			}

			continue
		}

		for source, content := range m {
			sourceMeta := ct.SourceMeta{
				RootPath:       o.Some(path),
				SourcePath:     o.Some(source),
				SourceMimeType: content.MimeType,
			}

			if project, err := t.loadProjectPath(sourceMeta, content.Data); err != nil {
				return projects, err
			} else {
				projects = append(projects, &project)
			}
		}
	}

	var typeProjects []*codegen_type.CodeGenProject

	typeLoad := func(project *ct.CodeGenProject, typ ct.CodeGenType) error {
		switch c := typ.(type) {
		case *ct.CodeGenTypePath:
			if c.SourcePath.Defined() {
				if ps, err := t.LoadProjectPath(project.RootPath, c.SourcePath.Get()); err != nil {
					return err
				} else {
					typeProjects = append(typeProjects, ps...)
				}
			}
		}

		return nil
	}

	var operationProjects []*codegen_type.CodeGenProject

	operationLoad := func(project *ct.CodeGenProject, operation *ct.CodeGenProjectOperation) error {
		if operation.SourcePath.Defined() {
			if ps, err := t.LoadProjectPath(project.RootPath, operation.SourcePath.Get()); err != nil {
				return err
			} else {
				operationProjects = append(operationProjects, ps...)
			}
		}

		return nil
	}

	if err := ct.WalkType(typeLoad, projects...); err != nil {
		return projects, err
	}

	if err := ct.WalkType(typePath, typeProjects...); err != nil {
		return projects, err
	}

	if err := ct.WalkType(typePath, projects...); err != nil {
		return projects, err
	}

	if err := ct.WalkOperation(operationLoad, projects...); err != nil {
		return projects, err
	}

	if err := ct.WalkOperation(operationPath, operationProjects...); err != nil {
		return projects, err
	}

	if err := ct.WalkOperation(operationPath, projects...); err != nil {
		return projects, err
	}

	projects = append(projects, typeProjects...)
	projects = append(projects, operationProjects...)

	return projects, nil
}

func (t *Loader) loadProjectPath(sourceMeta ct.SourceMeta, data []byte) (codegen_type.CodeGenProject, error) {
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
