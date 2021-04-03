package main

import (
	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"log"
	"net/http"
)

var db *sqlx.DB

func init() {
	db = sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
}

func main() {
	r := gin.Default()

	r.POST("/api/login", loginHandler)

	r.POST("/api/register", registerHandler)

	r.GET("/api/query_buyer_by_id", queryBuyerById)

	r.GET("/api/query_buyer_by_username", queryBuyerByUsername)

	err := r.Run()
	if err != nil {
		log.Println(err)
	}
}

// 成功的返回
func JsonOK(c *gin.Context, data gin.H) {
	c.JSON(http.StatusOK, gin.H{
		"code": 0,
		"msg":  "",
		"data": data,
	})
	c.Abort()
}

// 错误的返回
func JsonErr(c *gin.Context, msg string) {
	c.JSON(http.StatusOK, gin.H{
		"code": -1,
		"msg":  msg,
		"data": "",
	})
	c.Abort()
}
