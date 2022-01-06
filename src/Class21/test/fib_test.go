package test

import (
	"Class21"
	"fmt"
	cm "github.com/easierway/concurrent_map"
	"testing"
)

func TestFibList(t *testing.T) {
	var fibList []int
	fibList, _ = Class21.GetFibList(5)

	fmt.Println(fibList)
}

func TestConcurrentMap(t *testing.T) {
	m := cm.CreateConcurrentMap(99)
	m.Set(cm.StrKey("Key"), 10)
	t.Log(m.Get(cm.StrKey("Key")))
}
