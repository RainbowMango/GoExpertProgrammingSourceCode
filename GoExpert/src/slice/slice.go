package slice

import (
	"fmt"
	"reflect"
	"time"
)

// simple slice expression(a[low, high]) 切取字符串时，生成的仍然为字符串
func SliceString() {
	baseStr := "Hello World!"
	fmt.Printf("baseStr: %s\n", baseStr)                      // baseStr: Hello World!
	fmt.Printf("baseStr type: %s\n", reflect.TypeOf(baseStr)) // baseStr type: string

	newStr := baseStr[0:5]
	fmt.Printf("newStr: %s\n", newStr)                      // newStr: Hello
	fmt.Printf("newStr type: %v\n", reflect.TypeOf(newStr)) // newStr type: string
}

// simple slice expression(a[low, high]) 切取数组时，生成新的切片
func SliceArray() {
	baseArray := [10]string{"Hello", "World"}
	fmt.Printf("baseArray: %s\n", baseArray)                      // baseArray: [Hello World        ]
	fmt.Printf("baseArray type: %s\n", reflect.TypeOf(baseArray)) // baseArray type: [10]string

	newSlice := baseArray[0:5]
	fmt.Printf("newSlice: %v\n", newSlice)                      // newSlice: [Hello World   ]
	fmt.Printf("newSlice type: %v\n", reflect.TypeOf(newSlice)) // newSlice type: []string
}

// simple slice expression(a[low, high]) 切取切片时，low, high取值关系为： 0 <= low <= high <= cap(a)
// 特别需要注意的是low、high取值均可以超越len(a)
func SliceCap() {
	baseSlice := make([]int, 0, 10)
	newSlice := baseSlice[2:5]
	fmt.Printf("newSlice: %v", newSlice) // newSlice: [0 0 0]
}

// extend slice expression(a[low, high, max])切取数组时，max用于限制新生成片的capacity。
func ExtendSliceArray() {
	a := [5]int{1, 2, 3, 4, 5}
	b := a[1:4:4]
	fmt.Printf("cap(b): %d", cap(b)) // cap(b): 3
}

type Info struct {
	Version string    // version string
	Time    time.Time // commit time
}
