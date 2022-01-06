package Class12

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

func TestString(t *testing.T) {
	//s := "qwerty";
	//t.Logf("输出第2个byte=%s", s[2])
	s1 := "\xE1\xa3\xF1"
	t.Logf("输出字符串=%s", s1)
}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	c := []rune(s)      // 返回utf8的字符长度
	fmt.Println(len(c)) // 7
	t.Log(c)
	for _, strVal := range s {
		t.Logf("%[1]c %[1]d", strVal)
	}
}

func TestStrings(t *testing.T) {
	s := "hello,world"
	sArray := strings.Split(s, ",")
	for i, val := range sArray {
		t.Logf("输出第%d个元素为%s \n", i, val)
	}
	sArray1 := []string{"hello", "world"}
	newA := strings.Join(sArray1, "-")
	t.Logf("合并后的字符串=%s", newA)
}

func TestConv(t *testing.T) {
	//s := strconv.Itoa(10);
	//t.Log(s)
	//if i, err := strconv.Atoi("10"); err == nil {
	//	t.Logf("转换后的数字=%d", i)
	//}
	numString := strconv.Itoa(10)
	fmt.Printf("转换后的字符串=%s \n", numString)

	if num, err := strconv.Atoi("20"); err == nil {
		fmt.Printf("转换成整形的数字=%d \n", num)
	}

}
