package main

import (
    "testing"
)

func MakeSlice() []int {
    slice := make([]int, 10000000)
    for index := range slice {
        slice[index] = index
    }

    return slice
}

//下面函数通过遍历切片，打印切片的下标和元素值，请问性能上有没有可优化的空间？
func RangeSlice(slice []int) {
    for index, value := range slice {
        _, _ = index, value
    }
}
//程序解释：函数中使用for-range对切片进行遍历，获取切片的下标和元素素值，这里忽略函数的实际意义。
//参考答案：遍历过程中每次迭代会对index和value进行赋值，如果数据量大或者value类型为string时，对value的赋值操作可能是多余的，可以在for-range中忽略value值，使用slice[index]引用value值。

func BenchmarkSlice(b *testing.B) {
    slice := MakeSlice()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RangeSlice(slice)
    }
}

func MakeMap() map[int]string {
    myMap := make(map[int]string, 10000000)
    for i := 0; i < 10000000; i++ {
        myMap[i] = "Expert"
    }

    return myMap
}

//下面函数通过遍历Map，打印Map的key和value，请问性能上有没有可优化的空间？
func RangeMap(myMap map[int]string) {
    for key, _ := range myMap {
        _, _ = key, myMap[key]
    }
}
//程序解释：函数中使用for-range对map进行遍历，获取map的key值，并跟据key值获取获取value值，这里忽略函数的实际意义。
//参考答案：函数中for-range语句中只获取key值，然后跟据key值获取value值，虽然看似减少了一次赋值，但通过key值查找value值的性能消耗可能高于赋值消耗。能否优化取决于map所存储数据结构特征、结合实际情况进行。

func BenchmarkMap(b *testing.B) {
    myMap := MakeMap()

    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        RangeMap(myMap)
    }
}


//对于for-range语句的实现，可以从编译器源码中找到答案。

//编译器源码`gofrontend/go/statements.cc/For_range_statement::do_lower()`方法中有如下注释，对range支持的类型都会做一个C风格的循环。
// Arrange to do a loop appropriate for the type.  We will produce
//   for INIT ; COND ; POST {
//           ITER_INIT
//           INDEX = INDEX_TEMP
//           VALUE = VALUE_TEMP // If there is a value
//           original statements
//   }


//题目三:
//请问如下程序是否能正常结束？
//func main() {
//   v := []int{1, 2, 3}
//    for i:= range v {
//        v = append(v, i)
//    }
//}
//程序解释：main()函数中定义一个切片v，通过range遍历v，遍历过程中不断向v中添加新的元素。
//参考答案：能够正常结束。循环内改变切片的长度，不影响循环次数，循环次效在循环开始前就已经确定了。


//Slice:
// The loop we generate:
//   for_temp := range
//   len_temp := len(for_temp)
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = for_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }

//Array:
// The loop we generate:
//   len_temp := len(range)
//   range_temp := range
//   for index_temp = 0; index_temp < len_temp; index_temp++ {
//           value_temp = range_temp[index_temp]
//           index = index_temp
//           value = value_temp
//           original body
//   }

//Map:
// The loop we generate:
//   var hiter map_iteration_struct
//   for mapiterinit(type, range, &hiter); hiter.key != nil; mapiternext(&hiter) {
//           index_temp = *hiter.key
//           value_temp = *hiter.val
//           index = index_temp
//           value = value_temp
//           original body
//   }

//channel:
// The loop we generate:
//   for {
//           index_temp, ok_temp = <-range
//           if !ok_temp {
//                   break
//           }
//           index = index_temp
//           original body
//   }

//编程tips：
//- 遍历过程中可以适情况放弃接收index或value，可以一定程度上提升性能
//- 遍历channel时，如果channel中没有数据，可能会阻塞
//- 尽量避免遍历过程中修改原数据

//总结：
//for-range的实现实际上是C风格的for循环
//使用index,value接收range返回值会发生一次数据拷贝
