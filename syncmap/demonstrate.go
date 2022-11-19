package syncmap

import (
	"fmt"
	"sync"
)

func Demonstrate() {
	// sync.Map 的零值为空的map，声明后就可以直接使用，不需要再初始化
	var m sync.Map

	m.Store("Jim", 80)   // 写入Jim的成绩
	m.Store("Kevin", 90) // 写入Kevin的成绩
	m.Store("Jim", 90)   // 将修改Jim的成绩为90

	score, _ := m.Load("Jim") // 查询Jim的成绩
	fmt.Printf("Jim's score: %d\n", score.(int))

	m.Delete("Jim") // 删除Jim
}

func DemonstrateTypeUnsafe() {
	var m sync.Map
	m.Store("Kevin", 90) // 键类型为 string
	m.Store(80, "Kevin") // 键类型为 int

	fmt.Println("school report:")
	m.Range(func(key, value any) bool {
		fmt.Printf("Name: %s, score: %d\n", key.(string), value.(int)) // panic
		return true
	})
}

// DemonstrateCannotTransfer 用于演示sync.Map不能传递。
// 注：个人对编译器无法拦截此风险代码感到诧异，IDE反而可以给出告警
func DemonstrateCannotTransfer() {
	var m sync.Map
	m.Store("Jim", 80)   // 写入Jim的成绩
	m.Store("Kevin", 90) // 写入Kevin的成绩

	// 'func' passes a lock by the value: type 'sync.Map' contains 'sync.Mutex' which is 'sync.Locker'
	funcFoo := func(m sync.Map) {
		m.Range(func(key, value any) bool {
			fmt.Println("Name:", key, "Score:", value)
			return true
		})
	}

	funcFoo(m)
}

// DemonstrateInsertUncomparableMap 演示不能存储不可比较类型，即便内部map类型为map[any]*entry
func DemonstrateInsertUncomparableMap() {
	nativeMap := make(map[string]int)
	var m sync.Map
	m.Store(nativeMap, "should panic") // panic: runtime error: hash of unhashable type map[string]int
	m.Load(nativeMap)
}

func DemonstrateInsertUncomparableFunc() {
	funcFoo := func() {
		fmt.Println("I'm a func")
	}
	var m sync.Map
	m.Store(funcFoo, "should panic") // panic: runtime error: hash of unhashable type func()
}
