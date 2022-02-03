package Class41

import (
	"errors"
	"strconv"
)

var ToIntWrongParamsError = errors.New("input data should be [] string")

type ToIntFilter struct {
}

func NewToInterFilter() *ToIntFilter {
	return &ToIntFilter{}
}
func (toint *ToIntFilter) Process(data Request) (Response, error) {
	parts, ok := data.([]string)
	if !ok {
		return nil, ToIntWrongParamsError
	}
	ret := []int{}
	for _, part := range parts {
		s, err := strconv.Atoi(part)
		if err != nil {
			return nil, err
		}
		ret = append(ret, s)
	}
	return ret, nil
}
