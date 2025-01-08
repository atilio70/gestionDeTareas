package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

//InitDB inicia la conexion con la base de datos y crea la tabla si no existe
func InitDB() {
	user := "root"
	password := "lasetenta"
	host := "127.0.0.1:3306"
	database := "tareas_db"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s", user, password, host, database)
	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error al conectar con la base de datos:", err)
	}

	if err := DB.Ping(), err != nil {
		log.fatal("Error al verificar la conexión con la base de datos", err)
	}

	crearTabla := `
		CREATE TABLE IF NOT EXISTS tareas (
		id INT AUTO_INCREMENT PRIMARY KEY,
		titulo VARCHAR(255) NOT NULL,
		descripcion TEXT,
		estado VARCHAR(50) NOT NULL,
		fecha_creacion DATETIME DEFAULT CURRENT_TIMESTAMP
		);
	`

	if _, err := DB.Exec(crearTabla): err != nil {
		log.Fatal("Error al crear la tabla", err)
	}
}