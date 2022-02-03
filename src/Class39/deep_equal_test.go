package Class39

import (
	"errors"
	"reflect"
	"testing"
)

type Employee struct {
	EmployeeId string
	Name       string `json:"name"`
	Age        int
}

type Customer struct {
	CookieId string
	Name     string
	Age      int
}

func TestDeepEqual(t *testing.T) {
	map1 := map[int]string{1: "one", 3: "three", 2: "two"}
	map2 := map[int]string{1: "one", 2: "two", 3: "three"}
	t.Log(reflect.DeepEqual(map1, map2))
	t.Log(compare(map1, map2))

	s1 := []int{1, 2, 4, 3}
	s2 := []int{1, 2, 3, 4}
	s3 := []int{1, 2, 3, 4}
	t.Log("s1 == s2 ?", reflect.DeepEqual(s1, s2))
	t.Log("s2 == s3 ?", reflect.DeepEqual(s2, s3))

}

//比较下使用reflect比较的性能
func BenchmarkDeepEqual1(b *testing.B) {
	map1 := map[int]string{1: "one", 2: "two", 3: "three"}
	map2 := map[int]string{1: "one", 2: "two", 3: "three"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		reflect.DeepEqual(map1, map2)
	}
	b.StopTimer()
}

// 比较下手动map比较的性能
func BenchmarkDeepEqual2(b *testing.B) {
	map1 := map[int]string{1: "one", 2: "two", 3: "three"}
	map2 := map[int]string{1: "one", 2: "two", 3: "three"}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		compare(map1, map2)
	}
	b.StopTimer()
}

//手动比较map
func compare(map1 map[int]string, map2 map[int]string) bool {
	if len(map1) != len(map2) {
		return false
	}
	for ele, value := range map1 {
		if _, ok := map2[ele]; !ok {
			return false
		}
		if value != map2[ele] {
			return false
		}
	}
	return true
}

func TestFillNameAndAge(t *testing.T) {
	setting := map[string]interface{}{
		"Name": "张三", "Age": 20,
	}
	e := Employee{}
	if err := fillSettingNameAndAge(&e, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(e)
	c := new(Customer)
	if err := fillSettingNameAndAge(c, setting); err != nil {
		t.Fatal(err)
	}
	t.Log(*c)
}

func fillSettingNameAndAge(st interface{}, setting map[string]interface{}) error {
	if reflect.TypeOf(st).Kind() != reflect.Ptr {
		if reflect.TypeOf(st).Elem().Kind() != reflect.Struct {
			return errors.New("the first param should be a pointer to a struct")
		}
	}
	if setting == nil {
		return errors.New("setting is nil")
	}
	var (
		field reflect.StructField
		ok    bool
	)
	for k, v := range setting {
		if field, ok = (reflect.ValueOf(st)).Elem().Type().FieldByName(k); !ok {
			continue
		}
		if field.Type == reflect.TypeOf(v) {
			vstr := reflect.ValueOf(st)
			vstr = vstr.Elem()
			vstr.FieldByName(k).Set(reflect.ValueOf(v))
		}

	}
	return nil
}
