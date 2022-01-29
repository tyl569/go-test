package Class31

import (
	"fmt"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

// 假如有一个返回对应的结果，就返回
// 场景类似并发调用百度、Google，查询数据，有一个返回结果就进行输出
func firstResponse() string {
	numberOfRes := 10
	ch := make(chan string, numberOfRes)
	for i := 0; i < numberOfRes; i++ {
		go func(n int) {
			ret := runTask(n)
			ch <- ret // 写入协程数据
		}(i)
	}
	return <-ch
}

func TestForResponse(t *testing.T) {
	t.Log(firstResponse())
}
