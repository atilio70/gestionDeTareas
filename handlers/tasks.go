package handlers

import (
	"fmt"
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
	fmt.Println("IndexHandler llamado")
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

		fecha, err := time.Parse("2006-01-02 15:04:05", string(fechaCreacion))
		if err == nil {
			tarea.FechaCreacion = fecha
		}
		tareas = append(tareas, tarea)
	}
	fmt.Println("Tareas obtenidas:", tareas)
	c.HTML(http.StatusOK, "base", gin.H{
		"Title":  "Gestor de Tareas",
		"tareas": tareas,
	})
}
