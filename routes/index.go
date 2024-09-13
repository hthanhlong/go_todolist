package routes

type path struct {
	GetTodo     string
	CreateTodo  string
	GetTodoByID string
	UpdateTodo  string
	DeleteTodo  string
}

var Path = path{
	GetTodo:     "/todos",
	CreateTodo:  "todos",
	GetTodoByID: "/todos/:id",
	UpdateTodo:  "/todos/:id",
	DeleteTodo:  "/todos/:id",
}
