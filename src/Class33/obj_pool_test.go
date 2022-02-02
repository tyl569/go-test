package Class33

import (
	"errors"
	"fmt"
	"testing"
	"time"
)

type ResourcePool struct {
}

type ObjPool struct {
	bufChan chan *ResourcePool
}

func NewObjPool(numObj int) *ObjPool {
	objPool := ObjPool{}
	objPool.bufChan = make(chan *ResourcePool, numObj)
	for i := 0; i < numObj; i++ {
		objPool.bufChan <- &ResourcePool{}
	}
	return &objPool
}

func (p *ObjPool) GetObj(timeout time.Duration) (*ResourcePool, error) {
	select {
	case ret := <-p.bufChan:
		return ret, nil
	case <-time.After(timeout):
		return nil, errors.New("timeout")
	}
}

func (p *ObjPool) ReleaseObj(obj *ResourcePool) error {
	select {
	case p.bufChan <- obj:
		return nil
	default:
		return errors.New("overflow")
	}
}

func TestClass(t *testing.T) {
	test1 := new(ResourcePool)
	test2 := ResourcePool{}
	fmt.Println(test1)
	fmt.Println(&test2)
}

func TestObjPool(t *testing.T) {
	pool := NewObjPool(10)
	for i := 0; i < 11; i++ {
		if v, err := pool.GetObj(time.Second * 1); err != nil {
			t.Error(err)
		} else {
			t.Logf("get obj: %p", v)
			if err := pool.ReleaseObj(v); err != nil {
				t.Error(err)
			}
		}

	}
	fmt.Println("Done")
}
