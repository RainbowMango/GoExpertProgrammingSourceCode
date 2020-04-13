package _string

import (
	"fmt"
	"strings"
)

func Contains() {
	s := "Hello World!"
	sub := "World"
	if strings.Contains(s, sub) {
		fmt.Printf("%s contains %s\n", s, sub)
	}
}
