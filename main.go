package main

import (
	"taskmanager/db"
	"taskmanager/handlers"

	"github.com/gin-gonic/gin"
)

db.InitDB()

r := gin.Default()

//Cargar templates
r.LoadHTMLGlob("templates/*")

//Endpoints principales
r.GET("/", hadlers.IndexHandler)
r.GET("/tasks/create", handlers.CreateFormHandler)
r.POST("/tasks/create", handles.CreateHandler)
r.GET("/tasks/edit/:id", handlers.EditFormHandler)
r.POST("/tasks/update/:id", handlers.UpdateHandler)
r.POST("/tasks/delete/:id", handlers.DeleteHandler)

r.Run("0.0.0.0:8080")