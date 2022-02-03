package Class44

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
