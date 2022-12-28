package loader_context

import (
	"boundedinfinity/codegen/codegen_type"
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTypeManager struct {
	order     []TypeLoaderContext
	id2Type   mapper.Mapper[string, TypeLoaderContext]
	path2Type mapper.Mapper[string, TypeLoaderContext]
	id2path   mapper.Mapper[string, string]
	path2id   mapper.Mapper[string, string]
	root2path mapper.Mapper[string, []string]
	path2root mapper.Mapper[string, string]
	root2Type mapper.Mapper[string, []TypeLoaderContext]
}

func TypeManager() *CodeGenTypeManager {
	return &CodeGenTypeManager{
		order:     make([]TypeLoaderContext, 0),
		id2Type:   make(mapper.Mapper[string, TypeLoaderContext], 0),
		path2Type: make(mapper.Mapper[string, TypeLoaderContext], 0),
		root2Type: make(mapper.Mapper[string, []TypeLoaderContext], 0),
		id2path:   make(mapper.Mapper[string, string], 0),
		path2id:   make(mapper.Mapper[string, string], 0),
		root2path: make(mapper.Mapper[string, []string], 0),
		path2root: make(mapper.Mapper[string, string], 0),
	}
}

func (t *CodeGenTypeManager) Register(lc TypeLoaderContext) error {
	if lc.Schema.Base().Id.Defined() {
		t.id2Type[lc.Schema.Base().Id.Get()] = lc
		t.id2path[lc.Schema.Base().Id.Get()] = lc.FileInfo.Source
		t.path2id[lc.FileInfo.Source] = lc.Schema.Base().Id.Get()
	}

	t.path2Type[lc.FileInfo.Source] = lc
	t.path2root[lc.FileInfo.Source] = lc.FileInfo.Root
	util.MapListAdd(t.root2Type, lc.FileInfo.Root, lc)
	util.MapListAdd(t.root2path, lc.FileInfo.Root, lc.FileInfo.Source)

	t.order = append(t.order, lc)

	return nil
}

func (t CodeGenTypeManager) All() []TypeLoaderContext {
	return t.order
}

func (t CodeGenTypeManager) Has(id string) bool {
	return t.id2Type.Has(id)
}

func (t CodeGenTypeManager) Find(id o.Option[string]) o.Option[TypeLoaderContext] {
	a := t.id2Type.Get(id.Get())
	b := t.path2Type.Get(id.Get())
	return o.FirstOf(a, b)
}

func (t CodeGenTypeManager) Resolve(schema codegen_type.CodeGenType) o.Option[TypeLoaderContext] {
	switch c := schema.(type) {
	case *codegen_type.CodeGenTypeRef:
		return t.Find(c.Ref)
	case *codegen_type.CodeGenTypeArray:
		return t.Resolve(c.Items)
	default:
		return t.Find(schema.Base().Id)
	}
}

func (t CodeGenTypeManager) FindSource(id o.Option[string]) o.Option[string] {
	if id.Empty() {
		return o.None[string]()
	}

	c := t.Find(id)

	if c.Empty() || c.Get().Schema.Base().Id.Empty() {
		return o.None[string]()
	}

	return t.id2path.Get(c.Get().Schema.Base().Id.Get())
}
