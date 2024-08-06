package config

import (
	"database/sql"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb" // Import the SQL Server driver
	"log"
	"os"
)

func GetDBConnection() (*sql.DB, error) {
	// Replace these with your database configuration details
	server := os.Getenv("DB_SERVER")
	port := os.Getenv("DB_PORT")
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_NAME")

	// Connection string
	connString := fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s",
		user, password, server, port, database)

	// Open database connection
	db, err := sql.Open("sqlserver", connString)
	if err != nil {
		return nil, err
	}

	// Test the connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	log.Println("Connected to SQL Server!")
	return db, nil
}
