package timerExample

import (
    "time"
    "log"
)

// WaitChannel 用于演示timer的典型使用场景
func WaitChannel(conn <-chan string) bool {
    timer := time.NewTimer(1 * time.Second)

    select {
    case <- conn:
        timer.Stop()
        return true
    case <- timer.C: // 超时
        println("WaitChannel timeout!")
        return false
    }
}

// DelayFunction 用于演示timer的使用场景: 延迟调用
func DelayFunction() {
    timer := time.NewTimer(5 * time.Second)

    select {
    case <- timer.C:
        log.Println("Delayed 5s, start to do something.")
    }
}

// AfterDemo 用于演示time.After()用法
func AfterDemo() {
    log.Println(time.Now())
    <- time.After(1 * time.Second)
    log.Println(time.Now())
}

// AfterFuncDemo 用于演示time.AfterFunc用法
func AfterFuncDemo() {
    log.Println("AfterFuncDemo start: ", time.Now())
    time.AfterFunc(1 * time.Second, func() {
        log.Println("AfterFuncDemo end: ", time.Now())
    })

    time.Sleep(2 * time.Second) // 等待协程退出
}

func TickerDemo() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        log.Println("Ticker tick.")
    }
}

func RunTimerPackage() {
    channel := make(chan string)

    WaitChannel(channel)
    DelayFunction()
    AfterDemo()
    AfterFuncDemo()

    TickerDemo()
}