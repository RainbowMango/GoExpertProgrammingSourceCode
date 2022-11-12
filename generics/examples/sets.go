package examples

import "sync"

// Set 底层使用map实现
type Set[T comparable] map[T]struct{}

// MakeSet 构造某个（comparable）类型的 set
func MakeSet[T comparable]() Set[T] {
	return make(Set[T])
}

// Add 添加元素v
func (s Set[T]) Add(v T) {
	s[v] = struct{}{}
}

// Delete 删除元素v
// 如果v不存在等同于空操作
func (s Set[T]) Delete(v T) {
	delete(s, v)
}

// Contains 查询元素v是否已存在
func (s Set[T]) Contains(v T) bool {
	_, ok := s[v]
	return ok
}

// Len 返回元素个数
func (s Set[T]) Len() int {
	return len(s)
}

// Iterate 遍历set,并逐个调用函数 f
func (s Set[T]) Iterate(f func(T)) {
	for v := range s {
		f(v)
	}
}

func SetExample() {
	s := MakeSet[int]() // 构建一个存储int的set
	s.Add(1)
	if s.Contains(2) {
		// do something
	}
}

// ThreadSafeSet 行为同Set一致，但它可以是线程安全的。
type ThreadSafeSet[T comparable] struct {
	l sync.RWMutex
	m map[T]struct{}
}
