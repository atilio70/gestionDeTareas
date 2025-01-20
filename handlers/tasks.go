package handlers

import (
	"database/sql"
	"log"
	"net/http"
	"task-manager-api/db"
	"time"

	"github.com/gin-gonic/gin"
)

type Tarea struct {
	ID              int
	Titulo          string
	Descripcion     string
	Estado          string
	CategoriaID     *int
	CategoriaNombre string
	CategoriaColor  string
	FechaCreacion   time.Time
}

func IndexHandler(c *gin.Context) {
	categoriaID := c.Query("categoria")

	query := `
        SELECT 
        	t.id, 
            t.titulo, 
            t.descripcion, 
            t.estado, 
            DATE_FORMAT(t.fecha_creacion, '%Y-%m-%d') as fecha_creacion,
            t.categoria_id,
            COALESCE(c.nombre, 'Sin categoría') as categoria_nombre,
            COALESCE(c.color, 'primary') as categoria_color
        FROM tareas t 
        LEFT JOIN categorias c ON t.categoria_id = c.id`

	var rows *sql.Rows
	var err error

	if categoriaID != "" {
		query += " WHERE t.categoria_id = ?"
		rows, err = db.DB.Query(query, categoriaID)
	} else {
		rows, err = db.DB.Query(query)
	}

	var tareas []Tarea
	for rows.Next() {
		var tarea Tarea
		var fechaCreacion string
		err := rows.Scan(
			&tarea.ID,
			&tarea.Titulo,
			&tarea.Descripcion,
			&tarea.Estado,
			&fechaCreacion,
			&tarea.CategoriaID,
			&tarea.CategoriaNombre,
			&tarea.CategoriaColor,
		)
		if err != nil {
			log.Printf("Error scanning task: %v", err)
			continue
		}
		tarea.FechaCreacion, _ = time.Parse("2006-01-02", fechaCreacion)
		tareas = append(tareas, tarea)
	}

	categorias, err := GetCategorias()
	if err != nil {
		log.Printf("Error al obtener las categorias: %v", err)
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":           "Lista de Tareas",
		"Template":        "index",
		"Tareas":          tareas,
		"Categorias":      categorias,
		"CategoriaActual": categoriaID,
	})
}
