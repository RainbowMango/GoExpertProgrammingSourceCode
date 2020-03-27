package errors

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

var _ error = &NotFoundError{}

func NewNotFoundError(name string) error {
	return &NotFoundError{Name: name}
}
