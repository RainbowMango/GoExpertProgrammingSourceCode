package recover

import "fmt"

// 测试recover能否在具名函数中使用
func ForRecover() {
	if p := recover(); p != nil {
		fmt.Printf("catch a recover from named-defer-function\n")
	}
}

func NamedRecover() {
	defer ForRecover()

	panic("this is a panic")
}
