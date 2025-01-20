package main

import (
	"fmt"
	"log"
	"task-manager-api/db"
	"task-manager-api/handlers"
	"text/template"

	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()

	r := gin.Default()

	r.SetFuncMap(template.FuncMap{
		"deref": func(i *int) int {
			if i == nil {
				return 0
			}
			return *i
		},
	})

	r.LoadHTMLGlob("templates/*")

	gin.SetMode(gin.DebugMode)

	r.Use(func(c *gin.Context) {
		fmt.Printf("Request: %s %s\n", c.Request.Method, c.Request.URL.Path)
		c.Next()
	})

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
	r.GET("/categorias", handlers.ListCategoriasHandler)
	r.GET("/categorias/create", handlers.GetCategoriasFormHandler)
	r.POST("/categories/create", handlers.CreateCategoriaHandler)

	// Servir archivos estáticos
	r.Static("/static", "./static")

	fmt.Println("Servidor iniciando en http://localhost:8080")

	if err := r.Run(":8080"); err != nil {
		log.Fatal("Error al iniciar el servidor:", err)
	}
}
