package class26

import (
	"fmt"
	"testing"
	"time"
)

func service() string {
	time.Sleep(time.Millisecond * 500)
	return "Done"
}

func otherTask() {
	fmt.Println("working on something else")
	time.Sleep(time.Millisecond * 1000)
	fmt.Println("Task is done.")
}

func AsyncService() chan string {
	retCh := make(chan string)
	go func() {
		fmt.Println("returned result.")
		retCh <- service()
		fmt.Println("service exited.")
	}()
	return retCh
}
func TestAsync(t *testing.T) {
	select {
	case ret := <-AsyncService():
		t.Log(ret)
	case <-time.After(time.Microsecond * 100):
		t.Error("time out")
	}
}

func TestSelect(t *testing.T) {
	var ch1, ch2 chan int
	ch1 = CreateCh()
	ch2 = CreateCh()
	select {
	case <-ch1:
		t.Log("ch1 returns")
	case <-ch2:
		t.Log("ch2 returns")
	case <-time.After(time.Second * 1):
		t.Log("timeout")
		//default:
		//	t.Log("nothings returns")
	}
}

func CreateCh() chan int {
	ch1 := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			time.Sleep(time.Second * 2)
			ch1 <- i
		}
	}()

	return ch1
}
