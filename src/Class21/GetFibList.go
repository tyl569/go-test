package Class21

import (
	"errors"
	"fmt"
)

// 初始化函数，自动掉调用
// 一个包可以有多个init，会按照顺序执行一次
func init() {
	fmt.Println("run init1")
}
func init() {
	fmt.Println("run init2")
}
func GetFibList(n int) ([]int, error) {
	if n < 1 || n > 100 {
		return nil, errors.New("N is cannot less 1 or cannot larger 100")
	}
	fib := []int{1, 1}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib, nil
}
