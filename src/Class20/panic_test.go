package Class20

import (
	"fmt"
	"os"
	"testing"
)

func TestPanic(t *testing.T) {
	defer func() {
		fmt.Println("finally")
	}()
	fmt.Println("start")
	os.Exit(1)
	//panic(errors.New("Something wrong"))
}
