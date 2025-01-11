package main

import (
	"task-manager-api/db"
	"task-manager-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	//Cargar templates
	r.LoadHTMLGlob("templates/*")

	//Endpoints principales
	r.GET("/", handlers.IndexHandler)
	r.GET("/tasks/create", handlers.CreateFormHandler)
	r.POST("/tasks/create", handlers.CreateHandler)
	r.GET("/tasks/edit/:id", handlers.EditFormHandler)
	r.POST("/tasks/update/:id", handlers.UpdateHandler)
	r.POST("/tasks/delete/:id", handlers.DeleteHandler)

	r.Run("0.0.0.0:8080")
}
