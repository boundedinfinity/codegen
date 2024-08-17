package model

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

func setOptionAndReturn[B any, V comparable](builder B, opt *optioner.Option[V], value V) B {
	*opt = optioner.OfZero(value)
	return builder
}

func appendOptionAndReturn[B any, V comparable](builder B, opt *optioner.Option[[]V], value []V) B {
	var new []V
	new = append(new, opt.Get()...)
	new = append(new, value...)
	*opt = optioner.OfSlice(new)
	return builder
}

func mergeDescription(d1, d2 optioner.Option[string]) optioner.Option[string] {
	if d1.Defined() && d2.Empty() {
		return d1
	}

	if d1.Empty() && d2.Defined() {
		return d2
	}

	if d1.Defined() && d2.Defined() {
		desc := fmt.Sprintf("%v\n%v", d2.Get(), d1.Get())
		return optioner.Some(desc)
	}

	return optioner.None[string]()
}

func SetV[T any, V any](t T, c *V, n V) T {
	*c = n
	return t
}

func SetO[T any, V any](t T, c *optioner.Option[V], n V) T {
	*c = optioner.Some(n)
	return t
}

func EnsureName(typ CodeGenSchema) {
	if typ.Common().Lang.Name.Empty() {
		var name string

		if typ.Common().Name.Defined() {
			name = typ.Common().Name.Get()
		} else {
			name = typ.Common().Id.Get()
			name = pather.Paths.Base(name)
			name = caser.KebabToPascal(name)
		}

		typ.Common().Lang.Name = optioner.Some(name)
	}
}

func EnsurePackage(typ CodeGenSchema) {
	if typ.Common().Lang.Import.Empty() {
		pkg := typ.Common().Id.Get()
		pkg = pather.Paths.Dir(pkg)
		pkg = pather.Paths.Base(pkg)
		typ.Common().Lang.Import = optioner.Some(pkg)
	}
}

func EnsureJsonName(typ CodeGenSchema) {
	if typ.Common().JsonName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		typ.Common().JsonName = optioner.Some(name)
	}
}

func EnsureYamlName(typ CodeGenSchema) {
	if typ.Common().YamlName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		typ.Common().YamlName = optioner.Some(name)
	}
}

func EnsureSqlName(typ CodeGenSchema) {
	if typ.Common().SqlName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		name = caser.KebabToSnake(name)
		typ.Common().SqlName = optioner.Some(name)
	}
}
