package Class09

import "testing"

func TestArray(t *testing.T) {
	var arr [3]int
	t.Log(arr[0])
}

func TestArrayLoop(t *testing.T) {
	var a = [...]int{1, 2, 3, 5, 1}

	t.Log(a[1:])

	for i := 0; i < len(a); i++ {
		t.Log("输出", a[i])
	}
	for _, e := range a {
		t.Log("输出", e)
	}

	var name = [...]string{"张三", "李四"}
	for _, e := range name {
		t.Log("输出", e)
	}

	var b = [3][2]int{{1, 2}, {3, 4}, {5, 6}}
	for _, val := range b {
		for _, val1 := range val {
			t.Log("输出二维数组的元素", val1)
		}
	}
}

func TestSLice(t *testing.T) {
	var s0 []int
	t.Log(len(s0), cap(s0))
	s0 = append(s0, 1)
	t.Log(len(s0), cap(s0))
	s1 := []int{1, 2, 3, 4}
	t.Log(len(s1), cap(s1))

	s2 := make([]int, 3, 5) // 初始化一个长度为3，容量为5的数组
	t.Log(len(s2), cap(s2))
	t.Log(s2[0], s2[1], s2[2])
	s2 = append(s2, 1)
	t.Log(len(s2), cap(s2))
}

func TestSliceGrowing(t *testing.T) {
	s := []int{}
	for i := 0; i < 10; i++ {
		s = append(s, i)
		t.Log(len(s), cap(s)) // 当数组的容量不够的时候，数组会申请一个2倍的存储空间，把数组的值拷贝到新的存储空间
	}
}

func TestSliceShareMem(t *testing.T) {
	year := []string{"Jan", "Feb", "Mar", "Apr", "May", "Jun", "Jul", "Aug", "Sep", "Oct", "Nov", "Dec"}
	Q2 := year[3:6]
	t.Log(Q2, len(Q2), cap(Q2))
	summer := year[5:8]
	t.Log(summer, len(summer), cap(summer))
}
