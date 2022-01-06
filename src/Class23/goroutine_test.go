package Class23

import (
	"fmt"
	"testing"
)

func TestGoroutine(t *testing.T) {
	for i := 0; i < 10; i++ {
		go out(i)
	}
}

func out(i int) {
	fmt.Println(i)
}
