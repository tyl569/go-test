package Class25

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 50)
	return "Done"
}

func OtherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 100)
	fmt.Println("Task is done.")
}

func TestService(t *testing.T) {
	fmt.Println(service())
	OtherTask()
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		ret := service()
		fmt.Println("returned result.")
		retCh <- ret
		fmt.Println("service exited.")
	}()
	return retCh
}

// 开启新的协程
// 主协程先运行，输出 working on something else
// 子协程调用service函数
// 子协程输出 returned result.
// 子协程写入Done
// 主协程输出 Task is done.
// 主协程阻塞，拿到Done，并输出
// 子协程继续执行，输出service exited.
func TestAsync(t *testing.T) {
	retCh := AsyncService()
	OtherTask()
	fmt.Println(<-retCh)
}
