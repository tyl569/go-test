package Class34

import (
	"runtime"
	"sync"
	"testing"
)

func TestSync(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Created success")
			return 100
		},
	}

	v := pool.Get().(int) // 首次get的时候，会先看缓存中是否有值，如果没有值，那么就会进行创建
	t.Log("print v=", v)  // 打印
	pool.Put(3)
	runtime.GC()         // GC 会清除掉sync.pool里面的缓存的对象
	v = pool.Get().(int) // 获取新的值
	t.Log("print new v=", v)
}

func TestSyncPoolInMultiGoroutine(t *testing.T) {
	pool := &sync.Pool{
		New: func() interface{} {
			t.Log("Created success")
			return 10
		},
	}
	pool.Put(100)
	pool.Put(100)
	pool.Put(100)

	//runtime.GC()

	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(id int) {
			t.Log(pool.Get())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
