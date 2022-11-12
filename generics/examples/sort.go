package examples

import "sort"

// Ordered 定义适用于所有可排序类型的类型集合
// 可排序类型是指支持 <、<=、>=、>操作的类型
type Ordered interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64 |
		~uint | ~uint8 | ~uint16 | ~uint32 | ~uint64 | ~uintptr |
		~float32 | ~float64 |
		~string
}

// orderedSlice 实现了sort.Interface，所以可以使用sort.Sort排序
type orderedSlice[T Ordered] []T

func (s orderedSlice[T]) Len() int           { return len(s) }
func (s orderedSlice[T]) Less(i, j int) bool { return s[i] < s[j] }
func (s orderedSlice[T]) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

// OrderedSlice 按升序排列slice
func OrderedSlice[T Ordered](s []T) {
	// s先被转换为orderedSlice，然后再排序
	// 因此s不必再实现sort.Interface接口
	sort.Sort(orderedSlice[T](s))
}
