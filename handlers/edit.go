package handlers

import (
	"net/http"
	"task-manager-api/db"

	"github.com/gin-gonic/gin"
)

func EditFormHandler(c *gin.Context) {
	id := c.Param("id")
	row := db.DB.QueryRow("SELECT id, titulo, descripcion, estado FROM tareas WHERE id = ?", id)

	var tarea Tarea
	if err := row.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la tarea"})
		return
	}

	c.HTML(http.StatusOK, "edit.html", tarea)
}

func UpdateHandler(c *gin.Context) {
	id := c.Param("id")
	titulo := c.Param("titulo")
	descripcion := c.Param("descripcion")
	estado := c.Param("estado")

	_, err := db.DB.Exec("UPDATE tareas SET titulo = ?, descripcion = ?, estado = ? WHERE id = ?", titulo, descripcion, estado, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al actualizar la tarea"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
