package Class06

import "testing"

func TestType(t *testing.T) {
	var a int32 = 1
	var b int64
	//b = a // 报错，不能隐式转换
	b = int64(a)
	t.Log(b)
}

func TestPoint(t *testing.T) {
	a := 1
	aPtr := &a
	//aPtr = aPtr + 1 // 报错，不支持指针运算
	t.Log(a, aPtr)
	t.Logf("%T %T", a, aPtr)
}

func TestTring(t *testing.T) {
	var s string
	t.Log("==" + s + "==")
	var sPtr = &s
	t.Log(sPtr)

}
