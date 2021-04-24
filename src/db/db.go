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
<<<<<<< HEAD
	DB = sqlx.MustConnect("mysql", "root:root@(localhost:3306)/ishopping")
=======
	DB = sqlx.MustConnect("mysql", "root:1234@(localhost:3306)/ishopping")
>>>>>>> f5564b6decd2bfec35f7cdf56a23b2e7ae2856d8
}
