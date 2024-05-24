package model

type CodeGenInherit struct {
	CodeGenCommon
	Id string
}

func (t CodeGenInherit) TypeId() string {
	return "inherit"
}

var _ CodeGenType = &CodeGenInherit{}
