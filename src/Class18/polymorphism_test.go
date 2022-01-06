package class18

import (
	"fmt"
	"testing"
)

type Code string

type Programmer interface {
	WriteHelloWorld() Code
}

type GoProgrammer struct {
}

type JavaProgrammer struct {
}

func (g *GoProgrammer) WriteHelloWorld() Code {
	return "fmt.Println(\"Hello World\")"
}

func (j *JavaProgrammer) WriteHelloWorld() Code {
	return "System.out.Println(\"Hello World\")"
}

func writeFirstProgram(p Programmer) {
	fmt.Printf("%T %v \n", p, p.WriteHelloWorld())
}

func TestProgram(t *testing.T) {
	gPro := new(GoProgrammer)
	jPro := new(JavaProgrammer)
	writeFirstProgram(gPro)
	writeFirstProgram(jPro)
}
