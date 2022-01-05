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
