package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

type mysqlEnv struct {
	database string
	user     string
	password string
}

func main() {
	logConfig()
	env := loadEnvVariables()

	r := gin.Default()
	r.GET("/accounting-api", func(c *gin.Context) {
		getMethod(c, env)
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.POST("/accounting-api", func(c *gin.Context) {
		postMethod(c, env)
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.PUT("/accounting-api", func(c *gin.Context) {
		putMethod(c, env)
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	r.DELETE("/accounting-api", func(c *gin.Context) {
		deleteMethod(c, env)
		c.JSON(200, gin.H{
			"state": "success",
		})
	})
	log.Println("Start Server")
	r.Run(":8080")
}

// logConfig ログの諸設定を行う
func logConfig() {
	logFile, _ := os.OpenFile("log/log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	log.SetOutput(logFile)
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.SetPrefix("[LOG] ")
}

// loadEnvVariables 環境変数を読み込む
func loadEnvVariables() *mysqlEnv {
	var env mysqlEnv
	env.database = os.Getenv("MYSQL_DATABASE")
	env.user = os.Getenv("MYSQL_USER")
	env.password = os.Getenv("MYSQL_PASSWORD")
	return &env
}

func getMethod(c *gin.Context, env *mysqlEnv) {

}

func postMethod(c *gin.Context, env *mysqlEnv) {

}

func putMethod(c *gin.Context, env *mysqlEnv) {

}

func deleteMethod(c *gin.Context, env *mysqlEnv) {

}
