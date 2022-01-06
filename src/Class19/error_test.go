package Class19

import (
	"errors"
	"testing"
)

func Test(t *testing.T) {
	if fib, err := GetFeb(1000); err != nil {
		if err == cannotLessErr {
			t.Log("数字不能小于数字1")
		}
		if err == cannotLargerErr {
			t.Log("数字不能大于100")
		}
		return
	} else {
		t.Log(fib)
	}
}

var cannotLessErr = errors.New("N is cannot less than 1")
var cannotLargerErr = errors.New("N is cannot larger than 100")

// 返回一个斐波那契数列，和一个异常
func GetFeb(n int) ([]int, error) {
	if n < 1 {
		return nil, cannotLessErr
	}

	if n > 100 {
		return nil, cannotLargerErr
	}
	if n == 1 {
		return []int{1}, nil
	}
	if n == 2 {
		return []int{1, 2}, nil
	}
	fib := []int{1, 2}
	for i := 2; i < n; i++ {
		fib = append(fib, fib[i-2]+fib[i-1])
	}
	return fib, nil
}
