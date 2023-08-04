package services

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

const (
	dbUsername = "root"
	dbHost     = "db"
	dbPort     = "3306"
	dbName     = "hitderwoche"
)

func buildConnectionString() string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", dbUsername, os.Getenv("DB_PASSWORD"), dbHost, dbPort, dbName)
}

func GetDatabase() (db *gorm.DB, err error) {
	dataSourceName := buildConnectionString()
	return gorm.Open(mysql.Open(dataSourceName), &gorm.Config{})
}

func InitDatabase() {
	fmt.Println("Initializing Database ...")
	time.Sleep(30 * time.Second)

	db, err := GetDatabase()
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	// Migrate the schema
	db.AutoMigrate(&Track{})

	fmt.Println("Database-Migration applied")
}
