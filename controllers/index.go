package controller

import (
	"TodoList/database"
	"TodoList/structs"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetTodos(c *gin.Context) {
	var todos []structs.Todo
	database.DB.Find(&todos)
	c.JSON(http.StatusOK, gin.H{"data": todos})

}

func GetTodoByID(c *gin.Context) {
	todo := structs.Todo{}
	params := c.Param("id")
	if err := database.DB.Where("id = ?", params).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": todo})
}

func CreateTodo(c *gin.Context) {
	// create input sample
	var input structs.InputTodo

	if err := c.ShouldBindJSON(&input); err != nil {
		// response to client with an bad request
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
	}

	// save to db
	todo := structs.Todo{
		Title:  input.Title,
		Status: input.Status,
	}

	// response to client
	database.DB.Create(&todo)
	c.JSON(http.StatusOK, gin.H{
		"data": todo,
	})

}

func UpdateTodo(c *gin.Context) {
	var todo structs.Todo
	params := c.Param("id")

	if err := database.DB.Where("id = ?", params).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}

	var input structs.InputTodo

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	database.DB.Model(&todo).Updates(structs.Todo{Title: input.Title, Status: input.Status})

	c.JSON(http.StatusOK, gin.H{"data": todo})

}

func DeleteTodo(c *gin.Context) {
	var todo structs.Todo
	params := c.Param("id")
	if err := database.DB.Where("id = ?", params).First(&todo).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Todo not found"})
		return
	}
	database.DB.Delete(&todo)
	c.JSON(http.StatusOK, gin.H{"data": true})
}
