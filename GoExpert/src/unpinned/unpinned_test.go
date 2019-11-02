package unpined

import (
	"fmt"
	"strconv"
	"testing"
)

// TestProcess1 由于引用了循环变量，导致协程打印混乱
// 打印结果如下（每次执行均不一致）：
// E:\OpenSource\RainbowMango\GoExpertProgrammingSourceCode\GoExpert\src>go test ./unpined -v
// === RUN   TestProcess1
// Worker start process task: 3
// Worker start process task: 5
// Worker start process task: 5
// Worker start process task: 5
// Worker start process task: 10
// Worker start process task: 7
// Worker start process task: 8
// Worker start process task: 9
// Worker start process task: 4
// Worker start process task: 11
// Worker start process task: 11
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// Worker start process task: 19
// --- PASS: TestProcess1 (0.00s)
// PASS
// ok      unpined 0.406s
func TestProcess1(t *testing.T) {
	var tasks []string

	for i := 0; i < 20; i++ {
		tasks = append(tasks, strconv.Itoa(i))
	}
	Process1(tasks)
}

func TestProcess2(t *testing.T) {
	var tasks []string

	for i := 0; i < 20; i++ {
		tasks = append(tasks, strconv.Itoa(i))
	}
	Process2(tasks)
}

func TestDouble(t *testing.T) {
	var tests = []struct {
		name         string
		input        int
		expectOutput int
	}{
		{
			name:         "double 1 should got 2",
			input:        1,
			expectOutput: 2,
		},
		{
			name:         "double 2 should got 4",
			input:        2,
			expectOutput: 4,
		},
	}

	for _, test := range tests {
		tc := test // 显式绑定，每次循环都会生成一个新的tc变量
		fmt.Println("tc= ", &tc)
		t.Run(tc.name, func(t *testing.T) {
			if tc.expectOutput != Double(tc.input) {
				t.Fatalf("expect: %d, but got: %d", tc.input, tc.expectOutput)
			}
		})
	}
}
