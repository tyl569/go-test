package Class32

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func runTask(id int) string {
	time.Sleep(time.Millisecond * 10)
	return fmt.Sprintf("The result is from %d", id)
}

func TestForResponse(t *testing.T) {
	var wg sync.WaitGroup
	list := make([]string, 10)
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(n int) {
			ret := runTask(n)
			list = append(list, ret)
			wg.Done()
		}(i)
	}
	wg.Wait()
	fmt.Println(list)
}
