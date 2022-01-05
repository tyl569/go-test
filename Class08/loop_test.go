package Class08

import "testing"

func TestLoop(t *testing.T) {
	a := 1
	for a < 5 {
		t.Log("a=", a)
		a++
	}
}
