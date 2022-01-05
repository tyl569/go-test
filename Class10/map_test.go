package Class10

import "testing"

func TestMap(t *testing.T) {
	m1 := map[string]int{"one": 1, "two": 2, "three": 3}
	t.Log(m1)
	t.Log(m1["one"])
	t.Log(m1["four"])
	t.Logf("m1的长度=%d", len(m1))
	m2 := map[int]int{1: 1, 2: 2}
	t.Log(m2)
	t.Log(m2[3])
	m3 := make(map[int]int, 10)
	t.Log(m3)
	t.Logf("m3的长度=%d", len(m3))
	m4 := map[int]string{1: "test"}
	t.Log(m4[2])
	if _, ok := m4[3]; !ok {
		t.Log("数据不存在")
	}
	m5 := map[string]string{}
	m5["name"] = "张三"
	t.Log(m5)
}

func TestLoopMap(t *testing.T) {
	m1 := map[string]string{"name": "张三", "sex": "男"}
	for k, v := range m1 {
		t.Logf("输出的k=%s", k)
		t.Logf("输出的v=%s", v)
	}
}
