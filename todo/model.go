package todo

type TodoDecode struct {
	Mssage string `json:"mssage"`
}

type ITodo struct {
	Name   string   `json:"name"`
	Mssage []string `json:"mssage"`
}