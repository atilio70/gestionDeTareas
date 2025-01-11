package handlers

import (
	"net/http"
	"task-manager-api/db"

	"github.com/gin-gonic/gin"
)

func DeleteHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("DELETE FROM tareas WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la tarea"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
