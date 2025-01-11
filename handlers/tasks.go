package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Tarea struct {
	ID            int
	Titulo        string
	Descripcion   string
	Estado        string
	FechaCreacion time.Time
}

//IndexHandler maneja la pagina principal y muestra todas las tareas

func IndexHandler(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, titulo, descripcion, estado, fecha_creacion FROM tareas")
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las tareas"})
		return
	}
	defer rows.Close()

	var tareas []Tarea
	for rows.Next() {
		var tarea Tarea
		var fechaCreacion []byte

		if err := rows.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado, &fechaCreacion); err != nil {
			log.Printf("Error al procesar las filas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las tareas"})
			return
		}

		fecha, err := time.Parse("5006-01-02 15:04:05", string(fechaCreacion))
		if err == nil {
			tarea.FechaCreacion = fecha
		}
		tareas = append(tareas, tarea)
	}
	c.HTML(http.StatusOK, "index.html", gin.H{"tareas": tareas})
}
