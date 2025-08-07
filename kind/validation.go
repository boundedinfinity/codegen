package kind

type Validatable interface {
	Validate() error
	HasValidation() bool
}
