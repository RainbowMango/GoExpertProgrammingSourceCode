package stringExample

import (
	"fmt"
	"unsafe"
)

// 用于模拟string的数据结构stringStruct
type dummyString struct {
	str unsafe.Pointer
	len int
}

// ByteSliceToString 用于测试byte切片强制转换成string
func ByteSliceToString() {
	srcSlice := []byte{'a', 'b', 'c'}

	// byte slice强制转成dummyString
	dstString := *(*dummyString)(unsafe.Pointer(&srcSlice))
	fmt.Printf("string len: %d", dstString.len) // 输出：string len 3

	// dummyString强制转成string
	realString := *(*string)(unsafe.Pointer(&dstString))
	fmt.Printf("Real string: %s", realString) // 输出：Real string: abc
}

func GetStringBySlice(s []byte) string {
	return string(s)
}

func GetSliceByString(str string) []byte {
	return []byte(str)
}

func RunStringPackage() {
	ByteSliceToString()
}
