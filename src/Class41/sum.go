package Class41

import (
	"errors"
)

var SumWrongParamsError = errors.New("input data should be int")

type sumFilter struct {
}

func NewSumFiler() *sumFilter {
	return &sumFilter{}
}

func (sf *sumFilter) Process(data Request) (Response, error) {
	elems, ok := data.([]int)
	if !ok {
		return nil, SumWrongParamsError
	}
	ret := 0
	for _, elem := range elems {
		ret += elem
	}
	return ret, nil
}
