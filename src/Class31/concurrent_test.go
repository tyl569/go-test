package Class31

import (
	"fmt"
	"testing"
	"time"
)

// 假设这里是调用搜索引擎
func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func firstResponse() string {
	numberOfRes := 10
	ch := make(chan string, numberOfRes)
	for i := 0; i < numberOfRes; i++ {
		go func(n int) {
			ret := runTask(n)
			ch <- ret
		}(i)
	}
	return <-ch
}

func TestForResponse(t *testing.T) {
	t.Log(firstResponse())
}
