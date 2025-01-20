package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"

	"github.com/gin-gonic/gin"
)

func EditFormHandler(c *gin.Context) {
	id := c.Param("id")

	row := db.DB.QueryRow(`
        SELECT t.id, t.titulo, t.descripcion, t.estado, t.categoria_id 
        FROM tareas t WHERE t.id = ?`, id)

	var tarea Tarea
	if err := row.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado, &tarea.CategoriaID); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la tarea"})
		return
	}

	categorias, err := GetCategorias()
	if err != nil {
		log.Printf("Error al obtener categorías: %v", err)
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":      "Editar tarea",
		"Template":   "edit",
		"Tarea":      tarea,
		"Categorias": categorias,
	})
}

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	titulo := c.PostForm("titulo")
	descripcion := c.PostForm("descripcion")
	estado := c.PostForm("estado")
	categoriaID := c.PostForm("categoria_id")

	query := `
        UPDATE tareas 
        SET titulo = ?, 
            descripcion = ?, 
            estado = ?, 
            categoria_id = ?
        WHERE id = ?`

	log.Printf("Updating task ID: %s with categoria_id: %s", id, categoriaID)

	_, err := db.DB.Exec(query, titulo, descripcion, estado, categoriaID, id)
	if err != nil {
		log.Printf("Error updating task: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la tarea"})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
