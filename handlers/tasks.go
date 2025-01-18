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

func IndexHandler(c *gin.Context) {
	rows, err := db.DB.Query("SELECT id, titulo, descripcion, estado, fecha_creacion FROM tareas")
	if err != nil {
		log.Printf("Error al ejecutar la consulta: %v", err)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"Title":    "Error",
			"Template": "index",
			"Error":    "No se pudieron obtener las tareas",
		})
		return
	}
	defer rows.Close()

	var tareas []Tarea
	for rows.Next() {
		var tarea Tarea
		var fechaCreacion []byte
		if err := rows.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado, &fechaCreacion); err != nil {
			log.Printf("Error al escanear tarea: %v", err)
			continue
		}
		tarea.FechaCreacion, _ = time.Parse("2006-01-02", string(fechaCreacion))
		tareas = append(tareas, tarea)
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":    "Lista de Tareas",
		"Template": "index",
		"Tareas":   tareas, // Asegúrate de que este nombre coincida con el template
	})
}
