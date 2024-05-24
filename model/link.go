package model

type CodeGenLink struct {
	CodeGenCommon
	Id string
}

func (t CodeGenLink) TypeId() string {
	return "inherit"
}

var _ CodeGenType = &CodeGenLink{}
