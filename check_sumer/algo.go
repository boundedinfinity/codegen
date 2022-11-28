package check_sumer

//go:generate enumer -path=./algo.go

type CheckSumAlgo string

const (
	Md5    CheckSumAlgo = "md5"
	Sha1   CheckSumAlgo = "sha1"
	Sha256 CheckSumAlgo = "sha256"
)
