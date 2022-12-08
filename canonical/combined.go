package canonical

import (
	"fmt"

	o "github.com/boundedinfinity/go-commoner/optioner"
	"github.com/boundedinfinity/go-commoner/optioner/mapper"
)

type CanonicalCombined struct {
	typeMap mapper.Mapper[string, Canonical]
}

func NewCombinded() *CanonicalCombined {
	return &CanonicalCombined{
		typeMap: make(mapper.Mapper[string, Canonical], 0),
	}
}

func (t CanonicalCombined) Find(id string) o.Option[Canonical] {
	return o.FirstOf(t.FindById(id))
}

func (t CanonicalCombined) FindById(id string) o.Option[Canonical] {
	return t.typeMap.Get(id)
}

func (t CanonicalCombined) Id(c Canonical) o.Option[string] {
	var id o.Option[string]

	switch v := c.(type) {
	case CanonicalString:
		id = v.Id
	case CanonicalInteger:
		id = v.Id
	case CanonicalObject:
		id = v.Id
	}

	return id
}

func (t *CanonicalCombined) Register(c Canonical) error {
	id := t.Id(c)

	if id.Empty() {
		return nil
	}

	if t.typeMap.Has(id.Get()) {
		fmt.Printf("already contains %v", id.Get())
	}

	t.typeMap[id.Get()] = c

	return nil
}
