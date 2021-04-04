package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func init() {
	db = sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
}

func main() {
	r := gin.Default()

	r.POST("/api/login", loginHandler)

	r.POST("/api/register", registerHandler)

	r.GET("/api/query_buyer_by_id", queryBuyerByIdHandler)

	r.GET("/api/query_buyer_by_username", queryBuyerByUsernameHandler)

	r.POST("/api/update_buyer_info", updateBuyerInfoHandler)

	r.GET("/api/commodity_search", commoditySearchHandler)

	r.GET("/api/commodity_detail", commodityDetailHandler)

	r.POST("api/commodity_add",commodityAddHandler)

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
