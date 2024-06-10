package model

import (
	"fmt"

	"github.com/boundedinfinity/go-commoner/functional/optioner"
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
