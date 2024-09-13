package main

import (
	controller "TodoList/controllers"
	"TodoList/database"
	"TodoList/routes"
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
	router.GET(routes.Path.GetTodo, controller.GetTodos)
	router.POST(routes.Path.CreateTodo, controller.CreateTodo)
	router.GET(routes.Path.GetTodoByID, controller.GetTodoByID)
	router.PUT(routes.Path.UpdateTodo, controller.UpdateTodo)
	router.DELETE(routes.Path.DeleteTodo, controller.DeleteTodo)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port if not specified in .env
	}
	router.Run(":" + port)
}
