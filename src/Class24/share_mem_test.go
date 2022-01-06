package Class24

import (
	"sync"
	"testing"
	"time"
)

func TestCounter(t *testing.T) {
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			counter++
		}()
	}
	time.Sleep(time.Second * 1) // 增加事件等待，防止外部的协程比1000个协程先执行结束
	t.Logf("counter = %d", counter)
}

func TestCounterForSafe(t *testing.T) {
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		go func() {
			defer func() {
				mut.Unlock() // 解锁
			}()
			mut.Lock() // 加锁
			counter++
		}()
	}
	time.Sleep(time.Second * 1) // 增加事件等待，防止外部的协程比1000个协程先执行结束
	t.Logf("counter = %d", counter)
}

func TestCounterForWaitGroup(t *testing.T) {
	var wg sync.WaitGroup
	var mut sync.Mutex
	counter := 0
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				mut.Unlock() // 释放锁资源
			}()
			mut.Lock() // 加锁
			counter++
			wg.Done()
		}()
	}
	wg.Wait() //等到所有的任务执行结束
	t.Logf("counter = %d", counter)
}
