package Class28

import (
	"fmt"
	"testing"
	"time"
)

func TestCancel(t *testing.T) {
	cancelChan := make(chan struct{}, 0)
	for i := 0; i < 5; i++ {
		go func(i int, cancelCh chan struct{}) {
			for {
				if isCancelled(cancelChan) {
					break
				}
				time.Sleep(time.Millisecond * 5)
			}
			fmt.Println(i, "Canceled")
		}(i, cancelChan)
	}
	cancel2(cancelChan)
	time.Sleep(time.Second * 1)
}

func cancel1(cancelChan chan struct{}) {
	cancelChan <- struct{}{}
}

func cancel2(ch chan struct{}) {
	close(ch)
}

func isCancelled(cancelChan chan struct{}) bool {
	select {
	case <-cancelChan:
		return true
	default:
		return false
	}
}
