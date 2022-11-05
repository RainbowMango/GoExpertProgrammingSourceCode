package overall

// Vector 是可容纳任意类型的容器
type Vector[T any] []T

// Push 加入元素到容器
// 为泛型类型定义方法时，receiver类型必须带上类型参数
func (v *Vector[T]) Push(x T) {
	*v = append(*v, x)
}
