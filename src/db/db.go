package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

func init() {
	// TODO: 运行时记得修改数据库密码
	DB = sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
}
