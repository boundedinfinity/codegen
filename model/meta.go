package model

///////////////////////////////////////////////////////////////////
// CodeGenMeta
//////////////////////////////////////////////////////////////////

type CodeGenMeta struct {
}

//----------------------------------------------------------------
// Merge
//----------------------------------------------------------------

func (t *CodeGenMeta) Merge(obj CodeGenMeta) error {
	return nil
}

//----------------------------------------------------------------
// Validate
//----------------------------------------------------------------

func (t *CodeGenMeta) Validate() error {
	return nil
}

func (t *CodeGenMeta) HasValidation() bool {
	return false
}
