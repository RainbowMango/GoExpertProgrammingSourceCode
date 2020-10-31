package gotest_test

import (
	"testing"
	"time"

	"github.com/rainbowmango/goexpertprogrammingsourcecode/GoExpert/src/gotest"
)

func BenchmarkMakeSliceWithoutAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithoutAlloc()
	}
}

func BenchmarkMakeSliceWithPreAlloc(b *testing.B) {
	for i := 0; i < b.N; i++ {
		gotest.MakeSliceWithPreAlloc()
	}
}

func BenchmarkSetBytes(b *testing.B) {
	b.SetBytes(1024 * 1024)
	for i := 0; i < b.N; i++ {
		time.Sleep(1 * time.Second) // 模拟待测函数
	}
}
