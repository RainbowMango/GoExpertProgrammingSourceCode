package _struct

/*
关于下面类型的描述，正确的是：
单选：
- A：age为私有字段，只能通过类型方法访问
- B：Name和age字段在可见性上没有区别
- C：age为私有字段，变量初始化时不能直接赋值
- D：同一个package中，Name和age字段在可见性上没有区别
*/
type People struct {
	Name string
	age  int
}

func (p *People) SetName(name string) {
	p.Name = name
}

func (p *People) SetAge(age int) {
	p.age = age
}

func (p *People) GetAge() int {
	return p.age
}
