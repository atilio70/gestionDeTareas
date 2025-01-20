package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateFormHandler(c *gin.Context) {
	categorias, err := GetCategorias()
	if err != nil {
		log.Printf("Error al obtener categorías: %v", err)
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":      "Crear Tarea",
		"Template":   "create",
		"Categorias": categorias,
	})
}

func CreateHandler(c *gin.Context) {
	titulo := c.PostForm("titulo")
	descripcion := c.PostForm("descripcion")
	estado := c.DefaultPostForm("estado", "pendiente")
	categoriaID := c.PostForm("categoria_id")

	_, err := db.DB.Exec(`
        INSERT INTO tareas (titulo, descripcion, estado, categoria_id, fecha_creacion) 
        VALUES (?, ?, ?, ?, ?)`,
		titulo, descripcion, estado, categoriaID, time.Now())

	if err != nil {
		log.Printf("Error al crear la tarea: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
