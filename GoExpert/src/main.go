package main

import (
    "string"
    "timer"
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
}

