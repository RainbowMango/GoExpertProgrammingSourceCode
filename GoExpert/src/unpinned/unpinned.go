package unpined

import (
	"fmt"
)

func Process1(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func() {
			fmt.Printf("Worker start process task: %s\n", task)
		}()
	}
}

func Process2(tasks []string) {
	for _, task := range tasks {
		// 启动协程并发处理任务
		go func(t string) {
			fmt.Printf("Worker start process task: %s\n", t)
		}(task)
	}
}

func Double(a int) int {
	return a * 2
}
