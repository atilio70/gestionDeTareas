package main

import (
	"log"
	"task-manager-api/db"
	"task-manager-api/handlers"

	"github.com/gin-gonic/gin"
)

func main() {
	// Inicializar la base de datos
	db.InitDB()

	// Crear una nueva instancia del router de Gin
	r := gin.Default()

	// Cargar los archivos de plantillas desde la carpeta "templates"
	r.LoadHTMLGlob("templates/*")

	// Definir las rutas principales
	r.GET("/", handlers.IndexHandler)                       // Página principal que muestra las tareas
	r.GET("/tasks/create", handlers.CreateFormHandler)      // Formulario para crear nuevas tareas
	r.POST("/tasks/create", handlers.CreateHandler)         // Endpoint para crear una tarea
	r.GET("/tasks/edit/:id", handlers.EditFormHandler)      // Formulario para editar tareas
	r.POST("/tasks/update/:id", handlers.UpdateHandler)     // Endpoint para actualizar una tarea
	r.POST("/tasks/delete/:id", handlers.DeleteHandler)     // Endpoint para eliminar una tarea
	r.POST("/tasks/complete/:id", handlers.CompleteHandler) // Endpoint para completar una tarea

	// Ejecutar el servidor en el puerto 8080
	if err := r.Run("0.0.0.0:8080"); err != nil {
		log.Fatal("Error al ejecutar el servidor:", err)
	}
}
