package Class20

import (
	"errors"
	"fmt"
	"testing"
)

func TestRecover(t *testing.T) {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("Catch Error", err)
		}
	}()
	fmt.Println("start")
	panic(errors.New("this error"))
}
