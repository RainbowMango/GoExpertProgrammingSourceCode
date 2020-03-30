package errors

import (
	"errors"
	"fmt"
)

func ExampleCreateBasicError() {
	err := errors.New("this is demo error")

	basicErr := fmt.Errorf("some context: %v", err)           // 使用 %v
	if _, ok := basicErr.(interface{ Unwrap() error }); !ok { // 如果errBasic没有实现Unwrap接口
		fmt.Print("basicErr is a errorString")
	}
	// Output:
	// basicErr is a errorString
}

func ExampleCreateWrapError() {
	err := errors.New("this is demo error")

	wrapError := fmt.Errorf("some context: %w", err)          // 使用%w
	if _, ok := wrapError.(interface{ Unwrap() error }); ok { // 如果wrapError实现了Unwrap接口
		fmt.Print("wrapError is a wrapError")
	}
	// Output:
	// wrapError is a wrapError
}
