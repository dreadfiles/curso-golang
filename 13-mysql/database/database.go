package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() (*sql.DB, error) {
	//Load values from .env
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	// MySQL database connection string
	// Cadena de conexión a la base de datos MySQL
	connectionString := fmt.Sprintf("%v:%v@tcp(%v:%v)/%v",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Open database connection
	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	// Verify connection ping
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("MySql database connection successfully")
	return db, nil
}
