package main

import (
	controller "TodoList/controllers"
	"TodoList/database"
	"TodoList/utils"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
	"os"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	router := gin.Default()

	// InitDB
	database.InitDB()

	// router
	router.GET(utils.Path.GetTodo, controller.GetTodo)
	router.POST("/todos", CreateTodo)
	router.GET("/todos/:id", GetTodoByID)
	router.PUT("/todos/:id", UpdateTodo)
	router.DELETE("/todos/:id", DeleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}
	router.Run(":" + port)
}
