package dbops


import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err error
)

//初始化时创建连接
func init() {
	dbConn, err = sql.Open("mysql", "root:root1234@tcp(47.99.149.46:3306)/video_server?charset=utf8")
	if err != nil{
		panic(err.Error())
	}
}