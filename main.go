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
	r.Use(CORSMiddleware())

	r.POST("/api/login", loginHandler)
	r.POST("/api/register", registerHandler)
	//r.POST("/api/commodity_add",commodityAddHandler)
	r.GET("/api/commodity_search", commoditySearchHandler)
	r.GET("/api/commodity_detail", commodityDetailHandler)

	auth := r.Group("/api")
	auth.Use(authorizationHandler)
	auth.POST("/commodity_add", commodityAddHandler)

	if err := r.Run(":7001"); err != nil {
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

// 未登录时的返回
func forbiddenHandler(c *gin.Context, msg string) {
	c.JSON(http.StatusForbidden, gin.H{
		"code": -1,
		"msg":  fmt.Sprintf("error: %s", msg),
		"data": "",
	})
	c.Abort()
}

func authorizationHandler(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		forbiddenHandler(c, "invalid or expired token")
		return
	}
	token := c.Request.Header["Authorization"][0][7:]
	claims, err := ParseJWT(token)
	if err != nil {
		forbiddenHandler(c, "invalid or expired token")
		return
	}
	c.Set("UserID", claims.UserID)
	c.Next()
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:63342")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "true")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

// auth.Use(authMiddleware.MiddlewareFunc())
// auth.GET("/refresh_token", authMiddleware.RefreshHandler)
// auth.GET("/query_buyer_by_id", queryBuyerByIdHandler)
// auth.GET("/query_buyer_by_username", queryBuyerByUsernameHandler)
// auth.POST("/update_buyer_info", updateBuyerInfoHandler)
