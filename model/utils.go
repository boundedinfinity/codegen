package model

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/caser"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
)

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

func EnsureName(typ CodeGenType) {
	if typ.Common().Name.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		name = caser.KebabToPascal(name)
		typ.Common().Name = optioner.Some(name)
	}
}

func EnsurePackage(typ CodeGenType) {
	if typ.Common().Package.Empty() {
		pkg := typ.Common().Id.Get()
		pkg = pather.Paths.Dir(pkg)
		pkg = pather.Paths.Base(pkg)
		typ.Common().Package = optioner.Some(pkg)
	}
}

func EnsureImportPath(project CodeGenProject, typ CodeGenType) {
	if typ.Common().ImportPath.Empty() {
		pkg := typ.Common().Id.Get()
		pkg = pather.Paths.Dir(pkg)
		pkg = pather.Paths.Join(project.Package.Get(), pkg)
		typ.Common().ImportPath = optioner.Some(pkg)
	}
}

func EnsureJsonName(typ CodeGenType) {
	if typ.Common().JsonName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		typ.Common().JsonName = optioner.Some(name)
	}
}

func EnsureYamlName(typ CodeGenType) {
	if typ.Common().YamlName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		typ.Common().YamlName = optioner.Some(name)
	}
}

func EnsureSqlName(typ CodeGenType) {
	if typ.Common().SqlName.Empty() {
		name := typ.Common().Id.Get()
		name = pather.Paths.Base(name)
		name = caser.KebabToSnake(name)
		typ.Common().SqlName = optioner.Some(name)
	}
}
