package handler

import (
	"fmt"
	"ishopping/src/jwt"
	"ishopping/src/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

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

func AuthorizationHandler(c *gin.Context) {
	if len(c.Request.Header["Authorization"]) == 0 {
		forbiddenHandler(c, "invalid or expired token")
		return
	}
	token := c.Request.Header["Authorization"][0]
	claims, err := jwt.ParseJWT(token)
	if err != nil {
		forbiddenHandler(c, "invalid or expired token")
		return
	}
	c.Set("UserID", claims.UserID)
	c.Next()
}

func AdminAuthHandler(c *gin.Context) {
	uid := c.GetInt("UserID")
	utype, err := service.GetUserTypeByUid(uid)
	if err != nil {
		forbiddenHandler(c, "can not get type")
		return
	}
	if utype != 2 {
		forbiddenHandler(c, "you are not admin")
		return
	}
	c.Next()
}
