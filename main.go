package main

import (
	"fmt"
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
	r.Use(func(c *gin.Context) {
		c.Header("Access-Control-Allow-Origin", "*")
	})

	r.POST("/api/login", loginHandler)
	r.POST("/api/register", registerHandler)

	// r.NoRoute(authMiddleware.MiddlewareFunc(), func(c *gin.Context) {
	// 	claims := jwt.ExtractClaims(c)
	// 	log.Printf("NoRoute claims: %#v\n", claims)
	// 	c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	// })

	r.GET("/api/commodity_search", commoditySearchHandler)
	r.GET("/api/commodity_detail", commodityDetailHandler)

	auth := r.Group("/api")
	auth.Use(func(c *gin.Context) {
		e := func(c *gin.Context, msg string) {
			c.JSON(http.StatusForbidden, gin.H{
				"code": -1,
				"msg":  fmt.Sprintf("error: %s", msg),
				"data": "",
			})
			c.Abort()
		}

		cookie, err := c.Request.Cookie("jwt")
		if err != nil {
			e(c, "please login first")
			return
		}

		claims, err := ParseJWT(cookie.Value)
		if err != nil {
			e(c, "token is invalid or expired")
			return
		}

		c.Set("UserID", claims.UserID)

		c.Next()
	})

	// auth.Use(authMiddleware.MiddlewareFunc())
	// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
	// auth.GET("/query_buyer_by_id", queryBuyerByIdHandler)
	// auth.GET("/query_buyer_by_username", queryBuyerByUsernameHandler)
	// auth.POST("/update_buyer_info", updateBuyerInfoHandler)
	auth.POST("/commodity_add", commodityAddHandler)

	err := r.Run(":7001")
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
