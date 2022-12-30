package codegen_type

import "github.com/boundedinfinity/go-commoner/optioner"

func validateMinMax[T ~int | ~int64 | ~float64](s string, min, max optioner.Option[T]) error {
	if min.Defined() && max.Defined() {
		if max.Get() < min.Get() {
			return ErrMinMax
		}
	}

	return nil
}
