package _range

import "fmt"

func ForExpression() {
	s := []int{1, 2, 3}

	for i := 0; i < len(s); i++ {
		fmt.Println(s[i])
	}
}

func ForRangeExpression() {
	s := []int{1, 2, 3}

	for i := range s {
		fmt.Println(s[i])
	}
}
