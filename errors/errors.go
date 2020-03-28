package errors

import (
	"errors"
	"fmt"
)

type NotFoundError struct {
	Name string
}

func (e *NotFoundError) Error() string { return e.Name + ": not found" }

var _ error = &NotFoundError{}

func NewNotFoundError(name string) error {
	return &NotFoundError{Name: name}
}

func MakeByErrorsNew() error {
	return errors.New("new error")
}

func MakeByFmtErrorf() error {
	return fmt.Errorf("new error")
}
