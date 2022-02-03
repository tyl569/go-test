package Class38

import (
	"fmt"
	"reflect"
	"testing"
)

type People struct {
	name string
	age  int
}

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func (p *People) Eat(food string) {
	fmt.Println(p.name, " eat ", food)
}

func (p *People) UpdateAge(age int) {
	p.age = age
}

func CheckType(v interface{}) {
	t := reflect.TypeOf(v)
	switch t.Kind() {
	case reflect.Float32, reflect.Float64:
		fmt.Println("Float")
		break
	case reflect.Int, reflect.Int32, reflect.Int64:
		fmt.Println("Int")
		break
	default:
		fmt.Println("Unknown", t)
	}
}

func TestCheck(t *testing.T) {
	var f float64
	f = 12
	CheckType(f)
	var i = 12
	CheckType(i)
	CheckType(&i)
}

func TestPeople(t *testing.T) {
	p := new(People)
	p.name = "张三"
	p.age = 20
	reflect.ValueOf(p).MethodByName("UpdateAge").Call([]reflect.Value{reflect.ValueOf(21)})
	reflect.ValueOf(p).MethodByName("Eat").Call([]reflect.Value{reflect.ValueOf("Apple")})
	t.Log(p.age)
}

func TestBasicInfo(t *testing.T) {
	p := new(BasicInfo)
	p.Name = "张三"
	p.Age = 20
	if tag, ok := reflect.TypeOf(*p).FieldByName("Name"); ok {
		fmt.Println("The tag is", tag.Tag.Get("json"))
	} else {
		fmt.Println("cannot get filed")
	}

}

func TestGetType(t *testing.T) {
	var i = "name"
	t.Log("输出对应的类型", reflect.TypeOf(i))
	t.Log("输出对应的值", reflect.ValueOf(i))
	t.Log("输出对应的值的类型", reflect.ValueOf(i).Type())
}
