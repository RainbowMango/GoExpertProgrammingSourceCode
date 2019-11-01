package append

import (
	"errors"
	"os"
)

func Validation() []error {
	var errs []error

	errs = append(errs, errors.New("error 1"))
	errs = append(errs, errors.New("error 2"))
	errs = append(errs, errors.New("error 3"))

	return errs
}

func ValidateName(name string) error {
	if name != "" {
		return nil
	}

	return errors.New("empty name")
}

func Validations(name string) []error {
	var errs []error

	errs = append(errs, ValidateName(name))

	return errs
}

func foo() {
	errs := Validations("")

	if len(errs) > 0 {
		println(errs)
		os.Exit(1)
	}
}
