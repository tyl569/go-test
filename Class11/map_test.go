package Class11

import "testing"

func TestMapWithFuncValue(t *testing.T) {
	m := map[int]func(op int) int{}
	m[1] = func(op int) int { return op }
	m[2] = func(op int) int { return op * op }
	m[3] = func(op int) int {
		op++
		return op
	}
	m[4] = func(op int) int { return op * op * op }
	t.Logf("输出第一个元素的值=%d", m[1](2))
	t.Logf("输出第二个元素的值=%d", m[2](2))
	t.Logf("输出第三个元素的值=%d", m[3](2))
}

func TestMapForSet(t *testing.T) {
	m := map[int]bool{}
	m[1] = true
	n := 1
	if m[n] {
		t.Log("检查元素n存在")
	} else {
		t.Log("检查元素n不存在")
	}
	delete(m, 1)
	if m[n] {
		t.Log("检查元素n存在")
	} else {
		t.Log("检查元素n不存在")
	}
}
