package overall

import (
	"fmt"
	"testing"
)

func TestGenericTypesInstantiation(t *testing.T) {
	var intVec Vector[int]       // 实例化为整型容器
	fmt.Print(intVec)            // 为了消除"unused variable"
	var stringVec Vector[string] // 实例化为字符串容器
	fmt.Print(stringVec)         // 为了消除"unused variable"
}

func ExampleVectorPush() {
	var v Vector[string]
	v.Push("Hello, World!")
	fmt.Print(v)
	// Output:
	// [Hello, World!]
}
