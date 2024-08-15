package processor

import (
	"boundedinfinity/codegen/model"

	"github.com/boundedinfinity/go-commoner/errorer"
	"github.com/boundedinfinity/go-commoner/functional/optioner"
	"github.com/boundedinfinity/go-commoner/idiomatic/slicer"
)

var (
	ErrCodeGenPackageEmpty       = errorer.New("package empty")
	ErrCodeGenPackageMoreThanOne = errorer.New("more than one package found")
)

func (this *Processor) processProjectPackage(projects ...*model.CodeGenProject) error {
	filtered := slicer.Filter(
		func(_ int, p *model.CodeGenProject) bool { return p.Package.Defined() },
		projects...,
	)

	switch len(filtered) {
	case 0:
		return ErrCodeGenPackageEmpty
	case 1:
		this.Combined.Package = filtered[0].Package
	default:
		pkgs := slicer.Map(
			func(_ int, p *model.CodeGenProject) string { return p.Package.Get() },
			filtered...,
		)
		return ErrCodeGenPackageMoreThanOne.WithValue(slicer.Join(",", pkgs))
	}

	return nil
}

var (
	ErrCodeGenOutputRootMoreThanOne = errorer.New("more than one output root found")
)

func (this *Processor) processProjectOutputRoot(projects ...*model.CodeGenProject) error {
	filtered := slicer.Filter(
		func(_ int, p *model.CodeGenProject) bool { return p.OutputRoot.Defined() },
		projects...,
	)

	switch len(filtered) {
	case 0:
		this.Combined.OutputRoot = optioner.Some(".")
	case 1:
		this.Combined.OutputRoot = filtered[0].OutputRoot
	default:
		pkgs := slicer.Map(
			func(_ int, p *model.CodeGenProject) string { return p.Package.Get() },
			filtered...,
		)
		return ErrCodeGenOutputRootMoreThanOne.WithValue(slicer.Join(",", pkgs))
	}

	return nil
}
