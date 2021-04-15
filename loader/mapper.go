package loader

import "boundedinfinity/codegen/util"

func (t *Loader) processMapper() error {
	if err := util.UnmarshalFromFile(t.typeMapPath, &t.Mapper); err != nil {
		return err
	}

	return nil
}
