package main

import (
	"database/sql"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
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

func connect(env *mysqlEnv) *sql.DB {
	dbStr := env.user + ":" + env.password + "@/" + env.database
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getMethod(c *gin.Context, env *mysqlEnv) {
	db := connect(env)
	parameters := c.Request.URL.Query()

	// 必須のパラメーターの取得
	from, exists := parameters["from"]
	if !exists {
		log.Println("parameter 'from' is lacking")
	}
	to, exists := parameters["to"]
	if !exists {
		log.Println("parameter 'to' is lacking")
	}

	// 必須ではないパラメーターの取得
	var ifearning string
	var transactionType string
	var ifcash string
	if ifearningParam, exists := parameters["ifearning"]; exists {
		ifearning = ifearningParam[0]
	}
	if typeParam, exists := parameters["type"]; exists {
		transactionType = typeParam[0]
	}
	if ifcashParam, exists := parameters["ifcash"]; exists {
		ifcash = ifcashParam[0]
	}
}

func postMethod(c *gin.Context, env *mysqlEnv) {

}

func putMethod(c *gin.Context, env *mysqlEnv) {

}

func deleteMethod(c *gin.Context, env *mysqlEnv) {

}
