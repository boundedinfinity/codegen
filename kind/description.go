package kind

import "github.com/boundedinfinity/go-commoner/functional/optioner"

type Description struct {
	short optioner.Option[string]
	long  optioner.Option[string]
}
