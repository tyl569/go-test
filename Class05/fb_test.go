package Class05

import (
	"testing"
)

func TestFibList(t *testing.T) {
	var a int = 1
	var b int = 1
	t.Log(a)
	for i := 0; i < 5; i++ {
		tmp := a
		a = b
		b = tmp + a
	}
	t.Log(b)
}
