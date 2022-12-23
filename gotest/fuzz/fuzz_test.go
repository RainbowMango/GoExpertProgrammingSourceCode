package fuzz

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func ExampleReverse() {
	input := "The quick brown fox jumped over the lazy dog"
	rev, _ := Reverse(input)
	doubleRev, _ := Reverse(rev)
	fmt.Printf("%s\n", doubleRev)
	// Output:
	// The quick brown fox jumped over the lazy dog
}

func TestReverse(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{in: "Hello, world", want: "dlrow ,olleH"},
		{in: " ", want: " "},
		{in: "!12345", want: "54321!"},
	}
	for _, tc := range testcases {
		rev, err := Reverse(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}

func FuzzReverse(f *testing.F) {
	testcases := []string{"Hello, world", " ", "!12345"}
	for _, tc := range testcases {
		f.Add(tc) // 输入测试种子
	}
	f.Fuzz(func(t *testing.T, orig string) {
		if !utf8.ValidString(orig) { // 忽略构造的无效字符（非UTF-8编码字符串）
			return
		}
		rev, err := Reverse(orig)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if !utf8.ValidString(rev) {
			t.Fatalf("Reverse produced invalid UTF-8 string %q", rev)
		}

		doubleRev, err := Reverse(rev)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}

		if orig != doubleRev {
			t.Errorf("Before: %q, after: %q", orig, doubleRev)
		}
	})
}

func TestReverseV2(t *testing.T) {
	testcases := []struct {
		in, want string
	}{
		{in: "Hello, world", want: "dlrow ,olleH"},
		{in: " ", want: " "},
		{in: "!12345", want: "54321!"},
		{in: "中国", want: "国中"},
	}
	for _, tc := range testcases {
		rev, err := ReverseV2(tc.in)
		if err != nil {
			t.Errorf("Unexpected error: %v", err)
		}
		if rev != tc.want {
			t.Errorf("Reverse: %q, want %q", rev, tc.want)
		}
	}
}
