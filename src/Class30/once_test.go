package Class30

import (
	"fmt"
	"sync"
	"testing"
	"unsafe"
)

type Singleton struct {
}

var singleInstance *Singleton
var once sync.Once

func GetSingleObj() *Singleton {
	once.Do(func() {
		fmt.Println("Create Obj")
		singleInstance = new(Singleton)
	})
	return singleInstance
}

func TestGetSingle(t *testing.T) {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			obj := GetSingleObj()
			fmt.Printf("address: %x\n", unsafe.Pointer(obj))
			wg.Done()
		}()
	}
	wg.Wait()
}

func TestCurren(t *testing.T) {
	num := 1
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			num++
			wg.Done()
		}(num)
	}
	wg.Wait()
	t.Log("输出数据", num)
}

func TestCurren1(t *testing.T) {
	num := 1
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			once.Do(func() {
				num++
			})
			wg.Done()
		}(num)
	}
	wg.Wait()
	t.Log("输出数据", num)
}
