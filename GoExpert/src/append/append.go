package append

import (
	"errors"
	"fmt"
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

func AppendDemo3() {
	x := make([]int, 0, 10)
	x = append(x, 1, 2, 3)
	y := append(x, 4)
	z := append(x, 5)
	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)
}
