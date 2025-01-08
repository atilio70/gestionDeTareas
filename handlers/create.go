package handlers

import (
	"net/http"
	"taskmanager/db"

	"github.com/gin-gonic/gin"
)

func CreateFormHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "create.html", nil)
}

func CreateHandler(c *gin.Context) {
	titulo := c.PostForm("titulo")
	descripcion := c.PostForm("descripcion")
	estado := PostForm("estado")

	_, err := db.DB.Exec("INSERT INTO tareas (titulo, descripcion, estado) VALUES (?, ?, ?)", titulo, descripcion, estado)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
		return
	}
	c.Redirect(http.StatusFound, "/")
}
