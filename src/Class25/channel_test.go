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
	retCh := make(chan string) // 创建channel
	go func() {                // 创建一个协程
		ret := service() // 返回Done
		fmt.Println("return result.")
		retCh <- ret // 向channel里面写数据
		fmt.Println("service exited.")
	}()
	return retCh
}

/**
输出顺序:
return result.
service exited.
Done

原因：
在执行AsyncService的方法的时候
先创建一个Chanel
创建一个协程
return retCh
输出channel内容

这个时候，在另外的协程，正在执行
	1、调用service方法
	2、打印return result.
	3、向channel里面写数据
	4、输出service exited.

	在这个协程中，需要等待channel的消费，所以最终执行结果4在3的前面输出
*/
func TestAsyncService(t *testing.T) {
	retCh := AsyncService()
	fmt.Println(<-retCh)
}

/**
输出顺序:
return result.
Done
service exited.
*/
func TestAsyncService2(t *testing.T) {
	retCh := AsyncService()
	time.Sleep(time.Second * 1)
	fmt.Println(<-retCh)
}

func TestChannel1(t *testing.T) {
	ch1 := make(chan string)
	go func() {
		ch1 <- "测试"
	}()
	fmt.Println(<-ch1)
}

func TestChannel2(t *testing.T) {
	ch1 := make(chan string)
	go func() {
		fmt.Println(<-ch1)
	}()
	ch1 <- "测试"
}
