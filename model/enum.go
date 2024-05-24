package model

type Enum struct {
	CodeGenCommon
}

func (t Enum) TypeId() string {
	return "enum"
}

var _ CodeGenType = &Enum{}
