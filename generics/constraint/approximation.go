package constraint

import "fmt"

type MyString string

// AnyString 定义了一个宽泛的类型集合，只要类型底层类型为string即可
// 如果没有使用`~`，则即便底层类型为string，也会有编译错误：
// Type does not implement constraint 'AnyString' because type is not included in type set ('string')
type AnyString interface {
	~string
}

func SayHi[T AnyString](name T) {
	fmt.Printf("Hi %s", name)
}

/*
不能使用interface表示底层类型
type InterfaceIsNotAllowed interface {
	~AnyString // 编译错误：Invalid use of ~ ('AnyString' is an interface)
}
*/

/*
底层类型不是自己的，不能用于~之后
type UnderlyingTypeNotMatched interface {
	~MyString // Invalid use of ~ (underlying type of 'MyString' is not 'MyString')
}
*/
