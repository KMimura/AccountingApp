package main

import (
	"database/sql"
	"io/ioutil"
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
	comment         string
}

type postData struct {
	ID        string `json:"id"`
	Date      string `json:"date"`
	IfEarning string `json:"ifEarning"`
	Type      string `json:"type"`
	Comment   string `json:"comment"`
	Amount    string `json:"amount"`
}

type deleteData struct {
	ID string `json:"id"`
}

func main() {
	logConfig()
	env := loadEnvVariables()

	setTables(env)

	r := gin.Default()
	r.GET("/accounting-api/", func(c *gin.Context) {
		log.Println(c.Request.URL.Host)
		log.Println(c.Request.URL.Path)
		results := getMethod(c, env)
		response := make([]interface{}, len(*results))
		for i, r := range *results {
			response[i] = map[string]interface{}{
				"amount":    r.amount,
				"comment":   r.comment,
				"date":      r.date,
				"ifearning": r.ifEarning,
				"type":      r.transactionType,
			}
		}
		c.JSON(200, response)
	})
	r.POST("/accounting-api", func(c *gin.Context) {
		log.Println(c.Request.URL.Host)
		log.Println(c.Request.URL.Path)
		res := postMethod(c, env)
		status := 200
		message := "Success"
		if !res {
			status = 500
			message = "Internal Server Error"
		}
		c.JSON(status, gin.H{
			"state": message,
		})
	})
	r.DELETE("/accounting-api", func(c *gin.Context) {
		log.Println(c.Request.URL.Host)
		log.Println(c.Request.URL.Path)
		res := deleteMethod(c, env)
		status := 200
		message := "Success"
		if !res {
			status = 500
			message = "Internal Server Error"
		}
		c.JSON(status, gin.H{
			"state": message,
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

// connect DBと接続する
func connect(env *mysqlEnv) *sql.DB {
	dbStr := env.user + ":" + env.password + "@tcp(database)/" + env.database + "?parseTime=true"
	db, err := sql.Open("mysql", dbStr)
	if err != nil {
		panic(err.Error())
	}
	return db
}

// setTables 初回にテーブルを作成する
func setTables(env *mysqlEnv) {
	db := connect(env)
	defer db.Close()
	bytes, err := ioutil.ReadFile("/usr/src/api/init.sql")
	if err != nil {
		log.Println(err)
	}
	query := string(bytes)
	_, err = db.Exec(query)
	if err != nil {
		log.Println(err)
	}
}

func getMethod(c *gin.Context, env *mysqlEnv) *[]transactionData {
	db := connect(env)
	defer db.Close()
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
	if ifEarningParam, exists := parameters["ifearning"]; exists {
		ifEarning = ifEarningParam[0]
	}
	if typeParam, exists := parameters["type"]; exists {
		transactionType = typeParam[0]
	}

	// SQLインジェクション対策
	testValues := []*string{&from, &to, &ifEarning, &transactionType}
	forbiddenChars := []string{";", "-", "'"}
	for _, v := range testValues {
		for _, c := range forbiddenChars {
			if strings.Contains(*v, c) {
				*v = strings.Replace(*v, c, "", -1)
			}
		}
	}

	// クエリの組み立て
	query := "select t_date,amount,t_type,if_earning,comment from transactions where t_date between " + from + " and " + to
	if ifEarning != "" {
		query += " and if_earning = " + ifEarning
	}
	if transactionType != "" {
		query += " and t_type = " + transactionType
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
		var comment string
		if err := rows.Scan(&date, &amount, &transactionType, &ifEarning, &comment); err != nil {
			log.Println(err.Error())
			panic(err)
		}
		result := transactionData{date: date, amount: amount, transactionType: transactionType, ifEarning: ifEarning, comment: comment}
		results = append(results, result)
	}
	return &results
}

func postMethod(c *gin.Context, env *mysqlEnv) bool {
	db := connect(env)
	defer db.Close()

	var pd postData
	c.BindJSON(&pd)

	// 必須のパラメーターの取得
	dateParam := pd.Date
	if dateParam == "" {
		log.Println("parameter 'date' is lacking")
		return false
	}
	date := dateParam
	ifEarningParam := pd.IfEarning
	if ifEarningParam == "" {
		log.Println("parameter 'ifearning' is lacking")
		return false
	}
	ifEarning := ifEarningParam
	amountParam := pd.Amount
	if amountParam == "" {
		log.Println("parameter 'amount' is lacking")
		return false
	}
	amount := amountParam

	// 必須ではないパラメーターの取得
	var transactionType string
	typeParam := pd.Type
	if typeParam != "" {
		transactionType = typeParam
	}
	var comment string
	commentParam := pd.Comment
	if commentParam != "" {
		comment = commentParam
	}
	var updateID string
	updateIDParam := pd.ID
	if updateIDParam != "" {
		updateID = updateIDParam
	}

	// SQLインジェクション対策
	testValues := []*string{&date, &ifEarning, &transactionType, &comment, &updateID, &amount}
	forbiddenChars := []string{";", "-", "'"}
	for _, v := range testValues {
		for _, c := range forbiddenChars {
			if strings.Contains(*v, c) {
				*v = strings.Replace(*v, c, "", -1)
			}
		}
	}

	var query string
	if updateID == "" {
		// 新しく追加する場合
		query = "insert into transactions (t_date,if_earning,t_type,comment,amount) values ('" + date + "'," + ifEarning + ",'" + transactionType + "','" + comment + "'," + amount + ");"
	} else {
		// アップデートする場合
		query = "update transactions set t_date='" + date + "',if_earning=" + ifEarning + ",t_type='" + transactionType + "',comment='" + "',amount=" + amount + " where id=" + updateID + ";"
	}
	log.Println(query)

	// クエリの送信
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true
}

func deleteMethod(c *gin.Context, env *mysqlEnv) bool {
	db := connect(env)
	defer db.Close()

	var dd deleteData
	c.BindJSON(&dd)

	// 必須のパラメーターの取得
	idParam := dd.ID
	if idParam == "" {
		log.Println("parameter 'id' is lacking")
		return false
	}

	// SQLインジェクション対策
	forbiddenChars := []string{";", "-", "'"}
	for _, c := range forbiddenChars {
		if strings.Contains(idParam, c) {
			idParam = strings.Replace(idParam, c, "", -1)
		}
	}

	query := "delete from transactions where id = " + idParam + ";"
	log.Println(query)

	// クエリの送信
	_, err := db.Exec(query)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	return true

}
