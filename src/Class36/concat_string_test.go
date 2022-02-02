package Class36

import (
	"bytes"
	assert2 "github.com/stretchr/testify/assert"
	"testing"
)

func TestConcatStringByBytesBuffer(t *testing.T) {
	assert := assert2.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	assert.Equal("12345", ConcatStringByWritingBuffer(elems))
}

func TestConcatStringByAdd(t *testing.T) {
	assert := assert2.New(t)
	elems := []string{"1", "2", "3", "4", "5"}
	assert.Equal("12345", ConcatStringByAdd(elems))
}

func BenchmarkTestConcatStringByAdd(b *testing.B) {
	b.ResetTimer()
	var elems []string
	elems = append(elems, "1", "2", "3", "4", "5")
	ConcatStringByAdd(elems)
	b.StopTimer()
}

func BenchmarkTestConcatStringByWritingBuffer(b *testing.B) {
	b.ResetTimer()
	var elems []string
	elems = append(elems, "1", "2", "3", "4", "5")
	ConcatStringByWritingBuffer(elems)
	b.StopTimer()
}

func ConcatStringByAdd(elems []string) string {
	var str string
	for _, elem := range elems {
		str += elem
	}
	return str
}

func ConcatStringByWritingBuffer(elems []string) string {
	var buf bytes.Buffer
	for _, elem := range elems {
		buf.WriteString(elem)
	}
	return buf.String()
}
