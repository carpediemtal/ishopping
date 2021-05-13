package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

// TODO: 运行时记得在这里修改数据库用户名和密码
const (
	Username = "root"
	Password = "1234"
)

func init() {
	DB = sqlx.MustConnect("mysql", fmt.Sprintf("%s:%s@(localhost:3306)/ishopping", Username, Password))
}
