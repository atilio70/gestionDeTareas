package main

import (
	"fmt"
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

	r.LoadHTMLGlob("templates/*")

	// Configurar el modo de Gin para debug
	gin.SetMode(gin.DebugMode)

	// Middleware para debug
	r.Use(func(c *gin.Context) {
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

	// Cargar las plantillas HTML - Fix: remove return value assignment
	r.LoadHTMLGlob("templates/*")
	fmt.Println("Templates loaded successfully")

	// Definir las rutas principales
	r.GET("/", handlers.IndexHandler)
	r.GET("/tasks/create", handlers.CreateFormHandler)
	r.POST("/tasks/create", handlers.CreateHandler)
	r.GET("/tasks/edit/:id", handlers.EditFormHandler)
	r.POST("/tasks/update/:id", handlers.UpdateHandler)
	r.POST("/tasks/delete/:id", handlers.DeleteHandler)
	r.POST("/tasks/complete/:id", handlers.CompleteHandler)
	r.GET("/tasks/:id", handlers.DetalleHandler)

	// Servir archivos estáticos si los hay
	r.Static("/static", "./static")

	fmt.Println("Servidor iniciando en http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
