package cacher

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Cacher) FindSingle(id string) o.Option[*CachedData] {
	return o.FirstOf(t.FindBySource(id), t.FindByDest(id))
}

func (t *Cacher) FindList(id string) o.Option[[]*CachedData] {
	return o.FirstOf(t.FindByOrig(id))
}

func (t *Cacher) FindByOrig(id string) o.Option[[]*CachedData] {
	return t.orig2Data.Get(id)
}

func (t *Cacher) FindBySource(id string) o.Option[*CachedData] {
	return t.source2Data.Get(id)
}

func (t *Cacher) FindByDest(id string) o.Option[*CachedData] {
	return t.dest2Data.Get(id)
}
