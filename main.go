package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// Tarea representa una tarea en el sistema
type Tarea struct {
	ID            int
	Titulo        string
	Descripcion   string
	Estado        string
	FechaCreacion time.Time // Usar time.Time directamente
}

var db *sql.DB

// initDB inicia la conexión a la base de datos MySQL
func initDB() {
	user := "root"
	password := "lasetenta"
	host := "127.0.0.1:3306"
	database := "tareas_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)

	// Abrir la conexión
	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	// Comprobar la conexión
	if err := db.Ping(); err != nil {
		log.Fatal("Error al verificar la conexión con la base de datos", err)
	}

	// Crear la tabla si no existe
	crearTabla := `
		CREATE TABLE IF NOT EXISTS tareas (
			id INT AUTO_INCREMENT PRIMARY KEY,
			titulo VARCHAR(255) NOT NULL,
			descripcion TEXT,
			estado VARCHAR(50) NOT NULL DEFAULT 'pendiente',
			fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`
	if _, err := db.Exec(crearTabla); err != nil {
		log.Fatal("Error al crear la tabla", err)
	}
}

func main() {
	initDB()

	r := gin.Default()

	// Cargar templates
	r.LoadHTMLGlob("templates/*")

	// Endpoint para mostrar la página principal
	r.GET("/", func(c *gin.Context) {
		// Recuperar todas las tareas
		rows, err := db.Query("SELECT id, titulo, descripcion, estado, fecha_creacion FROM tareas")
		if err != nil {
			log.Printf("Error al ejecutar la consulta: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudieron obtener las tareas"})
			return
		}
		defer rows.Close()

		var tareas []Tarea
		for rows.Next() {
			var tarea Tarea
			var fechaCreacion []byte // Recibir la fecha como []byte

			// Escanear cada fila correctamente
			if err := rows.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado, &fechaCreacion); err != nil {
				log.Printf("Error al procesar las filas: %v", err)
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las tareas"})
				return
			}

			// Convertir la fecha de []byte a time.Time
			if len(fechaCreacion) > 0 {
				// Convertir el valor de []byte a time.Time
				fecha, err := time.Parse("2006-01-02 15:04:05", string(fechaCreacion))
				if err != nil {
					log.Printf("Error al convertir la fecha: %v", err)
					c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las fechas"})
					return
				}
				tarea.FechaCreacion = fecha
			}

			tareas = append(tareas, tarea)
		}

		// Si hubo un error durante la iteración de las filas
		if err := rows.Err(); err != nil {
			log.Printf("Error al iterar sobre las filas: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Error al procesar las tareas"})
			return
		}

		// Si todo salió bien, mostrar las tareas
		c.HTML(http.StatusOK, "index.html", gin.H{"tareas": tareas})
	})

	//mostrar formulario de edicion
	r.GET("/tasks/edit/:id", func(c *gin.Context) {
		id := c.Param("id")
		row := db.QueryRow("SELECT id, titulo, descripcion, estado FROM tareas WHERE id = ?", id)

		var tarea Tarea
		if err := row.Scan(&tarea.ID, &tarea.Titulo, &tarea.Descripcion, &tarea.Estado); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo obtener la tarea"})
			return
		}

		c.HTML(http.StatusOK, "edit.html", tarea)
	})

	//procesar la edicion de la tarea
	r.POST("/tasks/update/:id", func(c *gin.Context) {
		id := c.Param("id")
		titulo := c.PostForm("titulo")
		descripcion := c.PostForm("descripcion")
		estado := c.PostForm("estado")

		_, err := db.Exec("UPDATE tareas SET titulo = ?, descripcion = ?, estado = ? WHERE id = ?", titulo, descripcion, estado, id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo actualizar la tarea"})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	// Marcar tarea como completada
	r.POST("/tasks/complete/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("UPDATE tareas SET estado = 'completada' WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo completar la tarea"})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	//creando nuevas tareas

	// Endpoint para mostrar el formulario de creación
	r.GET("/tasks/create", func(c *gin.Context) {
		c.HTML(http.StatusOK, "create.html", nil)
	})

	// Endpoint para procesar la creación de una tarea nueva
	r.POST("/tasks/create", func(c *gin.Context) {
		titulo := c.PostForm("titulo")
		descripcion := c.PostForm("descripcion")
		estado := c.PostForm("estado")

		_, err := db.Exec("INSERT INTO tareas (titulo, descripcion, estado) VALUES (?, ?, ?)", titulo, descripcion, estado)
		if err != nil {
			log.Printf("Error al crear la tarea: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo crear la tarea"})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	//Eliminar una tarea
	r.POST("/tasks/delete/:id", func(c *gin.Context) {
		id := c.Param("id")

		_, err := db.Exec("DELETE FROM tareas WHERE id = ?", id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "No se pudo eliminar la tarea"})
			return
		}
		c.Redirect(http.StatusFound, "/")
	})

	r.Run("0.0.0.0:8080")
}
