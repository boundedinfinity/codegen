package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/pather"
	"github.com/boundedinfinity/go-commoner/idiomatic/stringer"
)

func (t *Processor) calculatePackageNames() error {
	translations := map[string]string{}

	for _, typ := range t.combined.Types {
		if err := t.calculatePackageName(typ, translations); err != nil {
			return err
		}
	}

	return nil
}

func (t *Processor) calculatePackageName(typ model.CodeGenType, translations map[string]string) error {
	var pkg string

	if typ.Common().Package.Defined() {
		pkg = typ.Common().Package.Get()
	} else {
		if t.combined.Package.Defined() {
			pkg = t.combined.Package.Get()
		}

		pkg = pather.Paths.Join(pkg, typ.TypeId().Get())
		pkg = pather.Paths.Dir(pkg)
	}

	for from, to := range translations {
		pkg = stringer.Replace(pkg, to, from)
	}

	typ.Common().Package = optioner.Some(pkg)

	return nil
}
