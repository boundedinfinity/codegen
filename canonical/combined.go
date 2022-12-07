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

func (t *CanonicalCombined) Register(c Canonical) error {
	var typ string

	switch v := c.(type) {
	case CanonicalString:
		typ = v.Id.Get()
	case CanonicalInteger:
		typ = v.Id.Get()
	case CanonicalObject:
		typ = v.Id.Get()
	}

	if t.typeMap.Has(typ) {
		fmt.Printf("already have %v", typ)
	}

	t.typeMap[typ] = c

	return nil
}
