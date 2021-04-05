package main

import (
	jwt "github.com/appleboy/gin-jwt/v2"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	db *sqlx.DB
)

func init() {
	// TODO: 运行时记得修改数据库密码
	db = sqlx.MustConnect("mysql", "root:asdf;lkj@(localhost:3306)/ishopping")
}

func main() {
	r := gin.Default()

	authMiddleware := initAuthMiddleware()

	r.POST("/api/login", authMiddleware.LoginHandler)
	r.POST("/api/register", registerHandler)

	r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
		claims := jwt.ExtractClaims(c)
		log.Printf("NoRoute claims: %#v\n", claims)
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	auth := r.Group("/api")
	auth.Use(authMiddleware.MiddlewareFunc())
	auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	auth.GET("/query_buyer_by_id", queryBuyerByIdHandler)
	auth.GET("/query_buyer_by_username", queryBuyerByUsernameHandler)
	auth.POST("/update_buyer_info", updateBuyerInfoHandler)
	auth.GET("/commodity_search", commoditySearchHandler)
	auth.GET("/commodity_detail", commodityDetailHandler)
	auth.POST("/commodity_add", commodityAddHandler)

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
