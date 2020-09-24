package actionsDb

import (
	"database/sql"
	_"github.com/go-sql-driver/mysql"
)

//全局数据库变量
var Db *sql.DB

//打开数据库
func OpenDb() error {
	if Db != nil{//如果MovieSpiderDb不为nil，说明已经调用过OpenDB函数了，
		return nil
	}
	db, err := sql.Open("mysql","root:dudu@tcp(127.0.0.1:3306)/lol?charset=utf8")
	if err != nil{
		return err
	}
	Db = db
	return nil
}

//关闭数据库
func CloseDb() error {
	if Db != nil{//如果Db不为nil，说明已经调用过OpenDB函数了，
		err := Db.Close()
		return err
	}
	return nil
}