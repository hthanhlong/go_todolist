package structs

type Todo struct {
	ID     uint   `json:"id" gorm:"primaryKey"`
	Title  string `json:"title"`
	Status bool   `json:"status"`
}

type InputTodo struct {
	Title  string `json:"title" binding:"required"`
	Status bool   `json:"status" binding:"required"`
}
