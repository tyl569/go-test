package Class07

import "testing"

func TestArray(t *testing.T) {
	a := [...]int{1, 2, 3, 4}
	//b := [...]int{1, 2, 3, 4, 6}
	c := [...]int{1, 2, 3, 5}
	d := [...]int{1, 2, 3, 4}
	//t.Log(a == b)// 报错，元素个数不一致
	t.Log(a == c)
	t.Log(a == d)
}
