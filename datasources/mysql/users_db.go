package mysql

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"os"
)

var (
	Client *sql.DB
	username = os.Getenv("MYSQL_USERNAME")
	password = os.Getenv("MYSQL_PASSWORD")
	host = os.Getenv("MYSQL_HOST")
	port = os.Getenv("MYSQL_PORT")
	dbName = os.Getenv("MYSQL_DBNAME")
)
func init() {
	var err error
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		username,password,host,port,dbName)
	Client, err = sql.Open("mysql", dataSourceName)
	if err != nil {
		panic(err)
	}
	if err := Client.Ping(); err != nil {
		panic(err)
	}
	log.Println("DataBase Successfully Connected!!!")
}
