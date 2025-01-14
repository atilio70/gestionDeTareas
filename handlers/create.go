package handlers

import (
	"fmt"
	"net/http"
	"task-manager-api/db"

	"github.com/gin-gonic/gin"
)

func CreateFormHandler(c *gin.Context) {
	fmt.Println("Renderizando formulario de creación") // Debug
	c.HTML(http.StatusOK, "create.html", gin.H{
		"Title": "Crear Nueva Tarea",
	})
}

func CreateHandler(c *gin.Context) {
	titulo := c.PostForm("titulo")
	descripcion := c.PostForm("descripcion")
	estado := c.PostForm("estado")

	_, err := db.DB.Exec("INSERT INTO tareas (titulo, descripcion, estado) VALUES (?, ?, ?)", titulo, descripcion, estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
