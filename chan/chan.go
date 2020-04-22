package _chan

import "fmt"

func ChanInitDeclare() {
	var ch chan int // 声明管道
	fmt.Println(ch == nil)
}

func ChanInitMake() {
	ch1 := make(chan string)    // 无缓冲管道
	ch2 := make(chan string, 5) // 带缓冲管道

	fmt.Println(ch1 == nil)
	fmt.Println(ch2 == nil)
}

func ChanOperator() {
	ch := make(chan int, 10)
	ch <- 1   // 数据流入管道
	d := <-ch // 数据流出管道
	fmt.Println(d)
}

func ChanParamRW(ch chan int) {
	// 管道可读写
}

func ChanParamR(ch <-chan int) {
	// 只能从管道读取数据
}

func ChanParamW(ch chan<- int) {
	// 只能向管道写入数据
}
