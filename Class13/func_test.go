package Class13

import (
	"fmt"
	"math/rand"
	"time"
)

//// 函数是一等公民
//func TestFunc(t *testing.T) {
//	a, _ := returnMultiValues()
//	t.Log(a)
//	funcRunTime := timeSpend(slow)
//	t.Log(funcRunTime(10))
//}
func main() {
	funcRunTime := timeSpend(slow)
	fmt.Println(funcRunTime(10))
}

/**
 传入的参数是一个函数: inner func(op int) int
	这个函数的入参是一个整形，返回的也是一个整形
 返回的参数也是一个函数：func(op int) int
	入参是一个整形，返回的也是一个整形
*/
func timeSpend(inner func(op int) int) func(op int) int {
	return func(n int) int {
		start := time.Now()
		ret := inner(n)
		fmt.Println("time spent:", time.Since(start).Seconds())
		return ret
	}
}

func slow(op int) int {
	time.Sleep(time.Second + 1)
	return op
}

func returnMultiValues() (int, int) {
	return rand.Intn(10), rand.Intn(20)
}


