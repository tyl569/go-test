package Class43

import (
	"encoding/json"
	"fmt"
	"testing"
)

type BasicInfo struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type JobInfo struct {
	Skill []string `json:"skill"`
}
type Employee struct {
	BasicInfo BasicInfo `json:"basic_info"`
	JobInfo   JobInfo   `json:"job_info"`
}

var jsonStr = `{
	"basic_info":{
		"name":"Mike",
		"age": 20
	},
	"job_info":{
		"skill":["Java","Go"]
	}
}`

func TestEmployee(t *testing.T) {
	e := new(Employee)
	err := json.Unmarshal([]byte(jsonStr), e)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(*e)
	if v, err := json.Marshal(e); err == nil {
		fmt.Println(string(v))
	} else {
		t.Error(err)
	}
}
