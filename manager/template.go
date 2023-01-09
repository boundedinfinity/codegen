package manager

import (
	ct "boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/codegen_type/codegen_type_id"
	"boundedinfinity/codegen/codegen_type/template_type"
	"boundedinfinity/codegen/util"
	"fmt"

	"github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTemplateManager struct {
	All             []*ct.CodeGenProjectTemplateFile
	source2template mapper.Mapper[string, *ct.CodeGenProjectTemplateFile]
	root2template   mapper.Mapper[string, []*ct.CodeGenProjectTemplateFile]
	source2root     mapper.Mapper[string, string]
	root2source     mapper.Mapper[string, []string]
	tt2template     mapper.Mapper[template_type.TemplateType, []*ct.CodeGenProjectTemplateFile]
	tId2template    mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.CodeGenProjectTemplateFile]
}

func TemplateManager() *CodeGenTemplateManager {
	return &CodeGenTemplateManager{
		All:             make([]*ct.CodeGenProjectTemplateFile, 0),
		source2template: make(mapper.Mapper[string, *ct.CodeGenProjectTemplateFile]),
		source2root:     make(mapper.Mapper[string, string]),
		root2template:   make(mapper.Mapper[string, []*ct.CodeGenProjectTemplateFile]),
		root2source:     make(mapper.Mapper[string, []string]),
		tt2template:     make(mapper.Mapper[template_type.TemplateType, []*ct.CodeGenProjectTemplateFile]),
		tId2template:    make(mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.CodeGenProjectTemplateFile]),
	}
}

func (t *CodeGenTemplateManager) Register(meta *ct.CodeGenProjectTemplateFile) {
	t.All = append(t.All, meta)
	t.source2template[meta.SourcePath.Get()] = meta
	t.source2root[meta.SourcePath.Get()] = meta.RootPath.Get()
	util.MapListAdd(t.root2template, meta.RootPath.Get(), meta)
	util.MapListAdd(t.root2source, meta.RootPath.Get(), meta.SourcePath.Get())
	util.MapListAdd(t.tt2template, meta.TemplateType, meta)

	if meta.Type.Defined() {
		util.MapListAdd(t.tId2template, meta.Type.Get(), meta)
	}
}

func (t *CodeGenTemplateManager) Find(id any) optioner.Option[[]*ct.CodeGenProjectTemplateFile] {
	s := fmt.Sprintf("%v", id)

	r1 := t.tId2template.Get(codegen_type_id.CodgenTypeId(s))
	r2 := t.tt2template.Get(template_type.TemplateType(s))

	return optioner.FirstOf(r1, r2)
}
