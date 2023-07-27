package main

import (
	"backend/cronjob"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/robfig/cron/v3"
)

func startCronJob() {
	// Create a new cron scheduler
	c := cron.New()

	// Schedule the cron job to run every 15 minutes
	c.AddFunc("*/15 * * * *", func() {
		fmt.Println("Running cronjob at:", time.Now())
		cronjob.UpdateData()
	})

	// Start the cron scheduler
	c.Start()
}

func startApi() {
	router := gin.Default()

	router.GET("/hello", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "Hello, world!"})
	})

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	go startCronJob()

	go startApi()

	select {}
}
