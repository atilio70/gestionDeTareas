package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"
	"time"

	"github.com/gin-gonic/gin"
)

func DetalleHandler(c *gin.Context) {
	id := c.Param("id")
	log.Printf("Buscando tarea con ID: %s", id)

	var tarea Tarea
	var fechaCreacionStr string

	query := `
        SELECT 
            id, 
            titulo, 
            descripcion, 
            estado,
            DATE_FORMAT(fecha_creacion, '%Y-%m-%d %H:%i:%s') as fecha_creacion
        FROM tareas 
        WHERE id = ?`

	err := db.DB.QueryRow(query, id).
		Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado, &fechaCreacionStr)

	if err != nil {
		log.Printf("Error al obtener tarea: %v", err)
		c.HTML(http.StatusNotFound, "base.html", gin.H{
			"Title":    "Error",
			"Template": "error",
			"Error":    "Tarea no encontrada",
		})
		return
	}

	tarea.FechaCreacion, err = time.Parse("2006-01-02 15:04:05", fechaCreacionStr)
	if err != nil {
		log.Printf("Error al parsear fecha: %v", err)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"Title":    "Error",
			"Template": "error",
			"Error":    "Error al procesar la fecha",
		})
		return
	}

	log.Printf("Tarea encontrada: %+v", tarea)

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":    tarea.Titulo,
		"Template": "detalle",
		"Tarea":    tarea,
	})
}
