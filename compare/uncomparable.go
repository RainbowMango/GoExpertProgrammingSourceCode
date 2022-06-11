package compare

import "fmt"

/*
func SliceCompare() {
	s1 := make([]int, 3)
	s2 := make([]int, 3)
	fmt.Println(s1 == s2) // the operator == is not defined on []int
}
*/

type Bird struct {
	Name      string
	SpeakFunc func() string
}

func (a Bird) Speak() string {
	return "I'm " + a.SpeakFunc()
}
func UncomparablePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic")
		}
	}()

	var b1 Animal = Bird{
		Name: "bird",
		SpeakFunc: func() string {
			return "I'm Poly"
		}}
	var b2 Animal = Bird{
		Name: "bird",
		SpeakFunc: func() string {
			return "I'm eagle"
		}}

	fmt.Println(b1 == b2) // panic: comparing uncomparable type compare.Bird
}
