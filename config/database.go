package config

import (
	"database/sql"
	"go-learn-crud-mysql/utility/baseUtility"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func ConnectDB() {
	// Load variabel lingkungan dari file .env
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Membaca variabel lingkungan untuk koneksi database
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")

	// Menghubungkan ke database MySQL
	db, err := sql.Open("mysql", dbUser+":"+dbPassword+"@("+dbHost+":"+dbPort+")/"+dbName+"?parseTime=true")
	baseUtility.CatchWithMessage(err, "Error connecting to database")

	DB = db
}
