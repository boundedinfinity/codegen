package loader

import "github.com/boundedinfinity/go-commoner/optioner"

func (t *Loader) FindSource(id optioner.Option[string]) optioner.Option[string] {
	if id.Empty() {
		optioner.None[string]()
	}

	return t.typeManager.FindSource(id)
}
