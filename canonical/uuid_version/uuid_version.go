package uuid_version

//go:generate enumer -path=./uuid_version.go

type UuidVersion string

const (
	V1 UuidVersion = "v1"
	V2 UuidVersion = "v2"
	V3 UuidVersion = "v3"
	V4 UuidVersion = "v4"
	V5 UuidVersion = "v5"
)
