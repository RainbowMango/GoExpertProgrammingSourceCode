package compare

import "fmt"

func StructTypeCompare() {
	type Student struct {
		name string
		age  int
	}

	zhangsan := Student{name: "Zhang San", age: 18}
	lisi := Student{name: "Li Si", age: 18}

	fmt.Println(zhangsan == lisi) // false
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
