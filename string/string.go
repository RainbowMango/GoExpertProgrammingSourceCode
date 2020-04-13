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
