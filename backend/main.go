package main

import (
	"backend/cronjob"
	"backend/services"
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
	c.AddFunc("* * * * *", func() {
		fmt.Println("Running cronjob at:", time.Now())
		cronjob.UpdateData()
	})

	// Start the cron scheduler
	c.Start()
}

func startApi() {
	router := gin.Default()

	router.GET("/tracks", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"items": services.GetTracks()})
	})

	if err := router.Run(":5000"); err != nil {
		log.Fatal(err)
	}
}

func main() {
	services.InitDatabase()

	go startCronJob()

	go startApi()

	select {}
}
