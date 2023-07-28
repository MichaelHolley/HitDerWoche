package services

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

const (
	driverName = "mysql"
	dbUsername = "root"
	dbPassword = "my-secret-password"
	dbHost     = "db"
	dbPort     = "3306"
	dbName     = "hitderwoche"
)

func buildConnectionString(includeDb bool) string {
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/", dbUsername, dbPassword, dbHost, dbPort)

	if includeDb {
		conn += dbName
	}

	return conn
}

func createDatabaseIfNotExists(db *sql.DB) error {
	_, err := db.Exec(fmt.Sprintf("CREATE DATABASE IF NOT EXISTS %s", dbName))
	return err
}

func createTableIfNotExists(db *sql.DB) error {
	createTableQuery := `
		CREATE TABLE IF NOT EXISTS Tracks (
			id VARCHAR(100) PRIMARY KEY,
			name VARCHAR(255),
		);`
	_, err := db.Exec(createTableQuery)
	return err
}

func InitDatabase() {
	fmt.Println("Initializing Database ...")

	// Connect to MySQL server
	dataSourceName := buildConnectionString(false)
	db, err := sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the database if it doesn't exist
	if err := createDatabaseIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	// Use the database
	dataSourceName = buildConnectionString(true)
	db, err = sql.Open(driverName, dataSourceName)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// Create the table if it doesn't exist
	if err := createTableIfNotExists(db); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Database and Table created (if not existed) successfully!")
}
