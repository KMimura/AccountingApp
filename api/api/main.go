package main

import (
	"database/sql"
	"log"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

// DBに接続するための情報
type mysqlEnv struct {
	database string
	user     string
	password string
}

// 取引に関するデータ
type transactionData struct {
	amount          int
	date            time.Time
	transactionType string
	ifEarning       bool
	ifCash          bool
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
	dbStr := env.user + ":" + env.password + "@tcp(database)/" + env.database + "?parseTime=true"
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

func getMethod(c *gin.Context, env *mysqlEnv) *[]transactionData {
	db := connect(env)
	parameters := c.Request.URL.Query()

	// 必須のパラメーターの取得
	fromParam, exists := parameters["from"]
	if !exists {
		log.Println("parameter 'from' is lacking")
		return nil
	}
	from := fromParam[0]
	toParam, exists := parameters["to"]
	if !exists {
		log.Println("parameter 'to' is lacking")
		return nil
	}
	to := toParam[0]

	// 必須ではないパラメーターの取得
	var ifEarning string
	var transactionType string
	var ifCash string
	if ifEarningParam, exists := parameters["ifearning"]; exists {
		ifEarning = ifEarningParam[0]
	}
	if typeParam, exists := parameters["type"]; exists {
		transactionType = typeParam[0]
	}
	if ifCashParam, exists := parameters["ifcash"]; exists {
		ifCash = ifCashParam[0]
	}

	// SQLインジェクション対策
	testValues := []*string{&from, &to, &ifEarning, &transactionType, &ifCash}
	forbiddenChars := []string{";", "-", "'"}
	for _, v := range testValues {
		for _, c := range forbiddenChars {
			if strings.Contains(*v, c) {
				*v = strings.Replace(*v, c, "", -1)
			}
		}
	}

	// クエリの組み立て
	query := "select * from " + env.database + " where date between " + from + " and " + to
	if ifEarning != "" {
		query += " and ifearning = " + ifEarning
	}
	if transactionType != "" {
		query += " and type = " + transactionType
	}
	if ifCash != "" {
		query += " and ifcash" + ifCash
	}
	query += ";"

	// クエリの送信
	rows, err := db.Query(query)
	if err != nil {
		log.Println(err.Error())
		panic(err)
	}
	defer rows.Close()
	// 結果を格納する
	var results []transactionData
	for rows.Next() {
		var date time.Time
		var amount int
		var transactionType string
		var ifEarning bool
		var ifCash bool
		if err := rows.Scan(&date, &amount, &transactionType, &ifEarning, &ifCash); err != nil {
			log.Println(err.Error())
			panic(err)
		}
		result := transactionData{date: date, amount: amount, transactionType: transactionType, ifEarning: ifEarning, ifCash: ifCash}
		results = append(results, result)
	}
	return &results
}

func postMethod(c *gin.Context, env *mysqlEnv) {

}

func putMethod(c *gin.Context, env *mysqlEnv) {

}

func deleteMethod(c *gin.Context, env *mysqlEnv) {

}
