package compare

import "fmt"

func StructTypeCompare() {
	type Student struct {
		name string
		age  int
	}

	zhangsan := Student{name: "Zhang San", age: 18}
	lisi := Student{name: "Li Si", age: 18}

	fmt.Println(zhangsan == lisi) // 输出 false
}

/*
func StructTypeCompare2() {
	type Student struct {
		name  string
		grade map[string]int
	}

	zhangsan := Student{name: "Zhang San"}
	lisi := Student{name: "Li Si"}

	fmt.Println(zhangsan == lisi) // Invalid operation: the operator == is not defined on Student
}
*/

type Animal interface {
	Speak() string
}

type Duck struct {
	Name string
}

func (a Duck) Speak() string {
	return "I'm " + a.Name
}

type Cat struct {
	Name string
}

func (a Cat) Speak() string {
	return "I'm " + a.Name
}

func InterfaceTypeCompare() {
	var d1, d2, c1 Animal
	d1 = Duck{Name: "Donald Duck"}
	d2 = Duck{Name: "Donald Duck"}
	fmt.Println(d1 == d2) // 输出 true

	c1 = Cat{Name: "Donald Duck"} // 伪装的 duck
	fmt.Println(d1 == c1)         // 输出 false
}
