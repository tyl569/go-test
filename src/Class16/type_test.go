package Class16

import (
	"fmt"
	"testing"
	"time"
)

type IntConv func(op int) int // 自定义一个IntConv类型，这个类型是一个方法

func timeSpend(inner IntConv) IntConv {
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

func TestFunc(t *testing.T) {
	funcRunTime := timeSpend(slow)
	t.Log(funcRunTime(10))
}

func main() {
	funcRunTime := timeSpend(slow)
	fmt.Println(funcRunTime(10))
}
