package main

import (
	"fmt"
	"unsafe"
)

type Person struct {
	Name string
	Sex  string
	Age  int
}

// 默认构造函数
func (p Person) String1() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&p))// 打印内存地址
	return fmt.Sprintf("Name: %s - Sex: %s - Age: %d", p.Name, p.Sex, p.Age)
}

//默认构造函数
func (p *Person) String2() string {
	fmt.Printf("Address is %x", unsafe.Pointer(p)) // 打印内存地址
	return fmt.Sprintf("Name: %s - Sex: %s - Age: %d", p.Name, p.Sex, p.Age)
}

func main() {
	person := Person{Name: "张三", Sex: "男", Age: 21}
	fmt.Printf("Address is %x \n", unsafe.Pointer(&person))
	person.String1()

	person1 := Person{Name: "张三", Sex: "男", Age: 21}
	fmt.Printf("Address is %x \n", unsafe.Pointer(&person1))
	person1.String2()
}