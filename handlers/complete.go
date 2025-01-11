package handlers

import (
	"net/http"
	"task-manager-api/db"

	"github.com/gin-gonic/gin"
)

func CompleteHandler(c *gin.Context) {
	id := c.Param("id")
	_, err := db.DB.Exec("UPDATE tareas SET estado = 'completada' WHERE id = ?", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo completar la acción"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
