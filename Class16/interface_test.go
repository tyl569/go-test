package Class16

import (
	"fmt"
	"testing"
)

func TestInterface(t *testing.T) {
	var p Programmer
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}

func main() {
	var p Programmer
	p = new(GoProgrammer)
	fmt.Println(p.WriteHelloWorld())
}

// 定义接口和方法
type Programmer interface {
	WriteHelloWorld() string
}

type GoProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() string {
	return "fmt.Println(\" Hello World\")"
}

func TestClient(t *testing.T) {
	var p Programmer;
	p = new(GoProgrammer)
	t.Log(p.WriteHelloWorld())
}
