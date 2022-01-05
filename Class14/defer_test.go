package Class14

import (
	"fmt"
	"testing"
)

/**
定义sum函数，入参ops是可变的int类型
*/
func Sum(ops ...int) {
	ret := 0;
	for _, op := range ops {
		ret += op
	}
	fmt.Println(ret)
}

func TestVarParam(t *testing.T) {
	Sum(1, 2, 3, 4)
	Sum(1, 3, 6, 7, 8, 9, 9)

}

func ClearResource1() {
	fmt.Println("Clear Resource1");
}

func ClearResource2() {
	fmt.Println("Clear Resource2");
}

func ClearResource3() {
	fmt.Println("Clear Resource3");
}

func TestDefer(t *testing.T) {
	defer ClearResource1();
	defer ClearResource2();
	defer ClearResource3();
	fmt.Println("execute something")
}

