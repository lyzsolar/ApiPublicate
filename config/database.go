package config

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// ConectarDB establece una conexión con la base de datos MySQL.
func ConectarDB() *sql.DB {
	dsn := "root:Solar123@tcp(localhost:3306)/veterinaria"
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Error conectando a la base de datos:", err)
	}
	return db
}
