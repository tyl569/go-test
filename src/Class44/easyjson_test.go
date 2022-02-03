package Class44

import (
	"encoding/json"
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

func BenchmarkEmployee_MarshalJSON(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := new(Employee)
		e.UnmarshalJSON([]byte(jsonStr))
		if _, err := e.MarshalJSON(); err == nil {
		}
	}
	b.StopTimer()
}

func BenchmarkEmployee(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		e := new(Employee)
		json.Unmarshal([]byte(jsonStr), e)
		if _, err := json.Marshal(e); err == nil {
		}
	}
	b.StopTimer()
}
