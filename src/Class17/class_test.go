package Class17

import (
	"fmt"
	"testing"
)

type Pet struct {
}

func (p *Pet) Speak() {
	fmt.Println("...")
}

func (p *Pet) SpeakTo(name string) {
	p.Speak()
	fmt.Println(name)
}

type Dog struct {
	Pet
}

func (d *Dog) Speak() {
	fmt.Println("Wang!")
}

func TestClass(t *testing.T) {
	d := new(Dog)
	d.Speak()
	d.SpeakTo("肥龙") // 这里会输出 ... 肥龙，并不会出现 Wang!肥龙
}
