package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	logConfig()

	r := gin.Default()
	r.GET("/accounting-api", func(c *gin.Context) {
		log.Println("GET")
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.POST("/accounting-api", func(c *gin.Context) {
		log.Println("POST")
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.PUT("/accounting-api", func(c *gin.Context) {
		log.Println("PUT")
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.DELETE("/accounting-api", func(c *gin.Context) {
		log.Println("DELETE")
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	log.Println("Start Server")
	r.Run(":8080")
}

func logConfig() {
	logFile, _ := os.OpenFile("log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("[LOG] ")
}
