package Class35

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestFunction(t *testing.T) {
	inputs := [3]int{1, 3, 4}
	outputs := [3]int{1, 9, 16}
	for i := 0; i < len(inputs); i++ {
		ret := square(inputs[i])
		assert.Equal(t, outputs[i], ret)
	}
}
