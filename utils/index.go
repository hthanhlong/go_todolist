package utils

type path struct {
	GetTodo    string
	CreateTodo string
}

var Path = path{
	GetTodo:    "/todos",
	CreateTodo: "todos/:id",
}
