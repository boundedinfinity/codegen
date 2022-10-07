package uri_scheme

//go:generate enumer -path=./main.go

type Scheme string

const (
	File  Scheme = "file"
	Http  Scheme = "http"
	Https Scheme = "https"
)
