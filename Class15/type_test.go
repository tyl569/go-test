package Class15

import (
	"fmt"
	"testing"
	"unsafe"
)

type Person struct {
	Name string
	Sex  string
	Age  int
}

// 默认构造函数
func (p Person) String() string {
	fmt.Printf("Address is %x \n", unsafe.Pointer(&p))
	return fmt.Sprintf("Name: %s - Sex: %s - Age: %d", p.Name, p.Sex, p.Age)
}

//默认构造函数
func (p *Person) String() string {
	fmt.Printf("Address is %x", unsafe.Pointer(p))
	return fmt.Sprintf("Name: %s - Sex: %s - Age: %d", p.Name, p.Sex, p.Age)
}

func TestPerson(t *testing.T) {
	person := Person{Name: "张三", Sex: "男", Age: 21}
	fmt.Printf("Address is %x \n", unsafe.Pointer(&person))
	t.Log(person.String())
	//t.Logf("用户名: Name=%s", person.Name)
	//person1 := Person{"李四", "男", 20}
	//t.Log(person1)
	//t.Logf("用户名: Name=%s", person1.Name)
	//person2 := new(Person) // 这里返回指针，相当于person2 := &Person{}
	//person2.Name = "王五"
	//person2.Sex = "女"
	//person2.Age = 30
	//t.Log(person2)
	//t.Logf("person is %T", person)
	//t.Logf("person2 is %T", person2)
}
