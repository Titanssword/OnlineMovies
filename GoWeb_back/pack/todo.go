package pack

type Todo struct {
	Id      int     `json:"id"`
	Name 	string   `json:"name"`
	Url     string   `json:"url"`
}

type TodoList []Todo
