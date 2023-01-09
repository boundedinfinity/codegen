package type_visibility

//go:generate enumer -path=./type_visability.go

type TypeVisibility string

const (
	Public    TypeVisibility = "public"
	Package   TypeVisibility = "package"
	Protected TypeVisibility = "protected"
	Private   TypeVisibility = "private"
)
