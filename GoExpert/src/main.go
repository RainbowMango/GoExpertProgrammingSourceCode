package main

import (
    "string"
    "timer"
    "sugar"
)

func GetFromMap() string {
    myMap := make(map[string]string, 10)

    name := myMap["Horen"]

    return name
}

func main() {
    GetFromMap()
    stringExample.RunStringPackage()
    timerExample.RunTimerPackage()
    sugar.SugerPackage()
}

