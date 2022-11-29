package cacher

import (
	o "github.com/boundedinfinity/go-commoner/optioner"
)

func (t *Cacher) FindByGroup(group string) o.Option[[]*CachedData] {
	return t.groupMap.Get(group)
}

func (t *Cacher) FindBySource(path string) o.Option[*CachedData] {
	return t.sourceMap.Get(path)
}

func (t *Cacher) FindByDest(path string) o.Option[*CachedData] {
	return t.destMap.Get(path)
}

func (t *Cacher) Find(path string) o.Option[*CachedData] {
	return o.FirstOf(t.FindBySource(path), t.FindByDest(path))
}
