package Class41

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var splitFilterWrongParamError = errors.New("input should be string")

type SplitFilter struct {
	delimiter string
}

func NewSplitFilter(delimiter string) *SplitFilter {
	return &SplitFilter{delimiter: delimiter}
}

func (sf *SplitFilter) Process(data Request) (Response, error) {
	str, ok := data.(string)
	if !ok {
		return nil, splitFilterWrongParamError
	}
	parts := strings.Split(str, sf.delimiter)
	return parts, nil
}

func TestStraightPipleline(t *testing.T) {
	spliter := NewSplitFilter(",")
	toInt := NewToInterFilter()
	sum := NewSumFiler()
	sp := NewStraightPipline("p1", spliter, toInt, sum)
	ret, err := sp.Process("1,2,3")
	if err != nil {
		t.Fatal(err)
	}
	if assert.Equal(t, ret, 6) {
		t.Log("The result is right")
	}
}
