package services

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

const (
	dbUsername = "root"
	dbPassword = "mysecretpassword"
	dbHost     = "db"
	dbPort     = "3306"
	dbName     = "hitderwoche"
)

func buildConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, dbPassword, dbHost, dbPort, dbName)
}

func createTableIfNotExists(db *sql.DB) error {
	createTableQuery := "CREATE TABLE IF NOT EXISTS Tracks(id INT AUTO_INCREMENT PRIMARY KEY, name VARCHAR(255) NOT NULL);"
	_, err := db.Exec(createTableQuery)
	return err
}

func InitDatabase() {
	fmt.Println("Initializing Database ...")
	time.Sleep(30 * time.Second)

	dataSourceName := buildConnectionString()
	db, err := sql.Open("mysql", dataSourceName)
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
