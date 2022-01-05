package class18

import (
	"fmt"
	"testing"
)

func doSomething(p interface{}) {
	// if i, ok := p.(int); ok {
	// 	fmt.Println("int", i)
	// 	return
	// }
	// if s, ok := p.(string); ok {
	// 	fmt.Println("string", s)
	// 	return
	// }
	// fmt.Println("unkonw type")
	switch v := p.(type) {
	case int:
		fmt.Println("int", v)
	case string:
		fmt.Println("string", v)
	default:
		fmt.Println("unkonw type")
	}
}
func TestEmpty(t *testing.T) {
	doSomething(1)
	doSomething("tengyunlong")
}
