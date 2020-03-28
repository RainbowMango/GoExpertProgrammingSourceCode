package errors

import (
	"errors"
	"fmt"
	"os"
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

func AssertError(err error) {
	if e, ok := err.(*os.PathError); ok {
		fmt.Printf("it's an os.PathError, operation: %s, path: %s, msg: %v", e.Op, e.Path, e.Err)
	}
}
