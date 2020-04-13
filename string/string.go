package _string

import "fmt"

func StringDoubleQuotationMarks() {
	s := "Hi, \nthis is \"RainbowMango\"."
	fmt.Println(s)
}

func StringBackslash() {
	s := `Hi, 
this is "RainbowMango".`
	fmt.Println(s)
}

func StringConnectSample1() string {
	var s string
	s = s + "a" + "b"
	return s
}

func StringConnectSample2() string {
	var s string
	s = s + "a"
	s = s + "b"
	return s
}
