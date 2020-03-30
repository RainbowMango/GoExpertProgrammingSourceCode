package errors

import (
	"errors"
	"fmt"
)

func CreateBasicError() error {
	err := errors.New("this is demo error")
	return fmt.Errorf("couldn't get unavailable numbers: %v", err)
}
