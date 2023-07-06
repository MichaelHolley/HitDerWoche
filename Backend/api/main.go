package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
  // Connect to the SQL database
  db, err := sql.Open("mysql", "username:password@tcp(localhost:3306)/mydatabase")
  if err != nil {
    log.Fatal(err)
  }
  defer db.Close()

  // Create a new Gin router
  router := gin.Default()

  // Define API endpoints
  router.GET("/users", func(c *gin.Context) {
    // Query the database for users
    rows, err := db.Query("SELECT * FROM users")
    if err != nil {
      log.Println(err)
      c.JSON(500, gin.H{"error": "Internal Server Error"})
      return
    }
    defer rows.Close()

    // Iterate over the results and build a response
    var users []string
    for rows.Next() {
      var name string
      if err := rows.Scan(&name); err != nil {
        log.Println(err)
        c.JSON(500, gin.H{"error": "Internal Server Error"})
        return
      }
      users = append(users, name)
    }

    // Return the response
    c.JSON(200, gin.H{"users": users})
  })

  // Run the API server
  if err := router.Run(":5000"); err != nil {
    log.Fatal(err)
  }
}
