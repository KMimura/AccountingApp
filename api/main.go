package main

import (
    "github.com/gin-gonic/gin"
    "log"
)

func main() {
    log_config()

    r := gin.Default()
    r.GET("/accounting-api",func(c *gin.Context) {
        c.JSON(200, gin.H{
	     "state":"success",
	})
    })
    r.POST("/accounting-api",func(c *gin.Context) {
        c.JSON(200, gin.H{
             "state":"success",
        })
    })
    r.PUT("/accounting-api",func(c *gin.Context) {
        c.JSON(200, gin.H{
             "state":"success",
        })
    })
    r.DELETE("/accounting-api",func(c *gin.Context) {
        c.JSON(200, gin.H{
             "state":"success",
        })
    })
    log.Println("Start Server")
    r.Run()
}

func log_config() {
    logFile, _ := os.OpenFile("log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
    log.SetOutput(logFile)
    log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
    log.SetPrefix("[LOG] ")
}
