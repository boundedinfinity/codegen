package model

type CodeGenObject struct {
	CodeGenCommon
	Properties []CodeGenType
}

func (t CodeGenObject) TypeId() string {
	return "object"
}

var _ CodeGenType = &CodeGenObject{}
