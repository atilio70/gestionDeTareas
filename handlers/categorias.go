package handlers

import (
	"log"
	"net/http"
	"task-manager-api/db"
	"task-manager-api/models"

	"github.com/gin-gonic/gin"
)

func GetCategorias() ([]models.Categoria, error) {
	rows, err := db.DB.Query("SELECT id, nombre FROM categorias")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var categorias []models.Categoria
	for rows.Next() {
		var cat models.Categoria
		if err := rows.Scan(&cat.ID, &cat.Nombre); err != nil {
			return nil, err
		}
		categorias = append(categorias, cat)
	}
	return categorias, nil
}

func ListCategoriasHandler(c *gin.Context) {
	categorias, err := GetCategorias()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al obtener categorías"})
		return
	}

	c.JSON(http.StatusOK, categorias)
}

func GetCategoriasFormHandler(c *gin.Context) {
	categorias, err := GetCategorias()
	if err != nil {
		log.Printf("Error al obtener categorías: %v", err)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"Title":    "Error",
			"Template": "error",
			"Error":    "Error al cargar categorías",
		})
		return
	}

	c.HTML(http.StatusOK, "base.html", gin.H{
		"Title":      "Nueva Categoría",
		"Template":   "categoria",
		"Categorias": categorias,
	})
}

func CreateCategoriaHandler(c *gin.Context) {
	nombre := c.PostForm("nombre")

	if nombre == "" {
		c.HTML(http.StatusBadRequest, "base.html", gin.H{
			"Title":    "Error",
			"Template": "categoria",
			"Error":    "El nombre de la categoría es requerido",
		})
		return
	}

	_, err := db.DB.Exec("INSERT INTO categorias (nombre) VALUES (?)", nombre)
	if err != nil {
		log.Printf("Error al crear categoría: %v", err)
		c.HTML(http.StatusInternalServerError, "base.html", gin.H{
			"Title":    "Error",
			"Template": "categoria",
			"Error":    "Error al crear la categoría",
		})
		return
	}

	c.Redirect(http.StatusFound, "/")
}
