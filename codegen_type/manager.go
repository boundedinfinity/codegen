package codegen_type

import (
	"boundedinfinity/codegen/util"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CodeGenTypeManager struct {
	order     []CodeGenType
	id2Type   mapper.Mapper[string, CodeGenType]
	path2Type mapper.Mapper[string, CodeGenType]
	id2path   mapper.Mapper[string, string]
	path2id   mapper.Mapper[string, string]
	root2path mapper.Mapper[string, []string]
	path2root mapper.Mapper[string, string]
	root2Type mapper.Mapper[string, []CodeGenType]
}

func Manager() *CodeGenTypeManager {
	return &CodeGenTypeManager{
		order:     make([]CodeGenType, 0),
		id2Type:   make(mapper.Mapper[string, CodeGenType], 0),
		path2Type: make(mapper.Mapper[string, CodeGenType], 0),
		root2Type: make(mapper.Mapper[string, []CodeGenType], 0),
		id2path:   make(mapper.Mapper[string, string], 0),
		path2id:   make(mapper.Mapper[string, string], 0),
		root2path: make(mapper.Mapper[string, []string], 0),
		path2root: make(mapper.Mapper[string, string], 0),
	}
}

func (t *CodeGenTypeManager) Register(root, path string, schema CodeGenType) error {
	if schema == nil {
		return nil
	}

	b := schema.Base()
	b.Source = path
	b.Root = root

	if b.Id.Defined() {
		t.id2Type[b.Id.Get()] = schema
		t.id2path[b.Id.Get()] = path
		t.path2id[path] = b.Id.Get()
	}

	t.path2Type[path] = schema
	t.path2root[path] = root
	util.MapListAdd(t.root2Type, root, schema)
	util.MapListAdd(t.root2path, root, path)

	t.order = append(t.order, schema)

	return nil
}

func (t CodeGenTypeManager) All() []CodeGenType {
	return t.order
}

func (t CodeGenTypeManager) Has(id string) bool {
	return t.id2Type.Has(id)
}

func (t CodeGenTypeManager) Find(id o.Option[string]) o.Option[CodeGenType] {
	return o.FirstOf(t.id2Type.Get(id.Get()), t.path2Type.Get(id.Get()))
}

func (t CodeGenTypeManager) FindSource(id o.Option[string]) o.Option[string] {
	if id.Empty() {
		return o.None[string]()
	}

	c := t.Find(id)

	if c.Empty() || c.Get().Base().Id.Empty() {
		return o.None[string]()
	}

	return t.id2path.Get(c.Get().Base().Id.Get())
}
