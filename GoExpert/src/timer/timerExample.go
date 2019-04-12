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

// TickerDemo 用于演示ticker基础用法
func TickerDemo() {
    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for range ticker.C {
        log.Println("Ticker tick.")
    }
}

func GetNewPassenger() string{
    return "task"
}

func Launch([]string) {
    return
}

// TickerLaunch用于演示ticker聚合任务用法
func TickerLaunch() {
    ticker := time.NewTicker(5 * time.Minute)
    maxPassenger := 30                   // 每车最大装载人数
    passengers := make([]string, 0, maxPassenger)

    for {
        passenger := GetNewPassenger() // 获取一个新乘客
        if passenger != "" {
            passengers = append(passengers, passenger)
        } else {
            time.Sleep(1 * time.Second)
        }

        select {
        case <- ticker.C:               // 时间到，发车
            Launch(passengers)
            passengers = []string{}
        default:
            if len(passengers) >= maxPassenger {  // 时间没到，车已座满，发车
                Launch(passengers)
                passengers = []string{}
            }
        }
    }
}

func WrongTicker() {
    for {
        select {
        case <-time.Tick(1 * time.Second):
            log.Printf("Resource leak!")
        }
    }
}

func RunTimerPackage() {
    channel := make(chan string)

    WaitChannel(channel)
    DelayFunction()
    AfterDemo()
    AfterFuncDemo()

    //TickerDemo()

    //WrongTicker()
}