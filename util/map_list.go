package util

import "github.com/boundedinfinity/go-commoner/optioner/mapper"

func MapListAdd[K comparable, V any](m mapper.Mapper[K, []V], k K, v V) {
	if !m.Has(k) {
		m[k] = make([]V, 0)
	}

	m[k] = append(m[k], v)
}
