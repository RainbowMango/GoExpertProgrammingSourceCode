package constraint

type Integer interface {
	int
}

type MyInteger interface {
	Integer
}

type MyType int

type MyTypeIface interface {
	MyType
}

/*
// 不能嵌入类型参数中的类型
// Cannot embed a type parameter
type EmbeddedParameter[T any] interface {
	T
}
*/
