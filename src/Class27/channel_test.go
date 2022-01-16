package Class27

import (
	"fmt"
	"sync"
	"testing"
)

func dataProduce(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		wg.Done()
	}()

}

func dataConsume(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			data := <-ch
			fmt.Printf("输出data=%d \n", data)
		}
		wg.Done()
	}()
}

func TestChannel(t *testing.T) {
	var wg sync.WaitGroup
	ch1 := make(chan int)
	wg.Add(1)
	dataProduce2(ch1, &wg)
	wg.Add(1)
	dataConsume2(ch1, &wg)
	wg.Wait()
}

func dataProduce2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for i := 0; i < 10; i++ {
			ch <- i
		}
		close(ch)
		wg.Done()
	}()
}

func dataConsume2(ch chan int, wg *sync.WaitGroup) {
	go func() {
		for true {
			if data, ok := <-ch; ok {
				fmt.Printf("输出data=%d \n", data)
			} else {
				fmt.Println("produce channel has closed.")
				break
			}
		}
		wg.Done()
	}()
}
