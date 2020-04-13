package _string

import "testing"

/*
TODO: 分析测试失败原因
	// Output:
	// Hi,
	// this is "RainbowMango".
*/
func ExampleStringDoubleQuotationMarks() {
	StringDoubleQuotationMarks()
}

/*
TODO: 分析测试失败原因
	// Output:
	// Hi,
	// This is "RainbowMango".
*/
func ExampleStringBackslash() {
	StringBackslash()
}

func BenchmarkStringConnectSample1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringConnectSample1()
	}
}

func BenchmarkStringConnectSample2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		StringConnectSample2()
	}
}
