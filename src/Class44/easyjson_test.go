package Class44

import (
	"fmt"
	"testing"
)

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age": 20
	},
	"job_info":{
		"skill":["Java","Go"]
	}
}`

func TestEasyjson(t *testing.T) {
	e := new(Employee)
	e.UnmarshalJSON([]byte(jsonStr))
	fmt.Println(e)
	if v, err := e.MarshalJSON(); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
