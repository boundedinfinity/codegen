package model

type WalkType string

const (
	WALKTYPE_NAMESPACE WalkType = "namespace"
	WALKTYPE_MODEL     WalkType = "model"
	WALKTYPE_PROPERTY  WalkType = "property"
	WALKTYPE_OPERATION WalkType = "operation"
)

type WalkFunc func(ctx *WalkContext) error

type WalkContext struct {
	// Model     InputModel
	Operation InputOperation
}
