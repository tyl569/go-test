package Class05

import "testing"

const (
	Monday = iota + 1
	Tuesday
	Wednesday
	Thursday
	Friday
	Saturday
	Sunday
)

func TestConst(t *testing.T) {
	t.Log(Monday)
	t.Log(Tuesday)
}
