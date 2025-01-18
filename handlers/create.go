package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateFormHandler(c *gin.Context) {
	// Renderiza la plantilla base, pero en este caso cargamos el formulario de creación
	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":    "Crear Tarea",
		"Template": "create",
	})
}

func CreateHandler(c *gin.Context) {
	// Aquí obtienes los datos del formulario y creas la tarea.
	titulo := c.PostForm("titulo")
	descripcion := c.PostForm("descripcion")
	estado := c.DefaultPostForm("estado", "pendiente")

	// Inserta los datos en la base de datos (asegúrate de validar los datos).
	_, err := db.DB.Exec("INSERT INTO tareas (titulo, descripcion, estado, fecha_creacion) VALUES (?, ?, ?, ?)",
		titulo, descripcion, estado, time.Now())
	if err != nil {
		log.Printf("Error al crear la tarea: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}

	// Redirige al usuario de vuelta a la página principal después de crear la tarea
	c.Redirect(http.StatusFound, "/")
}
