package main

import (
	deferexample "defer"
	"flag"
	"string"
	"sugar"
	"timer"
)

func GetFromMap() string {
	myMap := make(map[string]string, 10)

	name := myMap["Horen"]
	flag.Args()
	return name
}

func main() {
	deferexample.Dived(0)
	GetFromMap()
	stringExample.RunStringPackage()
	timerExample.RunTimerPackage()
	sugar.SugerPackage()
}
