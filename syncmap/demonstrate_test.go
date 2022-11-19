package syncmap

func ExampleDemonstrate() {
	Demonstrate()
	// output:
	// Jim's score: 90
}

/*
该示例仅用于演示sync.Map类型不安全的特点，会引发panic
func ExampleDemonstrateTypeUnsafe() {
	DemonstrateTypeUnsafe()
	// output:
	// panic
}
*/

func ExampleDemonstrateCannotTransfer() {
	DemonstrateCannotTransfer()
	// output:
	// Name: Jim Score: 80
	// Name: Kevin Score: 90
}

/*
该示例仅用于演示sync.Map不能存储map的事实
func ExampleDemonstrateInsertUncomparable() {
	DemonstrateInsertUncomparableMap()
	// output:
	//
}
*/

/*
该示例用于演示sync.Map不能存储函数的事实
func ExampleDemonstrateInsertUncomparableFunc() {
	DemonstrateInsertUncomparableFunc()
	// output:
	//
}
*/
