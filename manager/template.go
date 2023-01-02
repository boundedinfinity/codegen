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
	All             []*ct.TemplateContext
	source2template mapper.Mapper[string, *ct.TemplateContext]
	root2template   mapper.Mapper[string, []*ct.TemplateContext]
	source2root     mapper.Mapper[string, string]
	root2source     mapper.Mapper[string, []string]
	tt2template     mapper.Mapper[template_type.TemplateType, []*ct.TemplateContext]
	tId2template    mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.TemplateContext]
}

func TemplateManager() *CodeGenTemplateManager {
	return &CodeGenTemplateManager{
		All:             make([]*ct.TemplateContext, 0),
		source2template: make(mapper.Mapper[string, *ct.TemplateContext]),
		source2root:     make(mapper.Mapper[string, string]),
		root2template:   make(mapper.Mapper[string, []*ct.TemplateContext]),
		root2source:     make(mapper.Mapper[string, []string]),
		tt2template:     make(mapper.Mapper[template_type.TemplateType, []*ct.TemplateContext]),
		tId2template:    make(mapper.Mapper[codegen_type_id.CodgenTypeId, []*ct.TemplateContext]),
	}
}

func (t *CodeGenTemplateManager) Register(lc *ct.TemplateContext) {
	t.All = append(t.All, lc)
	t.source2template[lc.FileInfo.SourcePath.Get()] = lc
	t.source2root[lc.FileInfo.SourcePath.Get()] = lc.FileInfo.RootPath.Get()
	util.MapListAdd(t.root2template, lc.FileInfo.RootPath.Get(), lc)
	util.MapListAdd(t.root2source, lc.FileInfo.RootPath.Get(), lc.FileInfo.SourcePath.Get())
	util.MapListAdd(t.tt2template, lc.TemplateType, lc)
	util.MapListAdd(t.tId2template, lc.TypeId, lc)
}

func (t *CodeGenTemplateManager) Find(id any) optioner.Option[[]*ct.TemplateContext] {
	s := fmt.Sprintf("%v", id)

	return optioner.FirstOf(
		t.tId2template.Get(codegen_type_id.CodgenTypeId(s)),
		t.tt2template.Get(template_type.TemplateType(s)),
	)
}
